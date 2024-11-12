package models

type Config struct {
	Repositories  map[string]string                       `yaml:"repositories"`
	Databases     map[string]DatabaseConfig               `yaml:"databases"`
	ORMs          map[string]ORMConfig                    `yaml:"orms"`
	Combinations  map[string]map[string]CombinationConfig `yaml:"combinations"`
	MigrationCode map[string]string                       `yaml:"migrationCode"`
}

type DatabaseConfig struct {
	Name      string `yaml:"name"`
	DriverPkg string `yaml:"driverPkg"`
}

type ORMConfig struct {
	Name       string `yaml:"name"`
	ImportPath string `yaml:"importPath"`
}

type CombinationConfig struct {
	DSNTemplate       string   `yaml:"dsnTemplate"`
	InitFunc          string   `yaml:"initFunc"`
	AdditionalImports []string `yaml:"additionalImports"`
}
