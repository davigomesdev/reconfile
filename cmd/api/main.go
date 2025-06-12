package main

import (
	"time"

	"github.com/davigomesdev/reconfile/internal/adapters/handlers"
	"github.com/davigomesdev/reconfile/internal/adapters/interceptors"
	"github.com/davigomesdev/reconfile/internal/application/providers"
	authUC "github.com/davigomesdev/reconfile/internal/application/usecases/auth"
	supplierUC "github.com/davigomesdev/reconfile/internal/application/usecases/supplier"
	userUC "github.com/davigomesdev/reconfile/internal/application/usecases/user"
	"github.com/davigomesdev/reconfile/internal/infrastructure/database"
	db_repositories "github.com/davigomesdev/reconfile/internal/infrastructure/database/repositories"
	env_config "github.com/davigomesdev/reconfile/internal/infrastructure/env-config"
	jwt_auth "github.com/davigomesdev/reconfile/internal/infrastructure/jwt-auth"
	"github.com/davigomesdev/reconfile/internal/infrastructure/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env := env_config.LoadConfig()

	db := database.ConnectDB()
	defer database.DisconnectDB()

	database.RunMigrations()

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	app.Use(interceptors.ErrorFilter())

	jwtAuthService := jwt_auth.NewJWTAuthService(env)

	hashProvider := providers.NewHashProvider()
	xlsxParserProvider := providers.NewXLSXParserProvider()

	userRepository := db_repositories.NewUserRepository(db)
	supplierRepository := db_repositories.NewSupplierRepository(db)

	signinUseCase := authUC.NewSignInUseCase(userRepository, jwtAuthService, hashProvider)
	signupUseCase := authUC.NewSignUpUseCase(userRepository, jwtAuthService, hashProvider)
	refreshTokensUseCase := authUC.NewRefreshTokensUseCase(userRepository, jwtAuthService)

	getUserUseCase := userUC.NewGetUserUseCase(userRepository)
	searchUserUseCase := userUC.NewSearchUserUseCase(userRepository)
	createUserUseCase := userUC.NewCreateUserUseCase(userRepository, hashProvider)
	updateUserUseCase := userUC.NewUpdateUserUseCase(userRepository)
	updatePasswordUserUseCase := userUC.NewUpdatePasswordUserUseCase(userRepository, hashProvider)
	deleteUserUseCase := userUC.NewDeleteUserUseCase(userRepository)

	getSupplierUseCase := supplierUC.NewGetSupplierUseCase(supplierRepository)
	searchSupplierUseCase := supplierUC.NewSearchSupplierUseCase(supplierRepository)
	overviewSupplierUseCase := supplierUC.NewOverviewSupplierUseCase(supplierRepository)
	importSuppliersUseCase := supplierUC.NewImportSuppliersUseCase(supplierRepository, xlsxParserProvider)

	authHandler := handlers.NewAuthHandler(signinUseCase, signupUseCase, refreshTokensUseCase)
	userHandler := handlers.NewUserHandler(getUserUseCase, searchUserUseCase, createUserUseCase, updateUserUseCase, updatePasswordUserUseCase, deleteUserUseCase)
	supplierHandler := handlers.NewSupplierHandler(getSupplierUseCase, searchSupplierUseCase, overviewSupplierUseCase, importSuppliersUseCase)

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/signin", authHandler.SignIn)
	auth.POST("/signup", authHandler.SignUp)
	auth.POST("/refresh", authHandler.RefreshTokens)

	users := api.Group("/users", middlewares.JWTGuard(jwtAuthService))
	users.GET("/:id", userHandler.Get)
	users.GET("/current", userHandler.Current)
	users.GET("", userHandler.Search)
	users.POST("", userHandler.Create)
	users.PUT("/:id", userHandler.Update)
	users.PUT("/current", userHandler.UpdateCurrent)
	users.PATCH("/password", userHandler.UpdatePassword)
	users.DELETE("/:id", userHandler.Delete)

	suppliers := api.Group("/suppliers", middlewares.JWTGuard(jwtAuthService))
	suppliers.GET("/:id", supplierHandler.Get)
	suppliers.GET("", supplierHandler.Search)
	suppliers.GET("/overview", supplierHandler.Overview)
	suppliers.POST("/import", supplierHandler.Import)

	app.Run(":" + env.Port)
}
