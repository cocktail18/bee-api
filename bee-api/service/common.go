package service

import (
	"encoding/json"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/pkg/errors"
	"math"
	"sync"
)

type CommonSrv struct {
	BaseSrv
}

var commonSrvOnce sync.Once
var commonSrvInstance *CommonSrv

func GetCommonSrv() *CommonSrv {
	commonSrvOnce.Do(func() {
		commonSrvInstance = &CommonSrv{}
	})
	return commonSrvInstance
}

// GetDistance 获取距离
// lonA, latA分别为A点的纬度和经度
// lonB, latB分别为B点的纬度和经度
// 返回的距离单位为米
func (srv *CommonSrv) GetDistance(lonA, latA, lonB, latB float64) float64 {
	// 地球半径，单位米
	const R = 6367000

	c := math.Sin(latA)*math.Sin(latB)*math.Cos(lonA-lonB) + math.Cos(latA)*math.Cos(latB)
	return R * math.Acos(c) * math.Pi / 180
}

func (srv *CommonSrv) GetMapQQDistance(key string, mode string, from string, to string) (*proto.DistanceResp, error) {
	r := util.GetDefaultHttpClient()
	var res proto.Response

	resp, err := r.R().SetFormData(map[string]string{ // @todo 直接调用地图api
		"key":  key,
		"mode": mode,
		"from": from,
		"to":   to,
	}).SetResult(&res).Post(`https://api.it120.cc/common/map/qq/distance`)
	if err != nil {
		return nil, errors.Wrap(err, "获取距离失败")
	}
	if resp.IsError() {
		return nil, errors.Wrap(resp.Err, "获取距离失败")
	}
	if res.Code != 0 {
		return nil, enum.NewBussErr(nil, res.Code, res.Msg)
	}
	var data proto.DistanceResp
	bs, err := json.Marshal(res.Data)
	if err != nil {
		return nil, errors.Wrap(err, "获取距离失败")
	}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return nil, errors.Wrap(err, "获取距离失败")
	}
	return &data, nil
}
