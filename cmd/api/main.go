package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ryota1119/gin_webapi/cmd/api/bootstrap"
	"github.com/ryota1119/gin_webapi/internal/handler/controller"
	"github.com/ryota1119/gin_webapi/internal/infrastructure/database"
	"github.com/ryota1119/gin_webapi/internal/infrastructure/jwt_auth"
	"github.com/ryota1119/gin_webapi/internal/infrastructure/redis"
	"github.com/ryota1119/gin_webapi/internal/middleware"
	"github.com/ryota1119/gin_webapi/internal/repository"
	"github.com/ryota1119/gin_webapi/internal/usecase"
)

// @title			Gin WebAPI Example
// @version		1.0
// @description	Ginを使ったWebAPIのSwaggerドキュメント
// @host			localhost:8080
// @BasePath		/
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// データベース初期化
	if err = database.NewDB(); err != nil {
		log.Fatal(err)
	}
	// Redis初期化
	if err = redis.NewRedis(); err != nil {
		log.Fatal(err)
	}

	// jwtAuth Serviceの初期化
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		log.Fatal("JWT_SECRET_KEY is not set in environment variables")
	}
	jwtAuth := jwt_auth.NewJwtAuth(secretKey)

	// 依存性の注入
	// MySQL
	db := database.GetDB()
	// Redis
	redisClient := redis.GetRedisClient()

	// リポジトリ
	authRepo := repository.NewAuthRepository(redisClient)
	postRepo := repository.NewPostRepository(db)
	userRepo := repository.NewUserRepository(db)

	// ユースケース
	authUsecase := usecase.NewAuthUsecase(jwtAuth, authRepo, userRepo)
	postUsecase := usecase.NewPostUsecase(postRepo)

	// コントローラ
	authHandler := controller.NewAuthController(authUsecase)
	postHandler := controller.NewPostController(postUsecase)

	// ミドルウェアのセットアップ
	// 認証ミドルウェア
	authMiddleware := middleware.NewAuthMiddleware(authUsecase)

	// bootstrap初期化
	handler := bootstrap.NewHandler(
		authMiddleware,
		authHandler,
		postHandler,
	)
	//  Routerのセットアップ
	handler.SetupRouter(r)

	// ポート番号8080番で起動
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
