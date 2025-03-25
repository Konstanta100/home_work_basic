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
	"github.com/jackc/pgx/v5/pgtype"
)

type UserDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func convertUserToDto(user *repository.User) UserDto {
	var userDto UserDto
	uuidValue, err := user.ID.Value()
	if err != nil {
		return userDto
	}
	return UserDto{ID: uuidValue.(string), Name: user.Name, Email: user.Email}
}

func Users(ctx context.Context, w http.ResponseWriter, r *http.Request, repo repository.Querier) error {
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

	params := repository.UsersParams{
		Limit:  limit,
		Offset: offset,
	}

	users, err := repo.Users(ctx, params)
	if err != nil {
		return err
	}

	convertedUsers := make([]UserDto, 0, len(users))
	for _, user := range users {
		userDto := convertUserToDto(user)
		convertedUsers = append(convertedUsers, userDto)
	}

	resBody, err := json.Marshal(convertedUsers)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
	return nil
}

func User(ctx context.Context, w http.ResponseWriter, r *http.Request, repo repository.Querier) error {
	var (
		limit int64 = 1
		err   error
	)

	var userID pgtype.UUID
	id := r.URL.Query().Get("id")
	if id == "" {
		return errors.New("missing user id")
	}

	err = userID.Scan(id)
	if err != nil {
		return errors.New("invalid user id")
	}

	params := repository.UserByIdParams{
		ID:    userID,
		Limit: limit,
	}

	user, err := repo.UserById(ctx, params)
	if err != nil {
		return err
	}

	res := convertUserToDto(user)
	resBody, err := json.Marshal(res)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
	return nil
}

func CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request, repo repository.Querier) error {
	var err error
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var params repository.UserCreateParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		return err
	}
	if params.Name == "" {
		return errors.New("name is required")
	}
	if params.Email == "" {
		return errors.New("email is required")
	}
	if params.Password == "" {
		return errors.New("password is required")
	}

	uuid, err := repo.UserCreate(ctx, params)
	if err != nil {
		return err
	}
	if !uuid.Valid {
		return errors.New("invalid uuid of new user")
	}

	uuidValue, err := uuid.Value()
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		fmt.Sprintf(`{"id": "%v"}`, uuidValue.(string)),
	))
	return nil
}
