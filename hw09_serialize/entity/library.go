package entity

type Library struct {
	Books []Book `xml:"entity" bson:"books"`
}
