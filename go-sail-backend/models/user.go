package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	APIKey   []string           `bson:"api_key"`
	PlanID   primitive.ObjectID `bson:"plan_id"`
}

type APIKey struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Key    string             `bson:"key"`
	UserID primitive.ObjectID `bson:"user_id"`
	PlanID primitive.ObjectID `bson:"plan_id"`
}

type Plan struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Actions     []string             `bson:"actions"`
	Users       []primitive.ObjectID `bson:"users"`
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
