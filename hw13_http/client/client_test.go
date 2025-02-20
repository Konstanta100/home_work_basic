package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Konstanta100/home_work_basic/hw13_http/entity"
)

func TestSendCreateUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, got %s", r.Method)
		}

		if r.URL.EscapedPath() != "/user/create" {
			t.Errorf("Expected URL /user/create, got %s", r.URL.EscapedPath())
		}

		w.WriteHeader(http.StatusCreated)

		var user entity.User
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &user)

		respUser := entity.User{ID: user.ID, Name: user.Name, Age: user.Age}
		resp, _ := json.Marshal(respUser)
		w.Write(resp)
	}))
	defer ts.Close()

	sendCreateUser(ts.URL)
}

func TestSendGetPage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)

		page := entity.Page{}
		resp, _ := json.Marshal(page)
		w.Write(resp)
	}))
	defer ts.Close()

	sendGetPage(ts.URL)
}
