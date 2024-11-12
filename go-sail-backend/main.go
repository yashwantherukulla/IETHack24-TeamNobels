package main

import (
	"context"
	"log"

	"github.com/shreeyash-ugale/go-sail-server/controllers"
	"github.com/shreeyash-ugale/go-sail-server/database"
	"github.com/shreeyash-ugale/go-sail-server/middlewares"
	"github.com/shreeyash-ugale/go-sail-server/models"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("mongodb://localhost:27017")
	initPlansAndActions()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	apipass := router.Group("/user").Use(middlewares.UserAuth())
	{
		apipass.POST("/token", controllers.GenerateAPIKey)
		apipass.GET("/token", controllers.GetAPIKeys)
		apipass.DELETE("/token", controllers.RevokeAPIKey)
	}
	apiwkey := router.Group("/api").Use(middlewares.IsAPIOwner())
	{
		apiwkey.GET("/secured/eval", middlewares.IsAuthorized(), controllers.Ping)
		apiwkey.POST("/upgrade", controllers.UpgradePlan)
	}
	router.POST("/api/signup", controllers.Signup)
	return router
}

func initPlansAndActions() {
	// Define actions
	/*
		templateGenerate := models.Action{Name: "Template Generate", Description: "Generate templates"}
		dockerFileGenerate := models.Action{Name: "Docker File Generate", Description: "Generate Docker files"}
		codeEvaluation := models.Action{Name: "Code Evaluation", Description: "Evaluate code"}
		securityCheck := models.Action{Name: "Security Check", Description: "Perform security checks"}
	*/
	// Define plans
	// Define plans
	freePlan := models.Plan{
		Name:        "Free",
		Description: "Free plan with basic features",
		Actions:     []string{"templateGenerate", "dockerFileGenerate"},
	}
	premiumPlan := models.Plan{
		Name:        "Premium",
		Description: "Premium plan with additional features",
		Actions:     []string{"templateGenerate", "dockerFileGenerate", "codeEvaluation"},
	}
	executivePlan := models.Plan{
		Name:        "Executive",
		Description: "Executive plan with all features",
		Actions:     []string{"templateGenerate", "dockerFileGenerate", "codeEvaluation", "securityCheck"},
	}

	// Check if plans exist in the database, if not, add them
	var existingPlans []models.Plan
	cursor, err := database.PlanCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.Background(), &existingPlans); err != nil {
		log.Fatal(err)
	}

	plansMap := make(map[string]bool)
	for _, plan := range existingPlans {
		plansMap[plan.Name] = true
	}

	if !plansMap[freePlan.Name] {
		database.PlanCollection.InsertOne(context.Background(), freePlan)
	}
	if !plansMap[premiumPlan.Name] {
		database.PlanCollection.InsertOne(context.Background(), premiumPlan)
	}
	if !plansMap[executivePlan.Name] {
		database.PlanCollection.InsertOne(context.Background(), executivePlan)
	}

	/* Save actions and plans to the database
	database.ActionCollection.InsertOne(context.Background(), templateGenerate)
	database.ActionCollection.InsertOne(context.Background(), dockerFileGenerate)
	database.ActionCollection.InsertOne(context.Background(), codeEvaluation)
	database.ActionCollection.InsertOne(context.Background(), securityCheck)
	database.PlanCollection.InsertOne(context.Background(), freePlan)
	database.PlanCollection.InsertOne(context.Background(), premiumPlan)
	database.PlanCollection.InsertOne(context.Background(), executivePlan)
	*/
}
