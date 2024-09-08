package printer

import (
	"encoding/json"
	"errors"
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
	addUrl   string
	delUrl   string
}

var _ Printer = &DaQuPrinter{}

type DaQuRes struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type DaQuError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *DaQuError) Error() string {
	return fmt.Sprintf("大趋服务器返回失败 code = %d msg = %s", r.Code, r.Message)
}

func NewDaQu() *DaQuPrinter {
	return &DaQuPrinter{
		printUrl: "https://printer.juhesaas.com/openapi/print",
		addUrl:   "https://printer.juhesaas.com/openapi/addPrinter",
		delUrl:   "https://printer.juhesaas.com/openapi/delPrinter",
		client:   req.C().DevMode(),
	}
}

func (d *DaQuPrinter) DelPrinter(config *model.BeePrinter, codes []string) error {
	_, err := d.post(d.delUrl, config, codes)
	if err != nil {
		return err
	}
	return err
}

func (d *DaQuPrinter) AddPrinter(config *model.BeePrinter) error {
	data := []map[string]interface{}{
		{
			"sn":   config.Code, //打印机编号
			"key":  config.Key,  //密钥
			"name": config.Name,
		},
	}
	_, err := d.post(d.addUrl, config, data)
	if err != nil {
		var daquErr = &DaQuError{}
		if errors.Is(err, daquErr) {
			if daquErr.Code == 31412 {
				return ErrAlreadyAdd
			}
		}
		return err
	}
	return err
}

func (d *DaQuPrinter) Print(config *model.BeePrinter, voice, content string) error {
	data := map[string]interface{}{
		"sn":      config.Code, //打印机编号
		"content": content,     //内容
		"copies":  1,
		"voice":   voice,
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
	}).SetBodyJsonString(body).SetSuccessResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	if curlRes.IsErrorState() {
		return nil, curlRes.Err
	}
	if res.Code != 0 {
		return nil, &DaQuError{Code: res.Code, Message: res.Message}
	}
	return res, nil
}
