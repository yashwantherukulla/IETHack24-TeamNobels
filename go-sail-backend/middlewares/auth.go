package middlewares

import (
	"github.com/shreeyash-ugale/go-sail-server/database"
	"github.com/shreeyash-ugale/go-sail-server/models"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

type reqBody1 struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type reqBody2 struct {
	Email string `json:"email"`
	Key   string `json:"key"`
}

func UserAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		var requestBody reqBody1
		if err := context.ShouldBindJSON(&requestBody); err != nil {
			context.JSON(400, gin.H{"error": "invalid request body"})
			context.Abort()
			return
		}
		var result models.User
		err := database.UserCollection.FindOne(context, bson.M{"email": requestBody.Email}).Decode(&result)
		if err != nil {
			context.JSON(401, gin.H{"error": "unauthorized"})
			context.Abort()
			return
		}
		err = result.CheckPassword(requestBody.Password)
		if err != nil {
			context.JSON(401, gin.H{"error": "unauthorized"})
			context.Abort()
			return
		}
		context.Set("user", result)
		context.Next()
	}
}

func IsAPIOwner() gin.HandlerFunc {
	return func(context *gin.Context) {
		var requestBody reqBody2
		if err := context.ShouldBindJSON(&requestBody); err != nil {
			context.JSON(400, gin.H{"error": "invalid request body"})
			context.Abort()
			return
		}

		var result models.User
		err := database.UserCollection.FindOne(context, bson.M{"email": requestBody.Email}).Decode(&result)
		if err != nil {
			context.JSON(401, gin.H{"error": "unauthorized"})
			context.Abort()
			return
		}
		var apikey models.APIKey
		err = database.APIKeyCollection.FindOne(context, bson.M{"key": requestBody.Key}).Decode(&apikey)
		if err != nil {
			context.JSON(401, gin.H{"error": "unauthorized"})
			context.Abort()
			return
		}
		if apikey.UserID != result.ID {
			context.JSON(401, gin.H{"error": "unauthorized"})
			context.Abort()
			return
		}
		context.Set("user", result)
		context.Next()
	}
}

func IsAuthorized() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, _ := context.Get("user")
		userData := user.(models.User)
		var plantype models.Plan
		err := database.PlanCollection.FindOne(context, bson.M{"_id": userData.PlanID}).Decode(&plantype)
		if err != nil {
			context.JSON(401, gin.H{"error": "unauthorized"})
			context.Abort()
			return
		}

		actionPresent := false
		for _, action := range plantype.Actions {
			if action == "codeEvaluation" {
				actionPresent = true
				break
			}
		}
		if !actionPresent {
			context.JSON(403, gin.H{"error": "forbidden, not in plan"})
			context.Abort()
			return
		}
		context.Next()
	}
}
