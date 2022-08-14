package su

import "web-api/model"

type SuLabel struct {
	LabelCode  string
	PageCode   string
	NameTh     string
	NameEn     string
	ProjCode   string
	ModuleCode string
	model.BaseModel
}

type SuLabelQuery struct {
	LabelCode  string
	NameTh     string
	NameEn     string
	Xmin       int
}
