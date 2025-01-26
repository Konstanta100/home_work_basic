package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/Konstanta100/home_work_basic/hw09_serialize/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"gopkg.in/yaml.v3"
)

func main() {
	books := []entity.Book{
		{
			ID:     1,
			Title:  "Война и мир",
			Author: "Лев Николаевич Толстой",
			Year:   1867,
			Size:   20000,
			Rate:   4.2,
			Sample: []byte("Что это? я падаю! у меня ноги подкашиваются"),
		},
		{
			ID:     2,
			Title:  "Библия",
			Author: "Бог",
			Size:   10000,
			Rate:   4.8,
			Sample: []byte("Просите, и дано будет вам; ищите и найдёте; стучите, и отворят вам ..."),
		},
	}

	result, err := SerializeXML(books)
	if err != nil {
		fmt.Println(err)
	}

	books, err = DeserializeXML(result)
	if err != nil {
		fmt.Println(err)
	}

	result, err = SerializeJSON(books)
	if err != nil {
		fmt.Println(err)
	}

	books, err = DeserializeJSON(result)
	if err != nil {
		fmt.Println(err)
	}

	result, err = SerializeYAML(books)
	if err != nil {
		fmt.Println(err)
	}

	books, err = DeserializeYAML(result)
	if err != nil {
		fmt.Println(err)
	}

	result, err = SerializeGob(books)
	if err != nil {
		fmt.Println(err)
	}

	books, err = DeserializeGob(result)
	if err != nil {
		fmt.Println(err)
	}

	result, err = SerializeBson(books)
	if err != nil {
		fmt.Println(err)
	}

	books, err = DeserializeBson(result)
	if err != nil {
		fmt.Println(err)
	}

	for _, book := range books {
		test := book.String()
		fmt.Println(test)
	}
}

func SerializeXML(books []entity.Book) ([]byte, error) {
	library := entity.Library{Books: books}

	return xml.Marshal(library)
}

func DeserializeXML(data []byte) ([]entity.Book, error) {
	var library entity.Library
	err := xml.Unmarshal(data, &library)

	return library.Books, err
}

func SerializeJSON(books []entity.Book) ([]byte, error) {
	return json.Marshal(books)
}

func DeserializeJSON(data []byte) ([]entity.Book, error) {
	var books []entity.Book
	err := json.Unmarshal(data, &books)

	return books, err
}

func SerializeYAML(books []entity.Book) ([]byte, error) {
	return yaml.Marshal(books)
}

func DeserializeYAML(data []byte) ([]entity.Book, error) {
	var books []entity.Book
	err := yaml.Unmarshal(data, &books)

	return books, err
}

func SerializeGob(books []entity.Book) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(books)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func DeserializeGob(data []byte) ([]entity.Book, error) {
	var books []entity.Book
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)

	err := decoder.Decode(&books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func SerializeBson(books []entity.Book) ([]byte, error) {
	return bson.Marshal(entity.Library{Books: books})
}

func DeserializeBson(data []byte) ([]entity.Book, error) {
	var library entity.Library
	err := bson.Unmarshal(data, &library)

	return library.Books, err
}
