package printer

import (
	"encoding/json"
	"fmt"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/imroc/req/v3"
	"strconv"
	"time"
)

type DaQuPrinter struct {
	printUrl string
	client   *req.Client
}

type DaQuRes struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func NewDaQu() *DaQuPrinter {
	return &DaQuPrinter{
		printUrl: "https://printer.juhesaas.com/openapi/print",
		client:   req.C(),
	}
}

func (d *DaQuPrinter) AddPrinter(config *model.BeePrinter) error {
	data := []map[string]interface{}{
		{
			"sn":   config.Code, //打印机编号
			"key":  config.Key,  //密钥
			"name": config.Name,
		},
	}
	_, err := d.post(d.printUrl, config, data)
	if err != nil {
		return err
	}
	return err
}

func (d *DaQuPrinter) Print(config *model.BeePrinter, content string) error {
	data := map[string]interface{}{
		"sn":      config.Code, //打印机编号
		"content": content,     //内容
		"copies":  1,
	}
	_, err := d.post(d.printUrl, config, data)
	return err
}

func (d *DaQuPrinter) post(url string, config *model.BeePrinter, data any) (*DaQuRes, error) {
	unix := time.Now().Unix()
	uid := util.GetUUIDStr()
	body := util.ToJsonWithoutErr(data, "")
	sign := util.Md5(fmt.Sprintf("%s%s%d%s%s", uid, config.Appid, unix, config.AppSecret, body))
	// md5(uid+appid+stime+appsecrect+请求Json内容)

	var res = &DaQuRes{}
	curlRes, err := d.client.R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"appid":        config.Appid,
		"stime":        strconv.FormatInt(unix, 10),
		"uid":          uid,
		"sign":         sign,
	}).SetSuccessResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	if curlRes.IsErrorState() {
		return nil, curlRes.Err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("请求大趋服务器失败，code:%d msg:%s", res.Code, res.Message)
	}
	return res, nil
}
