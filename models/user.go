package models

import (
	"encoding/json"
	"errors"
)

// User 定义请求参数结构体
type User struct {
	UserID       int64  `json:"user_id" db:"user_id"`
	Phone        string `json:"phone" db:"phone"`
	Password     string `json:"password" db:"password"`
	AccessToken  string
	RefreshToken string
}

// UnmarshalJSON 解析JSON数据(自定义的JSON反序列化方法，用于将JSON数据解析到自定义的结构体中
// ////json.Unmarshal函数默认使用结构体的字段名作为JSON键名进行解析。但是，有时候我们可能需要自定义JSON键名，或者对解析过程进行一些额外的处理，这时就可以使用UnmarshalJSON方法。
func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserID   int64  `json:"user_id" db:"user_id"`
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	err = json.Unmarshal(data, &required) //将一个JSON对象（data）解析为一个名为require的变量。
	if err != nil {
		return
	} else if len(required.Phone) == 0 {
		return errors.New("缺少必填字段 phone")
	} else if len(required.Password) == 0 {
		return errors.New("缺少必填字段 password")
	} else {
		u.UserID = required.UserID
		u.Phone = required.Phone
		u.Password = required.Password
	}
	return
}

// RegisterForm 注册请求参数
type RegisterForm struct {
	UserID   uint64 `json:"user_id" db:"user_id"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UnmarshalJSON 解析JSON数据(自定义的JSON反序列化方法，用于将JSON数据解析到自定义的结构体中
func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserID          uint64 `json:"user_id" db:"user_id"`
		Phone           string `json:"phone" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Phone) == 0 {
		return errors.New("缺少必填字段phone")
	} else if len(required.Password) == 0 {
		return errors.New("缺少必填字段 password")
	} else if required.Password != required.ConfirmPassword {
		return errors.New("两次密码不一致")
	} else {
		r.UserID = required.UserID
		r.Phone = required.Phone
	}
	return
}

// LoginForm 登录请求参数

type LoginForm struct {
	Phone    string `json:"phone"        // binding:"required"`
	Password string `json:"password"    // binding:"required"`
}
