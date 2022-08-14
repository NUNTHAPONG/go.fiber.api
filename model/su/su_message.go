package su

import "web-api/model"

type SuMessage struct {
	MessageCode string
	NameTh      string
	NameEn      string
	model.BaseModel
}

type SuMessageQuery struct {
	MessageCode string
	NameTh      string
	NameEn      string
	Xmin        int
}
