package su

import (
	"web-api/model"
)

type SuPage struct {
	PageCode   string
	NameTh     string
	NameEn     string
	RoutePath  string
	ProjCode   string
	ModuleCode string
	model.BaseModel
}

type SuPageSave struct {
	PageCode   string
	NameTh     string
	NameEn     string
	RoutePath  string
	ProjCode   string
	ModuleCode string
	model.BaseModel
}