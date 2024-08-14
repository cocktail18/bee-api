package printer

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
	"strconv"
	"strings"
	"time"
)

// FeiEPrinter 飞蛾打印机
type FeiEPrinter struct {
	printUrl string
	client   *req.Client
}

var _ Printer = &FeiEPrinter{}

type FeiERes struct {
	Ret                int             `json:"ret"`
	Message            string          `json:"message"`
	Data               json.RawMessage `json:"data"`
	ServerExecutedTime int             `json:"serverExecutedTime"`
}

func NewFeiE() *FeiEPrinter {
	return &FeiEPrinter{
		printUrl: "http://api.feieyun.cn/Api/Open/",
		client:   req.C(),
	}
}

func (d *FeiEPrinter) DelPrinter(config *model.BeePrinter, codes []string) error {
	_, err := d.post(d.printUrl, config, map[string]string{
		"apiname": "Open_printerDelList",
		"snlist":  strings.Join(codes, "-"),
	})
	if err != nil {
		return err
	}
	return err
}

func (d *FeiEPrinter) AddPrinter(config *model.BeePrinter) error {
	data := map[string]string{
		"apiname":        "Open_printerAddlist",
		"printerContent": config.Code,
	}
	_, err := d.post(d.printUrl, config, data)
	if err != nil {
		return err
	}
	return err
}

func (d *FeiEPrinter) Print(config *model.BeePrinter, voice, content string) error {
	data := map[string]string{
		"apiname": "Open_printMsg",           //固定
		"sn":      config.Code,               //打印机编号
		"content": content,                   //打印内容
		"times":   cast.ToString(config.Num), //打印次数
	}
	_, err := d.post(d.printUrl, config, data)
	return err
}

func (d *FeiEPrinter) post(url string, config *model.BeePrinter, data map[string]string) (*FeiERes, error) {
	itime := time.Now().Unix()
	stime := strconv.FormatInt(itime, 10)
	sig := d.sha1(config.Appid + config.Key + stime) //生成签名
	// md5(uid+appid+stime+appsecrect+请求Json内容)

	dataCopy := make(map[string]string)
	for k, v := range data {
		dataCopy[k] = v
	}
	dataCopy["user"] = config.Appid
	dataCopy["stime"] = stime
	dataCopy["sig"] = sig

	var res = &FeiERes{}
	curlRes, err := d.client.R().SetFormData(dataCopy).SetSuccessResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	if curlRes.IsErrorState() {
		return nil, curlRes.Err
	}
	if res.Ret != 0 {
		return nil, fmt.Errorf("请求飞蛾服务器失败，code:%d msg:%s", res.Ret, res.Message)
	}
	return res, nil
}

func (d *FeiEPrinter) sha1(str string) string {
	s := sha1.Sum([]byte(str))
	strsha1 := hex.EncodeToString(s[:])
	return strsha1
}
