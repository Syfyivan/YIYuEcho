package mysql

import (
	"YiYuEcho/models"
	"YiYuEcho/settings"
	"crypto/md5"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"time"
	_ "time"
)

var db *gorm.DB

// InitDB 初始化数据库连接
func InitDB(config *settings.MySQLConfig) (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/yiyuser?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	//创建表的逻辑
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Diary{})
	if err != nil {
		return err
	}
	return nil
}

// CheckUserExist 检查用户是否存在
func CheckUserExist(userID int64) (err error) {
	var count int64
	result := db.Model(&models.User{}).Where("userID = ?", userID).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {

		//return errors.New(ErrorUsersExist)
		return errors.New("用户已存在")
	}

	//return nil
	return
}

// EncryptPassword 密码加密
func EncryptPassword(password string) (result string) {
	h := md5.New()
	h.Write([]byte(password))
	result = string(h.Sum(nil))
	return
}

// InsertUser  注册业务-向数据库中插入一条新用户
func InsertUser(user *models.User) (error error) {
	//对密码进行加密
	//user.Password = EncryptPassword(user.Password)
	//执行sql语句入库
	err := db.Create(user).Error
	return err
}

// Register  注册业务(函数首先检查用户名是否已经存在于数据库中，如果存在则返回错误。然后生成一个新的用户 ID)
func Register(user *models.User) (err error) {
	//sqlStr := "select count(user_id) from user where phone = ?"
	var count int64
	err = db.Model(&models.User{}).Where("phone = ?", user.Phone).Count(&count).Error
	if err != nil {

		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	//生成加密密码
	//user.Password = EncryptPassword(user.Password)
	//生成UserID
	//todo : 雪花算法
	//userID, err := snowflake.GetID

	userID := time.Now().UnixNano() + rand.Int63n(1000) // 使用当前时间戳加上一个随机数作为 user_id

	if err != nil {
		return errors.New("创建用户ID失败")
	}
	//把用户插入数据库
	user.UserID = userID
	//user.Password = password
	err = db.Create(user).Error
	return
}

// Login 登录业务
func Login(user *models.User) (err error) {
	originPassword := user.Password //记录原始密码(用户登录的密码)
	err = db.Model(&models.User{}).Where("phone = ?", user.Phone).First(&user).Error
	//查询数据库出错
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	//用户不存在
	if err == gorm.ErrRecordNotFound {
		return errors.New("用户不存在")
	}
	//判断密码是否正确
	if originPassword != user.Password {
		return errors.New("密码错误")
	}
	return nil

}

// GetUserById 根据id查询用户信息和文章
func GetUserById(id int64) (user *models.User, err error) {
	user = &models.User{}
	err = db.Model(&models.User{}).Where("user_id = ?", id).First(user).Error
	return
}
