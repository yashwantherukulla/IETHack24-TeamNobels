package scripts

import (
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/helpers"
)

func RunSetRepo(repoURL string) error {
	cm, err := helpers.NewConfigManager()
	if err != nil {
		return fmt.Errorf("failed to initialize config manager: %v", err)
	}

	config, err := cm.LoadConfig()
	if err != nil {
		return err
	}

	// Update repository URL
	config.RepoURL = repoURL

	if err := cm.SaveConfig(config); err != nil {
		return err
	}

	fmt.Printf("Successfully set repository URL to: %s\n", repoURL)
	return nil
}