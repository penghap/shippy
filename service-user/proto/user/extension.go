package user

import (
	"github.com/jinzhu/gorm"
	uuidX "github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuidX.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
