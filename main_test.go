package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adminevolsystemcl/hola-jungla-de-bits/internal/hostinfo"
)

func TestHostInfoCollectsRuntimeData(t *testing.T) {
	info, err := hostinfo.Collect()
	if err != nil {
		t.Fatalf("Collect() error = %v", err)
	}

	if info.Hostname == "" {
		t.Fatal("expected hostname to be set")
	}

	if info.OS == "" || info.Arch == "" || info.GoVersion == "" {
		t.Fatal("expected runtime metadata to be set")
	}

	if info.CPUs < 1 {
		t.Fatalf("expected CPUs > 0, got %d", info.CPUs)
	}

	if info.CollectedAt == "" {
		t.Fatal("expected collected timestamp to be set")
	}
}

func TestHostInfoHandlerReturnsJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/host-info", nil)
	rec := httptest.NewRecorder()

	hostInfoHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if got := rec.Header().Get("Content-Type"); !strings.Contains(got, "application/json") {
		t.Fatalf("expected json content type, got %q", got)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "\"hostname\"") {
		t.Fatalf("expected hostname in response body, got %q", body)
	}
}

func TestHostInfoHandlerReturnsTextForCLIConsumers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/host-info?format=text", nil)
	rec := httptest.NewRecorder()

	hostInfoHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if got := rec.Header().Get("Content-Type"); !strings.Contains(got, "text/plain") {
		t.Fatalf("expected text content type, got %q", got)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "hostname:") {
		t.Fatalf("expected text body, got %q", body)
	}
}

func TestHostInfoHandlerRejectsUnsupportedMethods(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/host-info", nil)
	rec := httptest.NewRecorder()

	hostInfoHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rec.Code)
	}
}
