package main
import (
	"fmt"
	"encoding/json"
)

type goods struct {
        AfterSale        string                 `json:"afterSale"`
        BuyRewardEnd     bool                   `json:"buyRewardEnd"`
        CategoryID       int64                  `json:"categoryId"`
        DateAdd          string                 `json:"dateAdd"`
        GoodsID          int64                  `json:"goodsId"`
        GoodsIDStr       string                 `json:"goodsIdStr"`
        GoodsName        string                 `json:"goodsName"`
        ID               int64                  `json:"id"`
        InviterId        int64                  `json:"inviter_id"`
        IsProcessHis     bool                   `json:"isProcessHis"`
        IsScoreOrder     bool                   `json:"isScoreOrder"`
        LogisticsType    int                    `json:"logisticsType"`
        Number           int64                  `json:"number"`
        NumberNoFahuo    int64                  `json:"numberNoFahuo"`
        OrderID          string                 `json:"orderId"`
        Persion          int64                  `json:"persion"`
        Pic              string                 `json:"pic"`
        PriceID          int64                  `json:"priceId"`
        Property         string                 `json:"property"`
        PropertyChildIds string                 `json:"propertyChildIds"`
        Purchase         bool                   `json:"purchase"`
        ShopID           int64                  `json:"shopId"`
        Type             int64                  `json:"type"`
        UID              int64                  `json:"uid"`
        Unit             string                 `json:"unit"`
        UserID           int64                  `json:"userId"`
}

func main() {
	goodsjsonstr := `[{"propertyChildIds":",211246:1686934,211247:1686939","goodsId":1808574,"number":2,"logisticsType":0,"inviter_id":"1","goodsTimesDay":"","goodsTimesItem":""}]`
	list := make([]goods, 0, 10)
	if err := json.Unmarshal([]byte(goodsjsonstr), &list); err != nil {
		fmt.Println(err)
	}
}
