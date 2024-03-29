package database

import (
	"context"
	"medods/pkg/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user model.User) string {
	result, err := User.InsertOne(context.TODO(), bson.M{"_id": primitive.NewObjectID(), "refresh_token": user.RefreshToken})
	if err != nil {
		logrus.Errorf("user cannot be created: %s", err.Error())
		return ""
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedID
}

func GetUser(user model.User) (string, error) {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		logrus.Errorf("invalid id: %s", err.Error())
		return user.ID, err
	}

	result := User.FindOne(context.TODO(), bson.M{"_id": id})
	err = result.Decode(&user)
	if err != nil {
		logrus.Errorf("unable to find user: %s", err.Error())
		return user.ID, err
	}

	return user.ID, nil
}

func UpdateUser(user model.User) (string, error) {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		logrus.Errorf("invalid id: %s", err.Error())
		return user.ID, err
	}

	result := User.FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, bson.D{{Key: "$set", Value: bson.D{{Key: "refresh_token", Value: user.RefreshToken}}}})
	err = result.Decode(&user)
	if err != nil {
		logrus.Errorf("unable to find user: %s", err.Error())
		return user.ID, err
	}

	return user.ID, nil
}

func CheckRefreshValid(user model.User) string {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		logrus.Errorf("invalid id: %s", err.Error())
		return user.ID
	}

	refreshToken := user.RefreshToken

	result := User.FindOne(context.TODO(), bson.M{"_id": id})
	err = result.Decode(&user)
	if err != nil {
		logrus.Errorf("user not found: %s", err.Error())
		return ""
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.RefreshToken), []byte(refreshToken)); err != nil {
		logrus.Errorf("invalid refresh token: %s", err.Error())
		return ""
	}

	return refreshToken
}
