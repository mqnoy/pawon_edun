package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	_cookerHttpDelivery "example/web-service-gin/cooker/delivery/http"
	_cookerRepository "example/web-service-gin/cooker/repository/mysql"
	_cookerUcase "example/web-service-gin/cooker/usecase"
	"example/web-service-gin/middleware"
	"example/web-service-gin/model"

	_recipeRepository "example/web-service-gin/recipe/repository/mysql"
)

func init() {
	// TODO: load config here
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	// TODO: inject logger
	// TODO: custom config

	dsn := "root:12345678@tcp(127.0.0.1:3306)/pawon_edun?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gorm_logrus.New(),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   false,
		},
	})

	sqlDB, errSql := db.DB()

	if err != nil || errSql != nil {
		log.Fatalln(err)
	} else {
		log.Info("Database is connected")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto migration
	// TODO: moved this
	db.AutoMigrate(&model.Cooker{})
	db.AutoMigrate(&model.Recipe{})

	g := gin.New()
	g.Use(middleware.LoggingMiddleware())
	g.Use(gin.Recovery())

	// setup global middleware

	// setup repository
	cr := _cookerRepository.NewMysqlCookerRepository(db)
	rr := _recipeRepository.NewMysqlRecipeRepository(db)

	// setup handler
	cu := _cookerUcase.NewCookerUsecase(cr, rr)
	_cookerHttpDelivery.NewCookerHandler(g, cu)

	// pingpong
	g.GET("/api/v1/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("pong"))
	})

	g.Use(middleware.ErrorHandler)
	g.Run("0.0.0.0:8080")

}
