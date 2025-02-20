package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Konstanta100/home_work_basic/hw13_http/entity"
)

const HTTP = "http://"

func main() {
	address := flag.String("address", "127.0.0.1:8888", "Адрес сайта")
	resource := flag.String("resource", "/about", "Ресурс на сайте")
	flag.Parse()

	sendGetPage(HTTP + *address + *resource)
	sendCreateUser(HTTP + *address)
}

func sendCreateUser(address string) {
	user := entity.User{
		ID:   1,
		Name: "Test",
		Age:  26,
	}

	resp, err := http.Post( //nolint:noctx
		address+"/user/create",
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	var newUser entity.User

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}

	fmt.Printf("User: %+v\n", newUser)
}

func sendGetPage(address string) {
	resp, err := http.Get(address) //nolint:gosec, noctx
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

	page := entity.Page{}

	err = json.Unmarshal(body, &page)
	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}

	fmt.Printf("Page: %+v\n", page)
}
