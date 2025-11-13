package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateAuthor_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	requestBody := map[string]string{
		"content":     "hello world",
		"created_at":  "2025-01-01T00:00:00Z",
		"modified_by": "bob",
	}
	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/author", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	ctx.Request = req

	// set path param "name" used by the handler
	ctx.Params = gin.Params{{Key: "name", Value: "Alice"}}

	CreateAuthor(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, w.Code, w.Body.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	// validate fields
	if resp["name"] != "Alice" {
		t.Errorf("expected name 'Alice', got %v", resp["name"])
	}
	if resp["content"] != requestBody["content"] {
		t.Errorf("expected content %q, got %v", requestBody["content"], resp["content"])
	}
	if resp["created_at"] != requestBody["created_at"] {
		t.Errorf("expected created_at %q, got %v", requestBody["created_at"], resp["created_at"])
	}
	if resp["modified_by"] != requestBody["modified_by"] {
		t.Errorf("expected modified_by %q, got %v", requestBody["modified_by"], resp["modified_by"])
	}
}

func TestCreateAuthor_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	invalidJSON := []byte(`{invalid-json`)

	req, _ := http.NewRequest(http.MethodPost, "/opening/Bob", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	ctx.Request = req

	ctx.Params = gin.Params{{Key: "name", Value: "Bob"}}

	CreateAuthor(ctx)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if _, ok := resp["error"]; !ok {
		t.Errorf("expected response to contain 'error' field, got: %v", resp)
	}
}
