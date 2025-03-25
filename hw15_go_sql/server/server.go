package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/config"
	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/repository"
	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/service"
)

func main() {
	ctx := context.Background()
	conf, err := config.Init()
	if err != nil {
		log.Panicln(err.Error())
	}

	fmt.Println(conf.HTTP.Port, conf.HTTP.Host)

	db, err := NewDB(ctx, conf.DB)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Connected to database")

	repo := repository.New(db)
	mux := http.NewServeMux()
	setupRoutes(mux, repo, db)

	server := &http.Server{
		Addr:              fmt.Sprintf("%v:%v", conf.HTTP.Host, conf.HTTP.Port),
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

func setupRoutes(mux *http.ServeMux, repo repository.Querier, db *pgxpool.Pool) {
	mux.HandleFunc("/orders", allowMethods(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if err := service.Orders(ctx, w, r, repo); err != nil {
				handleError(w, err, http.StatusBadRequest)
			}
		},
		http.MethodGet,
	))

	mux.HandleFunc("/order", allowMethods(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if err := service.Order(ctx, w, r, repo); err != nil {
				handleError(w, err, http.StatusBadRequest)
			}
		},
		http.MethodGet,
	))

	mux.HandleFunc("/order/create", allowMethods(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if err := service.OrderCreateFull(ctx, w, r, db); err != nil {
				handleError(w, err, http.StatusBadRequest)
			}
		},
		http.MethodPost,
	))

	mux.HandleFunc("/product/create", allowMethods(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if err := service.CreateProduct(ctx, w, r, repo); err != nil {
				handleError(w, err, http.StatusBadRequest)
			}
		},
		http.MethodPost,
	))

	mux.HandleFunc("/users", allowMethods(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if err := service.Users(ctx, w, r, repo); err != nil {
				handleError(w, err, http.StatusBadRequest)
			}
		},
		http.MethodGet,
	))

	mux.HandleFunc("/user", allowMethods(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if err := service.User(ctx, w, r, repo); err != nil {
				handleError(w, err, http.StatusBadRequest)
			}
		},
		http.MethodGet,
	))

	mux.HandleFunc("/user/create", allowMethods(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if err := service.CreateUser(ctx, w, r, repo); err != nil {
				handleError(w, err, http.StatusBadRequest)
			}
		},
		http.MethodPost,
	))
}

func allowMethods(next http.HandlerFunc, methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range methods {
			if r.Method == method {
				next.ServeHTTP(w, r)
				return
			}
		}
		handleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
	}
}

func handleError(w http.ResponseWriter, err error, statusCodes ...int) {
	statusCode := http.StatusBadRequest // default
	if len(statusCodes) > 0 {
		statusCode = statusCodes[0]
	}

	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
	log.Println(err)
}
