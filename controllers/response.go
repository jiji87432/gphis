package controllers

type BaseResp struct {
	Result bool        `json:"result"`
	Reason interface{} `json:"reason"`
}
