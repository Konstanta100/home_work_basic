package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Konstanta100/home_work_basic/hw13_http/entity"
)

func TestMiddleware(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMiddlewareMethodNotAllowed(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "HEAD", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestAbout(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(about)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var page entity.Page
	err = json.NewDecoder(rr.Body).Decode(&page)
	if err != nil {
		t.Fatal(err)
	}

	if page.Name != "About" {
		t.Errorf("handler returned unexpected body: got %v want %v", page.Name, "About")
	}
}

func TestAboutMethodNotAllowed(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "POST", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(about)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestCreateUser(t *testing.T) {
	user := entity.User{ID: 1, Name: "John Doe", Age: 30}
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "POST", "/user/create", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var returnedUser entity.User
	err = json.NewDecoder(rr.Body).Decode(&returnedUser)
	if err != nil {
		t.Fatal(err)
	}

	if returnedUser != user {
		t.Errorf("handler returned unexpected body: got %v want %v", returnedUser, user)
	}
}

func TestCreateUserMethodNotAllowed(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", "/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
