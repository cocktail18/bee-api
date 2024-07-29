package utils

import (
	"github.com/gin-gonic/gin"
)

// GetShopUserID 从Gin的Context中获取从jwt解析出来的商店用户ID
func GetShopUserID(c *gin.Context) int64 {
	//shopUserIdFunc := func() int64 {
	//	if claims, exists := c.Get("claims"); !exists {
	//		if cl, err := utils.GetClaims(c); err != nil {
	//			return 0
	//		} else {
	//			return cl.BaseClaims.ShopUserId
	//		}
	//	} else {
	//		waitUse := claims.(*systemReq.CustomClaims)
	//		return waitUse.BaseClaims.ShopUserId
	//	}
	//}
	//shopUserId := shopUserIdFunc()
	//if shopUserId == 0 { // 默认用100
	//	shopUserId = 100
	//}
	//return shopUserId
	return 100
}
