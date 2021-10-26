package settings

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ExposeRoutes         bool
	EnableFirebaseConfig bool
	FirebaseConfigPath   bool
	EnableFirebaseAuth   bool   // TODO: implement this. If enabled, all routes are secure by default (must set secure = false to allow anonymous)
	CloudRunProjectTag   string // TODO: should be configurable per-route as well
	CloudFunctionBaseUrl string // TODO: should be configureable per-route as well
	Routes               map[string]Route
}

func GetConfig() Config {
	config := Config{}

	log.Printf("Loading environment variables")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	exposeRoutes, wasSet := os.LookupEnv("EXPOSE_ROUTES")
	if !wasSet {
		exposeRoutes = "false"
	}

	config.ExposeRoutes = strings.ToLower(exposeRoutes) == "true"

	enableFirebaseConfig, wasSet := os.LookupEnv("ENABLE_FIREBASE_CONFIG")
	if !wasSet {
		enableFirebaseConfig = "false"
	}
	config.EnableFirebaseConfig = strings.ToLower(enableFirebaseConfig) == "true"

	firebaseConfigPath := os.Getenv("FIREBASE_CONFIG_PATH")

	if config.EnableFirebaseConfig && firebaseConfigPath == "" {
		log.Fatal("FIREBASE_CONFIG_PATH is required if ENABLE_FIREBASE_CONFIG is true")
	}

	if config.EnableFirebaseConfig {
		log.Fatal("ENABLE_FIREBASE_CONFIG is not yet supported")
	} else {
		routeConfig, wasSet := os.LookupEnv("ROUTES_CONFIG")
		if !wasSet {
			log.Fatal("ROUTES_CONFIG is required if ENABLE_FIREBASE_CONFIG is false")
		}

		var routes []Route
		err = json.Unmarshal([]byte(routeConfig), &routes)
		if err != nil {
			log.Fatalf("Could not parse ROUTES_CONFIG: %s", err)
		}

		routeMap := make(map[string]Route)
		for _, route := range routes {
			valid := route.validateRoute()
			if valid {
				routeMap[route.Prefix] = route
			}
		}

		config.Routes = routeMap
	}

	return config
}
