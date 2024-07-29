package util

import (
	"github.com/golang/geo/s2"
)

func GetDistance(lat, lon float64, lat2, lon2 float64) float64 {

	// 创建 S2 点
	p1 := s2.PointFromLatLng(s2.LatLngFromDegrees(lat, lon))
	p2 := s2.PointFromLatLng(s2.LatLngFromDegrees(lat2, lon2))

	// 计算球面距离
	return s2.LatLngFromPoint(p1).Distance(s2.LatLngFromPoint(p2)).Degrees() * 111.32 // 1度约等于111.32公里
}
