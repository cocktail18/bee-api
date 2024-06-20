package proto

type DistanceRespElements struct {
	Distance int `json:"distance"`
}

type DistanceRespRows struct {
	Elements []DistanceRespElements `json:"elements"`
}

type DistanceRespResult struct {
	Rows []DistanceRespRows `json:"rows"`
}
type DistanceResp struct {
	Message   string             `json:"message"`
	RequestID string             `json:"request_id"`
	Result    DistanceRespResult `json:"result"`
	Status    int                `json:"status"`
}
