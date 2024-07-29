package proto

type UploadFileResp struct {
	Msg          string `json:"msg"`
	OriginalName string `json:"originalName"`
	Size         string `json:"size"`
	ConfigId     string `json:"configId"`
	Name         string `json:"name"`
	Id           int64  `json:"id"`
	Type         string `json:"type"` //后缀名
	Url          string `json:"url"`
}
