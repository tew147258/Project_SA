package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tew147258/app/controllers"
	_ "github.com/tew147258/app/docs"
	"github.com/tew147258/app/ent"
)

type Users struct {
	User []User
}

type User struct {
	Email    string
	Password string
	Name     string
	Birthday string
	Tel      string
}

type Stadiums struct {
	Stadium []Stadium
}

type Stadium struct {
	Namestadium string
}

type Borrows struct {
	Borrow []Borrow
}

type Borrow struct {
	Type string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewStadiumController(v1, client)
	controllers.NewBorrowController(v1, client)
	controllers.NewConfirmationController(v1, client)

	//Set User Data
	users := Users{
		User: []User{
			User{"chanwit@gmail.com", "zxc123", "Chanwit Kaewkasi", "2006-01-02T15:04:05Z07:00", "0832642968"},
		},
	}
	for _, u := range users.User {
		time, _ := time.Parse(time.RFC3339, u.Birthday)
		client.User.
			Create().
			SetEmail(u.Email).
			SetPassword(u.Password).
			SetName(u.Name).
			SetBirthday(time).
			SetTel(u.Tel).
			Save(context.Background())
	}

	//Set Stadium Data
	stadiums := Stadiums{
		Stadium: []Stadium{
			Stadium{"สนามบาส"},
			Stadium{"สนามวอลเลย์บอล"},
			Stadium{"สนามฟุตบอล"},
		},
	}
	for _, s := range stadiums.Stadium {
		client.Stadium.
			Create().
			SetNamestadium(s.Namestadium).
			Save(context.Background())
	}

	//Set Borrow Data
	borrows := Borrows{
		Borrow: []Borrow{
			Borrow{"จะยืมอุปกรณ์"},
			Borrow{"ไม่ยืมอุปกรณ์"},
		},
	}
	for _, b := range borrows.Borrow {
		client.Borrow.
			Create().
			SetType(b.Type).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
