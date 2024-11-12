package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/99designs/keyring"
)
func RemoveFolders(rootFolder string, foldersToRemove []string) {
	for _, folder := range foldersToRemove {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}

func ResolveImportErr(dir string) error {
	currentDir, _ := os.Getwd()
	folder := filepath.Join(currentDir, dir)
	if err := os.Chdir(folder); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	commands := []struct {
		name string
		args []string
	}{
		{"go", []string{"mod", "download"}},
		{"go", []string{"mod", "tidy"}},
		{"go", []string{"install", "golang.org/x/tools/cmd/goimports@latest"}},
		{"goimports", []string{"-w", "."}},
		{"go", []string{"mod", "tidy"}},
	}

	for _, cmd := range commands {
		command := exec.Command(cmd.name, cmd.args...)
		
		command.Stdout = nil
		command.Stderr = nil

		if err := command.Run(); err != nil {
			return fmt.Errorf("failed to execute %s: %w", cmd.name, err)
		}
	}

	return nil
}

func StoreKey(key string, value string) error {
    kr, err := keyring.Open(keyring.Config{
        ServiceName: "go-sail",
    })
    if err != nil {
        return err
    }

    err = kr.Set(keyring.Item{
        Key:  key,
        Data: []byte(value),
    })
    if err != nil {
        return err
    }

    return nil
}

func GetKey(key string) (string, error) {
    kr, err := keyring.Open(keyring.Config{
        ServiceName: "go-sail",
    })
    if err != nil {
        return "", err
    }

    item, err := kr.Get(key)
    if err != nil {
        return "", err
    }

    return string(item.Data), nil
}

type Config struct {
	RepoURL string `json:"repo_url"`
}

type ConfigManager struct {
	configPath string
}

func NewConfigManager() (*ConfigManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %v", err)
	}

	configDir := filepath.Join(homeDir, ".myapp")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %v", err)
	}

	return &ConfigManager{
		configPath: filepath.Join(configDir, "config.json"),
	}, nil
}

func (cm *ConfigManager) LoadConfig() (*Config, error) {
	config := &Config{}

	data, err := os.ReadFile(cm.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

func (cm *ConfigManager) SaveConfig(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(cm.configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}