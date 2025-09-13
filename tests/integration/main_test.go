package integrationtest

import (
	"os"
	"testing"

	infrastructure "lettersoup/src/infrastructure"
	database "lettersoup/src/infrastructure/database/lettersoup"
	"lettersoup/src/infrastructure/providers"
)

func TestMain(m *testing.M) {
	infrastructure.Initialize()
	providers.Logger.Info("Running tests setup...")
	providers.Logger.Info("Migrating DummyEntity for RepositoryBase tests...")
	database.DB.AutoMigrate(&DummyEntity{})
	providers.Logger.Info("Tests setup completed.")
	code := m.Run()
	os.Exit(code)
}
