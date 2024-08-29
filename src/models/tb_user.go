package models

import (
	"awesomeProject/src/driver"
	"awesomeProject/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type TbUser struct {
	gorm.Model
	Uuid          *int64 `gorm:"column: uuid; type: bigint unique;" json:"Uuid"`
	Username      string `gorm:"column: username; type: varchar(16);" json:"Username"`
	Password      string `gorm:"column: password; type: varchar(16);" json:"Password"`
	Nickname      string `gorm:"column: nickname; type: varchar(16);" json:"Nickname"`
	PhoneNumber   string `gorm:"column: phone_number; type: varchar(11);" json:"PhoneNumber"`
	Email         string `gorm:"column: email; type: varchar(30);" json:"Email"`
	Gender        string `gorm:"column: gender; type: varchar(2);" json:"Gender"`
	ArticlesNum   *int   `gorm:"column: articles_num; type: int(11);" json:"ArticleNum"`
	DeletedActive *int   `gorm:"column: deleted_active; type: int(1);" json:"DeletedActive"`
}

type QueryUser struct {
	Uuid        *int64
	Username    string
	Password    string
	Nickname    string
	PhoneNumber string
	Email       string
	Gender      int
	ArticlesNum *int
}

type FormUsernameAndPassword struct {
	Username string
	Password string
}

type FormRegister struct {
	Uuid        int64
	Username    string
	Password    string
	Nickname    string
	PhoneNumber string
	Email       string
	Gender      string
}

func (table TbUser) TableName() string {
	return "tb_user"
}

func GetUserInfo(c *gin.Context) {
	db := driver.DB

	var users []QueryUser

	querySQL := `select uuid, username, password, nickname, phone_number, email, gender, articles_num from tb_user;`
	result := db.Raw(querySQL).Debug().Scan(&users)

	log.Println(users)

	if result.Error != nil {
		log.Println("Query Error!", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "Database query failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": users,
	})
}

func LoginHandle(c *gin.Context) {
	db := driver.DB

	webUsername := c.PostForm("username")
	webPassword := c.PostForm("password")

	if len(webUsername) == 0 || len(webPassword) == 0 {
		log.Println(webUsername, webPassword)
		log.Println("The username and password cannot be empty!")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  400,
			"error": "The username and password cannot be empty!",
		})
		return
	} else {
		var usernameAndPassword FormUsernameAndPassword

		queryUsernameAndPassword := `select username, password from tb_user where username = ?;`
		sqlUsernameAndPassword := db.Raw(queryUsernameAndPassword, webUsername).Scan(&usernameAndPassword)

		if sqlUsernameAndPassword.Error != nil {
			log.Println("Query Error!", sqlUsernameAndPassword.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  500,
				"error": "Database query failed",
			})
			return
		}

		if webUsername != usernameAndPassword.Username {
			log.Println(webUsername, webPassword)
			log.Println(usernameAndPassword)
			log.Println("The user does not exist! Please check if your username is correct.")
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  500,
				"error": "The user does not exist! Please check if your username is correct.",
			})
			return
		}

		if webUsername == usernameAndPassword.Username && webPassword != usernameAndPassword.Password {
			log.Println(webUsername, webPassword)
			log.Println(usernameAndPassword)
			log.Println("The password Error!")
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  500,
				"error": "The password error.",
			})
			return
		}

		if webUsername == usernameAndPassword.Username && webPassword == usernameAndPassword.Password {
			log.Println(webUsername, webPassword)
			log.Println(usernameAndPassword)
			log.Println("Login successful!")
			c.JSON(http.StatusOK, gin.H{
				"code":                200,
				"usernameAndPassword": usernameAndPassword,
				"message":             "Login successful!",
			})
			return
		}
		return
	}
}

func RegisterHandle(c *gin.Context) {

	nodeID := int64(1)
	snowFlake := utils.NewSnowflake(nodeID)
	uuid := snowFlake.Generate()

	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	phoneNumber := c.PostForm("phone")
	email := c.PostForm("email")
	gender := c.PostForm("gender")

	db := driver.DB

	var exists int
	querySelectUsername := `select count(*) from tb_user where username = ?`
	err := db.Raw(querySelectUsername, username).Row().Scan(&exists)

	if err != nil {
		panic(err)
	}

	if exists > 0 {
		fmt.Println("Username already exists.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Username already exists.",
		})
	} else {
		insertSQL := `insert into tb_user (uuid, username, password, nickname, phone_number, email, gender) values (?, ?, ?, ?, ?, ?, ?)`
		result := db.Exec(insertSQL, uuid, username, password, nickname, phoneNumber, email, gender)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  500,
				"error": "创建用户失败！",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "User registered successfully.",
		})
		return
	}
}
