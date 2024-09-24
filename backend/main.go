package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// TenantConfig represents tenant-specific configuration
type TenantConfig struct {
	TenantID    string `json:"tenant_id"`
	Theme       string `json:"theme"`
	Template    string `json:"template"`
	Description string `json:"description"`
}

// TenantTestData represents test data for tenants
type TenantTestData struct {
	TenantID string `json:"tenant_id"`
	Data     string `json:"data"`
}

// Sample tenant-specific configurations
var tenantConfigs = map[string]TenantConfig{
	"tenant1": {
		TenantID:    "tenant1",
		Theme:       "dark",
		Template:    "template_1",
		Description: "Configuration for Tenant 1",
	},
	"tenant2": {
		TenantID:    "tenant2",
		Theme:       "dark",
		Template:    "template_2",
		Description: "Configuration for Tenant 2",
	},
}

// Sample tenant-specific test data
var tenantTestData = map[string]TenantTestData{
	"tenant1": {
		TenantID: "tenant1",
		Data:     "This is test data for Tenant 1",
	},
	"tenant2": {
		TenantID: "tenant2",
		Data:     "This is test data for Tenant 2",
	},
}

// GetTenantConfig handles the /api/config endpoint
func GetTenantConfig(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Header.Get("tenant_id")

	if tenantID == "" {
		http.Error(w, "Tenant ID is required", http.StatusBadRequest)
		return
	}

	config, exists := tenantConfigs[tenantID]
	if !exists {
		http.Error(w, "Config not found for tenant", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

// GetTenantTestData handles the /api/test endpoint
func GetTenantTestData(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Header.Get("tenant_id")

	if tenantID == "" {
		http.Error(w, "Tenant ID is required", http.StatusBadRequest)
		return
	}

	testData, exists := tenantTestData[tenantID]
	if !exists {
		http.Error(w, "Test data not found for tenant", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testData)
}

// InitRouter initializes the HTTP routes
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Define the routes
	router.HandleFunc("/api/config", GetTenantConfig).Methods("GET")
	router.HandleFunc("/api/test", GetTenantTestData).Methods("GET")

	return router
}

func main() {
	router := InitRouter()

	// Enable CORS with default options (allow all origins, methods, headers)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // You can specify the frontend domain here
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
