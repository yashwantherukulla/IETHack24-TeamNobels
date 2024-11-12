package prompts

import (
    "github.com/AlecAivazis/survey/v2"
    "context"
    "errors"
)

var frameworks = []string{"fiber", "gin", "echo",}

var databases = []string{"postgres", "mysql", "None"}

var orms = []string{"gorm", "sqlx", "None"}

func SelectFramework(ctx context.Context) (string, error) {
    var framework string
    prompt := &survey.Select{
        Message: "🚀 Choose a Go framework:",
        Options: frameworks,
        Default: "fiber",
        Help:    "Select the framework you want to use for your project",
    }
    errCh := make(chan error, 1)
	go func() {
		errCh <- survey.AskOne(prompt, &framework)
	}()

	select {
	case <-ctx.Done():
		return "", errors.New("interrupt")
	case err := <-errCh:
		if err != nil {
			return "", err
		}
		return framework, nil
	}
}

func SelectDatabase(ctx context.Context) (string, error) {
    var database string
    prompt := &survey.Select{
        Message: "💾 Choose a database (or None):",
        Options: databases,
        Default: "None",
        Help:    "Select the database you want to use, or 'None' if you don't need one",
    }
    errCh := make(chan error, 1)
	go func() {
		errCh <- survey.AskOne(prompt, &database)
	}()

	select {
	case <-ctx.Done():
		return "", errors.New("interrupt")
	case err := <-errCh:
		if err != nil {
			return "", err
		}
		if database == "None" {
			return "", nil
		}
		return database, nil
	}
}

func SelectORM(ctx context.Context) (string, error) {
    var orm string
    prompt := &survey.Select{
        Message: "🔗 Choose an ORM (or None):",
        Options: orms,
        Default: "None",
        Help:    "Select an ORM for database interactions, or 'None' if you don't need one",
    }
    errCh := make(chan error, 1)
	go func() {
		errCh <- survey.AskOne(prompt, &orm)
	}()

	select {
	case <-ctx.Done():
		return "", errors.New("interrupt")
	case err := <-errCh:
		if err != nil {
			return "", err
		}
		if orm == "None" {
			return "", nil
		}
		return orm, nil
	}
}