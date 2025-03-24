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
)

func Products(ctx context.Context, w http.ResponseWriter, r *http.Request, repo repository.Querier) error {
	var (
		limit  int64 = 10
		offset int64 = 0
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

	params := repository.ProductsParams{
		Limit:  limit,
		Offset: offset,
	}

	res, err := repo.Products(ctx, params)
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

func CreateProduct(ctx context.Context, w http.ResponseWriter, r *http.Request, repo repository.Querier) error {
	var err error
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var params repository.ProductCreateParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		return err
	}
	if params.Name == "" {
		return errors.New("name is required")
	}

	if !params.Price.Valid {
		return errors.New("missing user id")
	}

	id, err := repo.ProductCreate(ctx, params)
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
