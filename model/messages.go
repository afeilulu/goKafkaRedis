package model

type Alert struct {
	Signature    string `json:"signature"`
	Signature_id int64 `json:"signature_id"`
}

type Http struct {
	Http_user_agent string `json:"http_user_agent"`
}

type Msg struct {
	Timestamp string `json:"timestamp"`
	Src_ip    string `json:"src_ip"`
	Alert     Alert  `json:"alert"`
	Http      Http   `json:"http"`
}

