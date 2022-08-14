package su

import "web-api/model"

type SuMenu struct {
	MenuCode   string
	PageCode   string
	MainMenu   string
	NameTh     string
	NameEn     string
	ProjCode   string
	ModuleCode string
	Icon       string
	Active     bool
	model.BaseModel
}

type SuMenuQuery struct {
	MenuCode   string
	PageCode   string
	MainMenu   string
	NameTh     string
	NameEn     string
	ModuleCode string
	Icon       string
	Active     bool
	Xmin       int
}
