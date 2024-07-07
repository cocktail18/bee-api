package enum

type CyTablePayMod int32

var CyTablePayModMap = map[CyTablePayMod]string{
	CyTablePayModPayBefore: "先付钱再用餐",
	CyTablePayModPayAfter:  "先用餐再付款",
}

const (
	CyTablePayModPayBefore CyTablePayMod = 0
	CyTablePayModPayAfter  CyTablePayMod = 1
)
