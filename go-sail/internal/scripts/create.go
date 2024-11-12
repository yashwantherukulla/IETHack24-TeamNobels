package scripts

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/TejasGhatte/go-sail/internal/helpers"
	"github.com/TejasGhatte/go-sail/internal/initializers"
	"github.com/TejasGhatte/go-sail/internal/models"
	"github.com/TejasGhatte/go-sail/internal/prompts"
	"github.com/TejasGhatte/go-sail/internal/signals"
	"github.com/briandowns/spinner"
)

func CreateProject(name string) error {
	framework, err := prompts.SelectFramework(context.Background())
	if err != nil {
		return err
	}
	database, err := prompts.SelectDatabase(context.Background())
	if err != nil {
		return err
	}

	var orm string
	if database != "" {
		orm, err = prompts.SelectORM(context.Background())
		if err != nil {
			return err
		}
	}

	fmt.Println("Generating project with the following options:")
	fmt.Printf("Framework: %s, Database: %s, ORM: %s\n", framework, database, orm)
	
	options := &models.Options{
		ProjectName: name,
		Framework:   framework,
		Database:    database,
		ORM:         orm,
	}

	// Use the HandleCancellation function to set up cancellation
	ctx := signals.HandleCancellation(context.Background())

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	s.Suffix = fmt.Sprintf(" Creating project: %s", name)
	s.Color("blue")
	defer s.Stop()

	err = PopulateDirectory(ctx, options, s)
	if err != nil {
		if err == context.Canceled {
			s.FinalMSG = fmt.Sprintf("Project creation cancelled: %s\n", name)
		} else {
			s.FinalMSG = fmt.Sprintf("Failed to create project: %s\n", name)
		}
		return err
	}

	s.FinalMSG = fmt.Sprintf("Created project: %s\n", name)
	return nil
}

func PopulateDirectory(ctx context.Context, options *models.Options, s *spinner.Spinner) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.Suffix = " Cloning template"
	if err := GitClone(options.ProjectName, options.Framework, initializers.Config.Repositories[options.Framework]); err != nil {
		return fmt.Errorf("error cloning repository: %v", err)
	}

	currentDir, _ := os.Getwd()
	folder := filepath.Join(currentDir, options.ProjectName, "initializers")

	if options.Database != "" && options.ORM != "" {
		provider, err := helpers.ProviderFactory(options.Database, options.ORM)
		if err != nil {
			return fmt.Errorf("error creating database provider: %v", err)
		}

		s.Suffix = " Generating database file"
		if err := ctx.Err(); err != nil {
			return err
		}
		err = helpers.GenerateDatabaseFile(folder, provider)
		if err != nil {
			return fmt.Errorf("error generating database file: %v", err)
		}

		s.Suffix = " Generating migration file"
		if err := ctx.Err(); err != nil {
			return err
		}
		err = helpers.GenerateMigrationFile(folder, provider)
		if err != nil {
			return fmt.Errorf("error generating migration file: %v", err)
		}

		s.Suffix = " Resolving imports"
		if err := ctx.Err(); err != nil {
			return err
		}
		err = helpers.ResolveImportErr(options.ProjectName)
		if err != nil {
			return fmt.Errorf("error resolving imports: %v", err)
		}
	}

	return nil
}