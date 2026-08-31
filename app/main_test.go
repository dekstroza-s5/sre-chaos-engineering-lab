package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	mode.Store("healthy")
	response := httptest.NewRecorder()
	health(response, httptest.NewRequest(http.MethodGet, "/health", nil))
	if response.Code != 200 { t.Fatalf("expected 200, got %d", response.Code) }
}

func TestInjectedError(t *testing.T) {
	mode.Store("error")
	response := httptest.NewRecorder()
	work(response, httptest.NewRequest(http.MethodGet, "/work", nil))
	if response.Code != 500 { t.Fatalf("expected 500, got %d", response.Code) }
}
