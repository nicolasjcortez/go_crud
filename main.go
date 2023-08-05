package main

import (
	"context"
	"fmt"
	"go_crud/controllers"
	"go_crud/docs"
	"go_crud/routes"
	"go_crud/services"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server *gin.Engine
	ctx    context.Context

	mongoclient *mongo.Client

	userService         services.UserService
	UserController      controllers.UserController
	userCollection      *mongo.Collection
	UserRouteController routes.UserRouteController
)

func init() {

	ctx = context.Background()

	// Connect to MongoDB
	DBUri := os.Getenv("GO_CRUD_MONGO_URI")
	mongoconn := options.Client().ApplyURI(DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// 👇 Instantiate the Constructors
	userCollection = mongoclient.Database("go_crud").Collection("users")
	userService = services.NewUserService(userCollection, ctx)
	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewUserControllerRoute(UserController)

	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	UserRouteController.UserRoute(router)

	// SWAGGER
	docs.SwaggerInfo.Title = "Users API"
	docs.SwaggerInfo.Description = "Users API"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run(":" + os.Getenv("PORT")))
}
