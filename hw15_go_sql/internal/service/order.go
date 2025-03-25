package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/repository"
	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/trx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Orders(ctx context.Context, w http.ResponseWriter, r *http.Request, repo repository.Querier) error {
	var (
		limit  int64 = 10
		offset int64
		err    error
	)
	limitRaw := r.URL.Query().Get("limit")
	offsetRaw := r.URL.Query().Get("offset")

	if limitRaw != "" {
		limit, err = strconv.ParseInt(limitRaw, 10, 64)
		if err != nil {
			return err
		}
	}
	if offsetRaw != "" {
		offset, err = strconv.ParseInt(offsetRaw, 10, 64)
		if err != nil {
			return err
		}
	}

	params := repository.OrdersParams{
		Limit:  limit,
		Offset: offset,
	}

	res, err := repo.Orders(ctx, params)
	if err != nil {
		return err
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
	return nil
}

func Order(ctx context.Context, w http.ResponseWriter, r *http.Request, repo repository.Querier) error {
	var (
		limit int64 = 1
		err   error
	)

	id := r.URL.Query().Get("id")
	if id == "" {
		return errors.New("missing user id")
	}

	orderID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	params := repository.OrderByIdParams{
		ID:    orderID,
		Limit: limit,
	}

	order, err := repo.OrderById(ctx, params)
	if err != nil {
		return err
	}

	resBody, err := json.Marshal(order)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
	return nil
}

func OrderCreateFull(ctx context.Context, w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) error {
	var err error
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var params trx.OrderFullParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		return err
	}

	if params.UserID == "" {
		return errors.New("missing user id")
	}

	if len(params.Products) == 0 {
		return errors.New("missing products")
	}

	id, err := trx.CreateOrderFull(ctx, params, db)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		fmt.Sprintf(`{"id": "%v"}`, id),
	))
	return nil
}
