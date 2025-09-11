package main

import (
	"context"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rasteiro11/MCABankCustomer/entities"
	pbAuthClient "github.com/rasteiro11/MCABankCustomer/gen/proto/go"
	"github.com/rasteiro11/MCABankCustomer/src/customer/delivery/http"
	"github.com/rasteiro11/MCABankCustomer/src/customer/delivery/http/middleware"
	customerRepo "github.com/rasteiro11/MCABankCustomer/src/customer/repository"
	customerService "github.com/rasteiro11/MCABankCustomer/src/customer/service"
	"github.com/rasteiro11/PogCore/pkg/config"
	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogCore/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	dbInstance, err := database.NewDatabase(database.GetMysqlEngineBuilder)
	if err != nil {
		logger.Of(ctx).Fatalf("[main] database.NewDatabase() returned error: %+v\n", err)
	}

	if err := dbInstance.Migrate(entities.GetEntities()...); err != nil {
		logger.Of(ctx).Fatalf("[main] database.Migrate() returned error: %+v\n", err)
	}

	db := dbInstance.Conn()

	customerRepo := customerRepo.NewCustomerRepository(db)

	customerSvc := customerService.NewCustomerService(customerRepo)

	credentials := insecure.NewCredentials()
	authConn, err := grpc.Dial(config.Instance().RequiredString("AUTH_GRPC_SERVICE"),
		grpc.WithTransportCredentials(credentials))
	if err != nil {
		logger.Of(ctx).Fatalf(
			"[main] grpc.Dial returned error: err=%+v", err)
	}

	authClient := pbAuthClient.NewAuthServiceClient(authConn)

	app := server.NewServer()
	app.Use("/customers", middleware.ValidateUserMiddleware(authClient))
	app.Use("/*", cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	http.NewHandler(app, http.WithCustomerService(customerSvc))

	app.PrintRouter()

	port := config.Instance().RequiredString("SERVER_PORT")
	if err := app.Start(port); err != nil {
		logger.Of(ctx).Fatalf("[main] server.Start() returned error: %+v\n", err)
	}
}
