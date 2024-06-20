package proto

import (
	"gitee.com/stuinfer/bee-api/util"
	"math"
)

type PageReqCommon struct {
	DateAddBegin string //	时间起,格式 2018-05-16	query	false
	DateAddEnd   string //	时间止,格式 2018-05-16	query	false
	Page         int64  //	获取第几页数据,示例值(1)	query	false
	PageSize     int64  //	每页显示多少数据,示例值(50)	query	false
}

type PageRespCommon struct {
	TotalRow  int64 `json:"totalRow"`
	TotalPage int64 `json:"totalPage"`
}

func (p PageReqCommon) GetPage() int64 {
	return util.IF(p.Page > 0, p.Page, 1)
}

func (p PageReqCommon) GetPageSize() int64 {
	return util.IF(p.PageSize > 0, p.PageSize, 50)
}

func GetPageCommon(cnt, pageSize int64) PageRespCommon {
	return PageRespCommon{
		TotalRow:  cnt,
		TotalPage: int64(math.Ceil(float64(cnt) / float64(pageSize))),
	}
}
