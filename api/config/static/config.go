package static

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"restaurant-api/api/config/static/model"
	"runtime"
)

var (
	config   *model.Config
	basePath = "./static"
)

func GetConfig() *model.Config {
	if config == nil {
		config = getPropertiesConfig()
	}

	return config
}

func getPropertiesConfig() *model.Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	if os.Getenv("scope") == "dev" {
		_, currentDirectory, _, ok := runtime.Caller(0)
		if ok {
			currentPath := filepath.Dir(currentDirectory)
			basePath = filepath.Join(currentPath, filepath.FromSlash("../../config/static"))
		}
	}
	filePath := "/dev/config.json"
	scope := os.Getenv("SCOPE")
	if scope != "" {
		filePath = fmt.Sprint("/scopes/", scope, ".json")
	}
	unifiedConfig := getUnifiedConfig(filePath)
	return &unifiedConfig
}

func getUnifiedConfig(pathConfig string) model.Config {
	config := &model.Config{}
	listConfig := []string{"/config.json", pathConfig}
	for _, jsonConfig := range listConfig {
		if jsonConfig != "" {
			getConfig(jsonConfig, config)
		}
	}
	return *config
}

func getConfig(pathFileConfig string, config *model.Config) {
	bytes, err := os.ReadFile(fmt.Sprint(basePath, pathFileConfig))
	if err != nil {
		print("Config static not found", pathFileConfig)
		return
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		panic(fmt.Errorf("error while loading json scope config - %v", err))
	}
}
