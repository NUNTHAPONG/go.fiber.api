package su

import "web-api/model"

type SuRole struct {
	RoleCode string
	NameTh   string
	NameEn   string
	RoleDesc string
	Active   bool
	model.BaseModel
}

type SuRoleQuery struct {
	RoleCode string
	NameTh   string
	NameEn   string
	RoleDesc string
	Active   bool
	Xmin     int
}
