package models

type _ResponseDiaryList struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []*ApiDiaryDetail `json:"data"`
}
