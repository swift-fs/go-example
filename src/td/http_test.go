package td

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{
			name:   "test01",
			param:  "test",
			expect: `{"message":"Hello, World!","name":"test"}`,
		},
		{
			name:   "test02",
			param:  "",
			expect: `{"message":"Hello, World!"}`,
		},
	}

	r := SetupWebServer()
	for _, tt := range tests {
		req := httptest.NewRequest("GET", "/hello?name="+tt.param, nil)
		resp, _ := r.Test(req)
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expect %v, got %v", http.StatusOK, resp.StatusCode)
		}
		body, _ := io.ReadAll(resp.Body)
		if string(body) != tt.expect {
			t.Errorf("expect %v, got %v", tt.expect, string(body))
		}
	}
}
