package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CrBy    string
	CrSys   string
	CrDate  time.Time
	ModBy   string
	ModSys  string
	ModDate time.Time
	Xmin    int      `json:"RowVersion"`
	State   RowState `json:"RowState"`
}

type RowState int

const (
	Normal RowState = iota
	Add
	Modified
	Remove
)

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CrBy = "admin"
	m.CrSys = "admin"
	m.CrDate = time.Now()
	m.ModBy = "admin"
	m.ModSys = "admin"
	m.ModDate = time.Now()
	return
}

func (m *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Print("BeforeSave is Working...\n")
	m.ModBy = "admin"
	m.ModSys = "admin"
	m.ModDate = time.Now()
	return
}


// func (d *BaseModel) SetAudit(tx *gorm.DB) {
// 	switch d.State {
// 	case Add:
// 		d.CrBy = "admin"
// 		d.CrSys = "admin"
// 		d.CrDate = time.Now()
// 		d.ModBy = "admin"
// 		d.ModSys = "admin"
// 		d.ModDate = time.Now()
// 	case Modified:
// 		d.ModBy = "admin"
// 		d.ModSys = "admin"
// 		d.ModDate = time.Now()
// 	}
// }
