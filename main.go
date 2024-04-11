package main

import (
	"context"
	"fmt"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test/src/config"
	"test/src/constructor"
	"test/src/middleware"
	"test/src/repository/mssql"
	testHTTPRepo "test/src/rpc/http"
	httptransport "test/src/transport/http"
	corsutil "test/src/utils/cors"
	mssqlConn "test/src/utils/database/mssql"
	healthcheckutil "test/src/utils/healthcheck"
	liblogger "test/src/utils/logger"
)

func main() {
	// main ctx
	ctx := context.Background()

	// init structured logger for the service
	logger := liblogger.NewServiceLogger("auth-api")
	_ = level.Info(logger).Log("msg", "service started")

	// init service configuration
	err := config.InitConfigs()
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// init mssql connection
	microsoftSQLConn, err := mssqlConn.InitConnect(
		config.MainConfig.MicrosoftSQLConfig.MSSQLHost,
		config.MainConfig.MicrosoftSQLConfig.MSSQLUser,
		config.MainConfig.MicrosoftSQLConfig.MSSQLPass,
		config.MainConfig.MicrosoftSQLConfig.MSSQLDB,
		config.MainConfig.MicrosoftSQLConfig.MSSQLPort,
	)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	mainRepo, err := mssql.NewStore(ctx, microsoftSQLConn, logger)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	httpClient := testHTTPRepo.InitHTTPClient()
	rpc := testHTTPRepo.NewRPC(httpClient, logger)

	svc := constructor.NewService(mainRepo, rpc, logger)

	// init endpoints (endpoints layer)
	endpoints := middleware.MakeEndpoints(svc)

	// init HTTP handler (transport layer)
	serverOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(httptransport.EncodeErrorResponse),
	}
	handler := httptransport.NewHTTPService(endpoints, serverOptions, logger)

	// add routes, prometheus and health check handlers
	http.Handle("/currency/", corsutil.CORS(handler))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/check", healthcheckutil.HealthCheck)

	// init errors chan and ticker
	errs := make(chan error)

	// make chan for syscall
	go func() {
		c := make(chan os.Signal, config.DefaultChanCapacity)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// init HTTP server
	go func() {
		_ = level.Info(logger).Log("transport", "HTTP", "port", ":"+config.MainConfig.Port)
		errs <- http.ListenAndServe(config.MainConfig.Port, nil)
	}()

	defer func() {
		_ = level.Info(logger).Log("msg", "service ended")
	}()

	_ = level.Error(logger).Log("exit", <-errs)
}
