package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Integration test: register the route on a gin engine and exercise it via ServeHTTP
func TestCreateAuthor_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// create router and register the handler under the same path used by the app
	router := gin.New()
	api := router.Group("/api/v1")
	api.POST("/author/:name", CreateAuthor)

	// build request body
	requestBody := map[string]string{
		"content":     "integration content",
		"created_at":  "2025-11-13T00:00:00Z",
		"modified_by": "tester",
	}
	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/author/Alice", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, w.Code, w.Body.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if resp["name"] != "Alice" {
		t.Errorf("expected name 'Alice', got %v", resp["name"])
	}
	if resp["content"] != requestBody["content"] {
		t.Errorf("expected content %q, got %v", requestBody["content"], resp["content"])
	}
}
