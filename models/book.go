package models

type Book struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Year   string `json:"year" bson:"year"`
	Code   string `json:"code" bson:"code" validate:"required,unique"` // Add the "code" field with not null and unique constraints
}
