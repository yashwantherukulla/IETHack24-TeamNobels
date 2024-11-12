package prompts

import (
	"context"

	"github.com/TejasGhatte/go-sail/internal/helpers"
)

func PromptUserSignupDetails(ctx context.Context) (string, string, string, error) {
	username := helpers.PromptUserInput(ctx, "Enter username: ")
	if ctx.Err() != nil {
		return "", "", "", ctx.Err()
	}

	email := helpers.PromptUserInput(ctx, "Enter email: ")
	if ctx.Err() != nil {
		return "", "", "", ctx.Err()
	}

	password := helpers.PromptUserPassword(ctx, "Enter password: ")
	if ctx.Err() != nil {
		return "", "", "", ctx.Err()
	}

	return username, email, password, nil
}