package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	FirstName string               `bson:"first_name" json:"first_name"`
	LastName  string               `bson:"last_name" json:"last_name"`
	Email     string               `bson:"email" json:"email"`
	Address   []primitive.ObjectID `bson:"address" json:"address"`
	Password  string               `bson:"password" json:"password"`
	Receipt   []primitive.ObjectID `bson:"receipt" json:"receipt"`
	RoleAdmin bool                 `bson:"role_admin" json:"role_admin"`
}
