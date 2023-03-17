package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	tests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{
			"home", "/", http.StatusOK,
		},
		{
			"404", "/fish", http.StatusNotFound,
		},
	}

	app := application{}
	routes := app.routes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	path2Templates = "./../../templates/"

	for _, tt := range tests {
		resp, err := ts.Client().Get(ts.URL + tt.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != tt.expectedStatusCode {
			t.Errorf("expected %d, but get %d\n", tt.expectedStatusCode, resp.StatusCode)
		}
	}
}
