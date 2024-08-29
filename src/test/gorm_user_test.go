package test

import (
	"awesomeProject/src/models"
	"awesomeProject/src/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	Uuid        int64
	Username    string
	Password    string
	Nickname    string
	PhoneNumber string
	Email       string
	Gender      int
}

// gorm插入用户数据测试单元

func TestGormUserTest(t *testing.T) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	nodeID := int64(1)
	snowFlake := utils.NewSnowflake(nodeID)

	// 生成一个新的雪花id
	uuid := snowFlake.Generate()

	// 使用SQL语句插入用户数据
	userInsertSQL := `insert into tb_user(uuid, username, password, nickname, phone_number, email, gender) values (?, ?, ?, ?, ?, ?, ?);`
	user := User{
		Uuid:        uuid,
		Username:    "admin",
		Password:    "123456",
		Nickname:    "管理员",
		PhoneNumber: "1129391231",
		Email:       "1465439890@qq.com",
		Gender:      1,
	}
	result := db.Exec(userInsertSQL, user.Uuid, user.Username, user.Password, user.Nickname, user.PhoneNumber, user.Email, user.Gender)
	if result.Error != nil {
		t.Fatal(result.Error)
	}

	var users []models.TbUser

	querySQL := `Select * from tb_user;`

	rows := db.Raw(querySQL).Scan(&users)
	if rows.Error != nil {
		t.Fatal(rows.Error)
	}

	if len(users) == 0 {
		t.Error("Expected to find at least one user, but found none")
	}

	for _, user := range users {
		fmt.Println(user)
		t.Log(user)
	}

}
