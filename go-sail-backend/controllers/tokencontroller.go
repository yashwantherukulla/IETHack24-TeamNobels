package controllers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shreeyash-ugale/go-sail-server/database"
	"github.com/shreeyash-ugale/go-sail-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenAuth struct {
	Email string `json:"email"`
	Key   string `json:"key"`
}

func GenerateAPIKey(c *gin.Context) {
	//var request TokenRequest
	var user models.User
	var apik models.APIKey
	/*
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}*/
	user = c.MustGet("user").(models.User)
	// check if email exists and password is correct
	//filter := bson.M{"email": request.Email}
	/*
		err := database.UserCollection.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}
		credentialError := user.CheckPassword(request.Password)
		if credentialError != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			c.Abort()
			c.Abort()
			return
		}*/
	apiKey, err := genkey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"email": user.Email}
	user.APIKey = append(user.APIKey, apiKey)
	_, err = database.UserCollection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"api_key": user.APIKey}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	apik.ID = primitive.NewObjectID()
	apik.Key = apiKey
	apik.UserID = user.ID
	apik.PlanID = user.PlanID
	_, err = database.APIKeyCollection.InsertOne(ctx, apik)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"api_key": apiKey})
}

func genkey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func RevokeAPIKey(c *gin.Context) {
	var user models.User
	var request TokenAuth
	user = c.MustGet("user").(models.User)
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	/*

		// check if email exists
		filter := bson.M{"email": request.Email}


		err := database.UserCollection.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}*/

	// remove the API key from user's APIKey slice
	for i, key := range user.APIKey {
		if key == request.Key {
			user.APIKey = append(user.APIKey[:i], user.APIKey[i+1:]...)
			break
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"email": user.Email}
	_, err := database.UserCollection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"api_key": user.APIKey}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// delete the API key from APIKeyCollection
	_, err = database.APIKeyCollection.DeleteOne(ctx, bson.M{"key": request.Key})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "API key deleted successfully"})
}

func GetAPIKeys(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	c.JSON(http.StatusOK, gin.H{"api_keys": user.APIKey})
}
