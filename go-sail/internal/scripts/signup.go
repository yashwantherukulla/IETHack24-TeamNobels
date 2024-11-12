package scripts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/helpers"
	"github.com/TejasGhatte/go-sail/internal/prompts"
	"github.com/TejasGhatte/go-sail/internal/signals"
)

func Signup(ctx context.Context) error {
	ctx = signals.HandleCancellation(ctx)

	username, email, password, err := prompts.PromptUserSignupDetails(ctx)
	if err != nil {
		return err
	}

	payload := map[string]interface{}{
		"username": username,
		"email":    email,
		"password": password,
		"plan": "Premium",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	response, err := helpers.MakeReq("url", jsonData)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var responseBody struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		ApiKey  string `json:"api_key"`
		Username string `json:"username"`
	}

	if err := helpers.StoreKey(fmt.Sprintf("sail-api-key-%s", username), responseBody.ApiKey); err != nil {
		return err
	}

	if err := helpers.StoreKey("username", responseBody.Username); err != nil {
		return err
	}

	if response.StatusCode != 200 {
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&responseBody); err == nil {
			return err
		} else {
			fmt.Printf("Error decoding response: %v\n", err)
		}
		return fmt.Errorf("error calling server")
	}
	fmt.Println("Signup successful!")
	return nil
}