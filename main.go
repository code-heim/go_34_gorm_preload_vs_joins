package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"size:128"`
	Username string `gorm:"size:64"`
	Password string `gorm:"size:255"`
	Notes    []Note // Has many notes
}

type Note struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey"`
	Name    string `gorm:"size:255"`
	Content string `gorm:"type:text"`
	UserID  uint64 `gorm:"index"`
	User    User   // belongs to a user
}

var DB *gorm.DB

func connectDatabase() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Microsecond, // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,             // Disable color
		},
	)

	database, err := gorm.Open(mysql.Open("codeheim:tmp_pwd@tcp(127.0.0.1:3306)/gorm_belongs_to?charset=utf8&parseTime=true"), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func dbMigrate() {
	DB.AutoMigrate(&User{}, &Note{})
}

func main() {
	connectDatabase()
	dbMigrate()

	r := gin.Default()

	r.Use(gin.Logger())

	r.LoadHTMLGlob("views/*")

	r.GET("/", HomeHandler)

	log.Println("Server started!")
	r.Run() // Default Port 8080
}

func HomeHandler(c *gin.Context) {
	var notes []Note
	DB.Preload("User").Find(&notes)

	// Username and the title(name) of the note
	// for a report

	type queryResult struct {
		Name     string
		Username string
	}

	var data []queryResult
	DB.Model(&Note{}).Select("notes.name, users.username").
		Joins("left join users on notes.user_id = users.id").Scan(&data)

	// List of all user's email with the count of notes they have created

	type rawResult struct {
		Username  string
		NoteCount int
	}

	var countResult []rawResult
	DB.Raw("SELECT u.username, COUNT(n.id) as note_count FROM users u JOIN notes n ON n.user_id = u.id GROUP BY 1").Scan(&countResult)

	c.HTML(
		http.StatusOK,
		"index.tpl",
		gin.H{
			"notes":     notes,
			"joinData":  data,
			"noteCount": countResult,
		},
	)
}
