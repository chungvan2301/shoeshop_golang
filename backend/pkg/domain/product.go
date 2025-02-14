package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Brand         string             `bson:"brand" json:"brand"`
	Gender        string             `bson:"gender" json:"gender"`
	Category      string             `bson:"category" json:"category"`
	Price         float64            `bson:"price" json:"price"`
	IsInInventory bool               `bson:"is_in_inventory" json:"is_in_inventory"`
	ItemsLeft     int                `bson:"items_left" json:"items_left"`
	ImageURL      string             `bson:"imageURL" json:"imageURL"`
	Slug          string             `bson:"slug" json:"slug"`
}
