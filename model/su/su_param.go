package su

import "web-api/model"

type SuParam struct {
	ParamCode string
	SeqOrder  int64
	NameTh    string
	NameEn    string
	PChar1    string
	PChar2    string
	PChar3    string
	PChar4    string
	PNum1     int64
	PNum2     int64
	PNum3     int64
	PNum4     int64
	PBool1    bool
	PBool2    bool
	PBool3    bool
	PBool4    bool
	model.BaseModel
}

type SuParamQuery struct {
	ParamCode string
	SeqOrder  int
	NameTh    string
	NameEn    string
	PChar1    string
	PChar2    string
	model.BaseModel
}
