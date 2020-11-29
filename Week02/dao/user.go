package dao

import (
	"errors"
	"fmt"
	"week002/common"
	"week002/model"

	"github.com/jinzhu/gorm"
	xerrors "github.com/pkg/errors"
)

type UserDao struct {
}

func (u UserDao) SelectById(userid int) (*model.User, error) {

	var user model.User

	err := common.GetDB().Where("id=?", userid).First(&user).Error

	
	if errors.Is(err,gorm.ErrRecordNotFound) {

		err=xerrors.Wrap(err,fmt.Sprintf("error select userid =%d not found",userid))

		return nil,err
	}

	return &user, err
}
