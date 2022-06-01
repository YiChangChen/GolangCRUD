package models

type Result struct {
	IsSuccess     bool   `json:"isSuccess"`
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
}

type ResultWithData[T any] struct {
	Result
	Data T `json:"data"`
}
