package scripts

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/TejasGhatte/go-sail/internal/helpers"
)

func GitClone(projectName, templateType, templateURL string) error {

	if templateType == "" || templateURL == "" {
		return fmt.Errorf("project template not found")
	}

	currentDir, _ := os.Getwd()

	folder := filepath.Join(currentDir, projectName)

	_, errPlainClone := git.PlainClone(
		folder,
		false,
		&git.CloneOptions{
			URL: getAbsoluteURL(templateURL),
		},
	)
	if errPlainClone != nil {
		return fmt.Errorf("repository `%v` was not cloned", templateURL)
	}

	go helpers.RemoveFolders(folder, []string{".git", ".github"})

	return nil
}

func getAbsoluteURL(templateURL string) string {
	templateURL = strings.TrimSpace(templateURL)
	u, _ := url.Parse(templateURL)

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	return u.String()
}