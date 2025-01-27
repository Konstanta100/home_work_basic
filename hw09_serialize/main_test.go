package main

import (
	"bytes"
	"testing"

	"github.com/Konstanta100/home_work_basic/hw09_serialize/entity"
)

func TestSerialization(t *testing.T) {
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

	xmlBooks, err := SerializeXML(books)
	if err != nil {
		t.Errorf("XML сериализация не прошла: %v ", err)
	}

	deserializedBooks, err := DeserializeXML(xmlBooks)

	if err != nil || !compareBooks(books, deserializedBooks) {
		t.Errorf("XML десериализация не прошла: %v ", err)
	}

	jsonBooks, err := SerializeJSON(books)
	if err != nil {
		t.Errorf("Json сериализация не прошла: %v ", err)
	}

	deserializedBooks, err = DeserializeJSON(jsonBooks)

	if err != nil || !compareBooks(books, deserializedBooks) {
		t.Errorf("Json десериализация не прошла: %v ", err)
	}

	yamlBooks, err := SerializeYAML(books)
	if err != nil {
		t.Errorf("Yaml сериализация не прошла: %v ", err)
	}

	deserializedBooks, err = DeserializeYAML(yamlBooks)

	if err != nil || !compareBooks(books, deserializedBooks) {
		t.Errorf("Yaml десериализация не прошла: %v ", err)
	}

	gobBooks, err := SerializeGob(books)
	if err != nil {
		t.Errorf("Gob сериализация не прошла: %v ", err)
	}

	deserializedBooks, err = DeserializeGob(gobBooks)

	if err != nil || !compareBooks(books, deserializedBooks) {
		t.Errorf("Gob десериализация не прошла: %v ", err)
	}

	bsonBooks, err := SerializeBson(books)
	if err != nil {
		t.Errorf("Bson сериализация не прошла: %v ", err)
	}

	deserializedBooks, err = DeserializeBson(bsonBooks)

	if err != nil || !compareBooks(books, deserializedBooks) {
		t.Errorf("Bson десериализация не прошла: %v ", err)
	}
}

func compareBooks(originalBooks, migratedBooks []entity.Book) bool {
	if len(originalBooks) != len(migratedBooks) {
		return false
	}

	for i, origin := range originalBooks {
		migrate := migratedBooks[i]

		if origin.Title != migrate.Title || origin.Year != migrate.Year || origin.Size != migrate.Size ||
			origin.Rate != migrate.Rate || origin.ID != migrate.ID || origin.Author != migrate.Author {
			return false
		}

		if !bytes.Equal(origin.Sample, migrate.Sample) {
			return false
		}
	}

	return true
}
