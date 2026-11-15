package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreetRoute(t *testing.T) {
	_, server := newApp(ServerConfig{})
	cases := []struct {
		path       string
		wantStatus int
		wantKey    string
		wantValue  string
	}{
		{path: "/greet/Ada", wantStatus: http.StatusOK, wantKey: "message", wantValue: "Hello, Ada!"},
		{path: "/greet/%20%20", wantStatus: http.StatusBadRequest, wantKey: "error", wantValue: "greeting: name is empty"},
		{path: "/greet/" + strings.Repeat("a", 65), wantStatus: http.StatusBadRequest, wantKey: "error", wantValue: "greeting: name is too long"},
	}
	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			rr := httptest.NewRecorder()
			server.Handler().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, tc.path, nil))
			if rr.Code != tc.wantStatus {
				t.Fatalf("GET %s status = %d, want %d, body %s", tc.path, rr.Code, tc.wantStatus, rr.Body)
			}
			var body map[string]string
			if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
				t.Fatalf("decode body: %v", err)
			}
			if body[tc.wantKey] != tc.wantValue {
				t.Fatalf("GET %s %s = %q, want %q", tc.path, tc.wantKey, body[tc.wantKey], tc.wantValue)
			}
		})
	}
}
