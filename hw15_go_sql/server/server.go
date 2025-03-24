package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/config"
	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/repository"
	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/service"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	conf, err := config.Init()
	if err != nil {
		log.Panicln(err.Error())
	}

	db, err := NewDB(ctx, conf.DB)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Connected to database")

	repo := repository.New(db)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /orders", func(w http.ResponseWriter, r *http.Request) {
		if err = service.Orders(ctx, w, r, repo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	mux.HandleFunc("GET /order", func(w http.ResponseWriter, r *http.Request) {
		if err = service.Order(ctx, w, r, repo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /order/create", func(w http.ResponseWriter, r *http.Request) {
		if err = service.OrderCreateFull(ctx, w, r, db); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	mux.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		if err = service.Products(ctx, w, r, repo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /product/create", func(w http.ResponseWriter, r *http.Request) {
		if err = service.CreateProduct(ctx, w, r, repo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		if err = service.Users(ctx, w, r, repo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	mux.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
		if err = service.User(ctx, w, r, repo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /user/create", func(w http.ResponseWriter, r *http.Request) {
		if err = service.CreateUser(ctx, w, r, repo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", conf.HTTP.Port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	if err = server.ListenAndServe(); err != nil {
		log.Fatalln("Error listen server")
	}

	fmt.Println("Server start")
}

func NewDB(ctx context.Context, dbCfg config.DB) (*pgxpool.Pool, error) {
	fmt.Println(dbCfg.User, dbCfg.Password, net.JoinHostPort(dbCfg.Host, strconv.Itoa(dbCfg.Port)), dbCfg.Database)
	connConfig, err := pgx.ParseConfig(
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?TimeZone=Europe/Moscow",
			dbCfg.User,
			dbCfg.Password,
			net.JoinHostPort(dbCfg.Host, strconv.Itoa(dbCfg.Port)),
			dbCfg.Database,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create DSN for DB connection: %w", err)
	}
	dbc, err := pgxpool.New(ctx, connConfig.ConnString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB : %w", err)
	}
	if err = dbc.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return dbc, nil
}
