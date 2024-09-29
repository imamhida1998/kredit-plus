package main

import (
	"fmt"
	"kredit-plus/config"
	"kredit-plus/lib/db"
	"kredit-plus/service/handler"
	"kredit-plus/service/repo"
	"kredit-plus/service/usecase"
	"os"
	"sync"

	redisCluster "github.com/go-redis/redis/v8"
)

func main() {
	app := fiber.New()
	env := config.New()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := db.InitoDatabase()
		if err != nil {
			fmt.Println("MySQL :", err.Error())
			os.Exit(1)
		}
	}()

	wg.Wait()
	fmt.Println("DB  :", db.EngineSQL.DataSourceName())

	store := redisCluster.NewClient(&redisCluster.Options{
		Addr: env.Get("REDIS_URL") + ":" + env.Get("REDIS_PORT"),
	})

	repoUser := repo.NewUserRepository()
	repoTransaction := repo.NewTransactionRepository()
	repoRedis := repo.NewRedisRepository(store)

	UseCaseCustomer := usecase.NewCustomerService(repoUser)
	UsecaseTransaction := usecase.NewTransaction(repoUser, repoTransaction, repoRedis)

	NewHandler := handler.NewHandler(UseCaseCustomer, UsecaseTransaction)

	NewHandler.Route(app)

	app.Listen(":3000")

}
