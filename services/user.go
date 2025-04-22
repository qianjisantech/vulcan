/*
Copyright 2021 The DnsJia Authors.
WebSite:  https://github.com/dnsjia/luban
Email:    OpenSource@dnsjia.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package services

import (
	"errors"
	"fmt"
	"github.com/dnsjia/luban/common"
	"github.com/dnsjia/luban/models"
	"gorm.io/gorm"
)

func UserRegister(u models.User) (userInter models.User, err error) {
	var user models.User

	if !errors.Is(common.DB.Where("username = ? ", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New(fmt.Sprintf("user %v already exists", u.UserName))
	}
	err = common.DB.Create(&u).Error

	return u, err
}

func Login(l models.LoginUser) (models.User, error) {
	var user models.User

	err := common.DB.Preload("Role").Where("email = ?", l.Email).First(&user).Error
	//sql := `select  m.name,m.path, u.username,r.name from users u
	//     left join  role r on u.role_id=r.id
	//     left join  relation_role_menu rrm on r.id=rrm.role_id
	//     left join  menu m on m.id=rrm.menu_id
	//     where u.email='?' AND  password='?'`
	return user, err
}
