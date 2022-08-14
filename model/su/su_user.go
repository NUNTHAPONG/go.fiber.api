package su

import (
	"github.com/google/uuid"
	"time"
	"web-api/model"
)

type SuUser struct {
	UserId          uuid.UUID
	Username        string
	Password        string
	RoleCode        string
	Active          bool
	LastLoginDate   time.Time
	model.BaseModel 
}

type SuUserQuery struct {
	Username string
	Password string
	RoleCode string
	Active   bool
	Xmin     int
}
