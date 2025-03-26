package trx

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/Konstanta100/home_work_basic/hw15_go_sql/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateOrderFull(ctx context.Context, params OrderFullParams, db *pgxpool.Pool) (int64, error) {
	var id int64
	tx, err := db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		return id, err
	}
	defer tx.Rollback(ctx)

	dbTX := repository.Queries{}
	repo := dbTX.WithTx(tx)

	var userID pgtype.UUID
	err = userID.Scan(params.UserID)
	if err != nil {
		return id, errors.New("invalid user id")
	}

	paramsUser := repository.UserByIdParams{ID: userID, Limit: 1}
	user, err := repo.UserById(ctx, paramsUser)
	if err != nil {
		return id, errors.New("user not found")
	}

	productIDs := make([]int64, 0, len(params.Products))
	for _, product := range params.Products {
		productIDs = append(productIDs, product.ID)
	}

	paramsProducts := repository.ProductsByIdsParams{Column1: productIDs, Limit: 1000}
	products, err := repo.ProductsByIds(ctx, paramsProducts)
	if len(products) == 0 || err != nil {
		return id, errors.New("products not found")
	}

	var totalAmount pgtype.Numeric
	totalAmount, err = calculateTotalAmount(products, params)
	if err != nil {
		return id, err
	}

	orderParam := repository.OrderCreateParams{UserID: user.ID, TotalAmount: totalAmount}
	id, err = repo.OrderCreate(ctx, orderParam)
	if err != nil {
		return id, errors.New("order not create")
	}

	for _, product := range params.Products {
		orderID := id
		productID := product.ID
		productParams := repository.LinkOrderToProductCreateParams{
			OrderID:   &orderID,
			ProductID: &productID,
			Count:     product.Count,
		}
		_, err = repo.LinkOrderToProductCreate(ctx, productParams)
		if err != nil {
			return id, errors.New("link product to order not create")
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return id, err
	}

	return id, nil
}

func calculateTotalAmount(products []*repository.Product, params OrderFullParams) (pgtype.Numeric, error) {
	var totalAmount pgtype.Numeric
	totalAmount.Int = big.NewInt(0)
	totalAmount.Exp = -2
	totalAmount.Valid = true

	for _, product := range products {
		for _, paramProduct := range params.Products {
			if paramProduct.ID == product.ID {
				countProduct := pgtype.Numeric{}
				countProduct.Int = big.NewInt(paramProduct.Count)
				countProduct.Exp = 0
				countProduct.Valid = true

				if !product.Price.Valid || !countProduct.Valid {
					return pgtype.Numeric{}, fmt.Errorf("цена или количество товара равно NULL")
				}

				productTotal, err := multiplyNumeric(product.Price, countProduct)
				if err != nil {
					return pgtype.Numeric{}, fmt.Errorf("ошибка при умножении: %w", err)
				}

				totalAmount, err = addNumeric(totalAmount, productTotal)
				if err != nil {
					return pgtype.Numeric{}, fmt.Errorf("ошибка при сложении: %w", err)
				}
			}
		}
	}

	return totalAmount, nil
}

func multiplyNumeric(a, b pgtype.Numeric) (pgtype.Numeric, error) {
	if !a.Valid || !b.Valid {
		return pgtype.Numeric{}, fmt.Errorf("один из аргументов равен NULL")
	}

	result := pgtype.Numeric{}
	result.Int = new(big.Int).Mul(a.Int, b.Int)
	result.Exp = a.Exp + b.Exp
	result.Valid = true

	return result, nil
}

func addNumeric(a, b pgtype.Numeric) (pgtype.Numeric, error) {
	if !a.Valid || !b.Valid {
		return pgtype.Numeric{}, fmt.Errorf("один из аргументов равен NULL")
	}

	if a.Exp > b.Exp {
		b.Int.Mul(b.Int, big.NewInt(10).Exp(big.NewInt(10), big.NewInt(int64(a.Exp-b.Exp)), nil))
		b.Exp = a.Exp
	} else if a.Exp < b.Exp {
		a.Int.Mul(a.Int, big.NewInt(10).Exp(big.NewInt(10), big.NewInt(int64(b.Exp-a.Exp)), nil))
		a.Exp = b.Exp
	}

	result := pgtype.Numeric{}
	result.Int = new(big.Int).Add(a.Int, b.Int)
	result.Exp = a.Exp
	result.Valid = true

	return result, nil
}
