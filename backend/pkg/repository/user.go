package repository

import (
	"context"
	"log"

	"github.com/chungvan2301/shoeshop/backend/pkg/domain"
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserRepository(userCollection *mongo.Collection, ctx context.Context) *UserRepository {
	return &UserRepository{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserRepository) GetUserDetail(userID string) (models.UserDetail, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Format user ID fail!")
	}

	filter := bson.M{"_id": objID}

	var userDetail models.UserDetail

	err = u.userCollection.FindOne(u.ctx, filter).Decode(&userDetail)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("User not found!")
			return models.UserDetail{}, err
		}
		log.Println("Error retrieving user:", err)
		return models.UserDetail{}, err
	}

	return userDetail, nil
}

func (u *UserRepository) RegisterUser(user models.UserInput) error {

	newUser := domain.User{
		ID:        primitive.NewObjectID(),
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
		RoleAdmin: false,
	}

	// Thực hiện thao tác InsertOne để thêm vào MongoDB
	_, err := u.userCollection.InsertOne(u.ctx, newUser)
	if err != nil {
		log.Println("Lỗi khi thêm user: ", err)
		return err
	}

	log.Println("Thêm user thành công: ", newUser)
	return nil
}

func (u *UserRepository) EditUser(user models.UserUpdate) error {

	objID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		log.Println("Lỗi khi update user: ", err)
	}

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"password":   user.NewPassword,
		},
	}

	result, err := u.userCollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		log.Println("user không tìm thấy")
	}

	log.Println("Update user thành công: ", user)
	return nil
}

func (u *UserRepository) GetUserPassword(userID string) (string, error) {

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Format ID error!")
	}

	var user domain.User
	err = u.userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		}
		log.Println("Error retrieving user: ", err)
		return "", err
	}

	return user.Password, nil
}

func (u *UserRepository) DeleteUser(userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Format user ID fail!")
	}

	filter := bson.M{"_id": objID}

	result, err := u.userCollection.DeleteOne(u.ctx, filter)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	if result.DeletedCount == 0 {
		log.Println("User not found!")
		return mongo.ErrNoDocuments
	}

	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (models.UserLoginResponse, error) {
	var (
		user              domain.User
		userLoginResponse models.UserLoginResponse
	)
	err := u.userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userLoginResponse, err
		}
		log.Println("Error retrieving user: ", err)
		return userLoginResponse, err
	}

	userLoginResponse.ID = user.ID.Hex()
	userLoginResponse.Password = user.Password

	return userLoginResponse, nil
}
