package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Phone    int                `bson:"phone" json:"phone"`
	Street   string             `bson:"street" json:"street"`
	Ward     string             `bson:"ward" json:"ward"`
	District string             `bson:"district" json:"district"`
	City     string             `bson:"city" json:"city"`
}
