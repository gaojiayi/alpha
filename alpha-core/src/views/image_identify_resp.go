package views

type ImageIdentifyResp struct {
	FileName     string        `json:"file_name"`
	IdentifyRets []IdentifyRet `json:"identify_rets"`
}

type IdentifyRet struct {
	Name   string `json:"name"`
	Percet string `json:"percet"`
}
