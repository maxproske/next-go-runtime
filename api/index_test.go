package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	Handler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want application/json", contentType)
	}

	expected := Response{Message: "Hello from Go!"}
	var got Response
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", got, expected)
	}
}
