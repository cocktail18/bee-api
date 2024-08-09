package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"strings"
	"sync"
)

type GoodsSrv struct {
	BaseSrv
}

var goodsSrvOnce sync.Once
var goodsSrvInstance *GoodsSrv

func GetGoodsSrv() *GoodsSrv {
	goodsSrvOnce.Do(func() {
		goodsSrvInstance = &GoodsSrv{}
	})
	return goodsSrvInstance
}

func (srv *GoodsSrv) GetGoodsList(c context.Context, shopId int64, categoryId int64, page, pageSize int) ([]*model.BeeShopGoods, error) {
	var list []*model.BeeShopGoods
	dbIns := db.GetDB().Where(map[string]interface{}{
		"hidden":      0,
		"category_id": categoryId,
	})
	dbIns = dbIns.Where("user_id = ?", kit.GetUserId(c))
	if shopId > 0 {
		dbIns = dbIns.Where("shop_id = ?", shopId)
	}
	err := dbIns.Order("paixu desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, err
}

func (srv *GoodsSrv) GetCategoryAll(c context.Context) ([]*model.BeeShopGoodsCategory, error) {
	var list []*model.BeeShopGoodsCategory
	err := db.GetDB().Where("user_id = ?", kit.GetUserId(c)).
		Where("is_use=?", 1).Order("paixu desc").Find(&list).Error
	return list, err
}

func (srv *GoodsSrv) GetGoodsDetail(c context.Context, id int64, regionId string) (*proto.GoodsDetailResp, error) {
	resp := &proto.GoodsDetailResp{}
	var goods model.BeeShopGoods
	var err error
	err = db.GetDB().Where("id = ?", id).Take(&goods).Error
	if err != nil {
		return nil, err
	}
	var category model.BeeShopGoodsCategory
	err = db.GetDB().Where("id = ?", goods.CategoryId).First(&category).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	var content model.BeeShopGoodsContent
	err = db.GetDB().Where("goods_id = ?", goods.Id).First(&content).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	var pics []*model.BeeShopGoodsImages
	err = db.GetDB().Where("goods_id = ?", goods.Id).Find(&pics).Error
	if err != nil {
		return nil, err
	}

	var skuList []*model.BeeShopGoodsSku
	err = db.GetDB().Where("goods_id = ?", goods.Id).Find(&skuList).Error
	if err != nil {
		return nil, err
	}
	resp.BasicInfo = &goods
	resp.Category = &category
	resp.Content = content.Content
	resp.Pics = pics
	for _, pic := range resp.Pics {
		resp.Pics2 = append(resp.Pics2, pic.Pic)
	}
	propIds := make([]int64, 0)
	parentId2props := make(map[int64][]*model.BeeShopGoodsProp)
	id2prop := make(map[int64]*model.BeeShopGoodsProp)
	for _, sku := range skuList {
		for _, propStr := range strings.Split(sku.PropertyChildIds, ",") {
			if propStr == "" {
				continue
			}
			_propIds := lo.Map(strings.Split(propStr, ":"), func(item string, _ int) int64 {
				return cast.ToInt64(item)
			})
			propIds = append(propIds, _propIds...)
		}
	}
	propIds = lo.Uniq(propIds)
	var beeProps []*model.BeeShopGoodsProp
	err = db.GetDB().Where("(id in ?) or (property_id in ?)", propIds, propIds).Order("property_id asc").Find(&beeProps).Error
	if err != nil {
		return nil, err
	}
	for _, prop := range beeProps {
		pid := prop.PropertyId
		if _, ok := parentId2props[pid]; !ok {
			parentId2props[pid] = make([]*model.BeeShopGoodsProp, 0)
		}
		id2prop[prop.Id] = prop
		parentId2props[pid] = append(parentId2props[pid], prop)
	}
	resp.Properties = make([]*model.BeeShopGoodsProp, 0, 10)
	for _, tmpProp := range parentId2props[0] {
		prop := tmpProp
		prop.ChildsCurGoods = parentId2props[prop.Id]
		resp.Properties = append(resp.Properties, prop)
	}
	for _, sku := range skuList {
		propName := make([]string, 0)
		for _, propStr := range strings.Split(sku.PropertyChildIds, ",") {
			if propStr == "" {
				continue
			}
			_propIds := lo.Map(strings.Split(propStr, ":"), func(item string, _ int) int64 {
				return cast.ToInt64(item)
			})
			propName = append(propName, id2prop[_propIds[0]].Name+":"+id2prop[_propIds[1]].Name)
		}
		sku.PropertyChildNames = strings.Join(propName, ",") + ","
	}
	resp.SkuList = skuList
	resp.Logistics, err = GetFreightTplSrv().GetBeeLogistics(c, goods.LogisticsId, regionId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (srv *GoodsSrv) GetGoodsWithSku(c context.Context, goodsId int64, propertyChildIds string) (*model.BeeShopGoods, *model.BeeShopGoodsSku, error) {
	var goods model.BeeShopGoods
	var err error
	err = db.GetDB().Where("id = ?", goodsId).Take(&goods).Error
	if err != nil {
		return nil, nil, err
	}

	var curSku model.BeeShopGoodsSku
	propertyChildIds = kit.GetDBPropertyChildIds(propertyChildIds)
	err = db.GetDB().Where("goods_id = ? and property_child_ids = ?", goods.Id, propertyChildIds).Take(&curSku).Error
	if err != nil {
		return nil, nil, err
	}
	propChildNames, err := srv.getPropertyChildNames(c, propertyChildIds)
	if err != nil {
		return nil, nil, err
	}
	curSku.PropertyChildNames = propChildNames
	return &goods, &curSku, nil
}

func (srv *GoodsSrv) getPropertyChildNames(c context.Context, propertyChildIds string) (string, error) {
	propIds := kit.GetPropIdsByStr(propertyChildIds)
	var props []*model.BeeShopGoodsProp
	err := db.GetDB().Where("id in (?) ", propIds).Find(&props).Error
	if err != nil {
		return "", err
	}

	id2prop := lo.KeyBy(props, func(item *model.BeeShopGoodsProp) int64 {
		return item.Id
	})
	propChildNames := strings.Builder{}
	for _, str := range strings.Split(propertyChildIds, ",") {
		if str == "" {
			continue
		}
		strArr := strings.Split(str, ":")
		propParent := id2prop[cast.ToInt64(strArr[0])]
		propChild := id2prop[cast.ToInt64(strArr[1])]
		if propParent == nil {
			propParent = &model.BeeShopGoodsProp{
				Name: "未知",
			}
		}
		if propChild == nil {
			propChild = &model.BeeShopGoodsProp{
				Name: "未知",
			}
		}
		propChildNames.WriteString(propParent.Name)
		propChildNames.WriteString(":")
		propChildNames.WriteString(propChild.Name)
		propChildNames.WriteString(",")
	}
	return propChildNames.String(), nil
}

func (srv *GoodsSrv) GetPropsByIds(c context.Context, propIds []int64) ([]*model.BeeShopGoodsProp, error) {
	var props []*model.BeeShopGoodsProp
	err := db.GetDB().Where("id in ?", propIds).Find(&props).Error
	if err != nil {
		return nil, err
	}
	return props, nil
}

func (srv *GoodsSrv) GetPrice(c context.Context, goodsId int64, propertyChildIds string) (*proto.GoodsPriceResp, error) {
	var goods model.BeeShopGoods
	var err error
	err = db.GetDB().Where("id = ?", goodsId).Take(&goods).Error
	if err != nil {
		return nil, err
	}
	propertyChildIds = kit.GetDBPropertyChildIds(propertyChildIds) //重排序
	propIds := kit.GetPropIdsByStr(propertyChildIds)
	props, err := srv.GetPropsByIds(c, propIds)
	if err != nil {
		return nil, err
	}
	var curSku *model.BeeShopGoodsSku
	err = db.GetDB().Where("goods_id = ? and property_child_ids = ?", goods.Id, propertyChildIds).
		Take(&curSku).Error
	if err != nil {
		return nil, errors.Wrap(err, "找不到该sku")
	}
	propsId2props := lo.KeyBy(props, func(item *model.BeeShopGoodsProp) int64 {
		return item.Id
	})

	propChildNames := strings.Builder{}
	for _, str := range strings.Split(propertyChildIds, ",") {
		if str == "" {
			continue
		}
		strArr := strings.Split(str, ":")
		propParent := propsId2props[cast.ToInt64(strArr[0])]
		propChild := propsId2props[cast.ToInt64(strArr[1])]
		propChildNames.WriteString(propParent.Name)
		propChildNames.WriteString(":")
		propChildNames.WriteString(propChild.Name)
		propChildNames.WriteString(",")
	}
	curSku.PropertyChildNames = propChildNames.String()
	resp := &proto.GoodsPriceResp{
		FxType:             curSku.FxType,
		GoodsId:            goodsId,
		Id:                 goodsId,
		OriginalPrice:      curSku.OriginalPrice,
		PingtuanPrice:      curSku.PingtuanPrice,
		Price:              curSku.Price,
		PropertyChildIds:   propertyChildIds,
		PropertyChildNames: curSku.PropertyChildNames,
		Score:              curSku.Score,
		Stores:             curSku.Stores,
	}
	return resp, nil
}

func (srv *GoodsSrv) GetGoodsAddition(c context.Context, goodsId int64) ([]*model.BeeShopGoodsAddition, error) {
	userId := kit.GetUserId(c)
	var list []*model.BeeShopGoodsAddition
	if err := db.GetDB().Model(&model.BeeShopGoodsAddition{}).
		Where("goods_id = ? and user_id = ?", goodsId, userId).Order("pid asc").Find(&list).Error; err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, enum.ErrEmpty
	}
	additionMap := make(map[int64][]*model.BeeShopGoodsAddition)
	for _, addition := range list {
		if _, ok := additionMap[addition.Pid]; !ok {
			additionMap[addition.Pid] = make([]*model.BeeShopGoodsAddition, 0)
		}
		additionMap[addition.Pid] = append(additionMap[addition.Pid], addition)
	}
	retList := make([]*model.BeeShopGoodsAddition, 0)
	for _, addition := range additionMap[0] {
		if items, ok := additionMap[addition.Id]; ok {
			addition.Items = items
		} else {
			addition.Items = make([]*model.BeeShopGoodsAddition, 0)
		}
		retList = append(retList, addition)
	}
	return retList, nil
}
