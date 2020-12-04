package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"disbursement-service/config"
	"disbursement-service/migration"
	"disbursement-service/model"
	"disbursement-service/repository/flip"
	"disbursement-service/repository/postgres"
	transportHTTP "disbursement-service/transport/http"
	"disbursement-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func initDB(config *model.Postgres) *sqlx.DB {
	// sql.Open("postgres", "user=test password=test dbname=test sslmode=disable")
	var err error
	connection := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass, config.Name)

	// val := url.Values{}
	// val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := sqlx.Connect(`postgres`, connection)
	if err != nil {
		log.Panic("error connect :", err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}

func main() {
	config, err := config.NewConfig()
	if nil != err {
		panic(err)
	}

	httpAddr := flag.String("http.addr", config.AppPort, "HTTP listen address")

	// setting database
	db := initDB(config.DB)
	dbConn := postgres.NewDisbursementRepository(db)

	errChan := make(chan error)
	// starting migration
	go func() {
		connection := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Pass, config.DB.Name)
		migration.RunMigrations(connection)
	}()

	// setting flip
	client := &http.Client{}
	flipRepo := flip.NewDisbursementAPI(config.Flip.Host, config.Flip.Authorization, client)

	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	uc := usecase.NewDisbursement(dbConn, flipRepo, logger)

	// start usecase
	go func() {
		log.Println("[HTTP][Info] Starting APP DISBURSEMENT HTTP")
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		mux := http.NewServeMux()
		ctx := context.Background()
		mux.Handle("/", transportHTTP.MakeHandler(ctx, uc, logger))
		errChan <- http.ListenAndServe(*httpAddr, accessControl(mux))
	}()

	err = <-errChan
	if nil != err {
		log.Println("error: ", err)
		panic(err)
	}
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, OPTIONS")
		r.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, scope, state, hd, code")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, scope, state, hd, code")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
