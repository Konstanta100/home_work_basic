package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const HTTP = "http"

func main() {
	address := flag.String("address", "127.0.0.1:8888", "Адрес сайта")
	resource := flag.String("resource", "/about", "Ресурс на сайте")
	flag.Parse()

	resp, err := http.Get(getURLAddress(*address, *resource)) //nolint:noctx
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка HTTP-ответа: %d %s\n", resp.StatusCode, resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	page := Page{}

	err = json.Unmarshal(body, &page)
	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}

	fmt.Printf("Page: %+v\n", page)

	user := User{
		ID:   1,
		Name: "Test",
		Age:  26,
	}

	resp, err = http.Post( //nolint:noctx
		getURLAddress(*address, "/user/create"),
		"application/json",
		strings.NewReader(user.String()),
	)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}

	fmt.Println("User created successfully")

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	var newUser User

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}

	fmt.Printf("User: %+v\n", newUser)
}

func getURLAddress(host string, res string) string {
	return HTTP + "://" + host + res
}

type Page struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Host        string `json:"host"`
	Resource    string `json:"resource"`
	Method      string `json:"method"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) String() string {
	body, _ := json.Marshal(u)
	return string(body)
}
