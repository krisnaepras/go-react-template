package httpx

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeRejectsUnknownField(t *testing.T) {
	r := httptest.NewRequest("POST", "/", strings.NewReader(`{"title":"x","unknown":true}`))
	r.Header.Set("Content-Type", "application/json")
	var dst struct {
		Title string `json:"title"`
	}
	if err := DecodeJSON(httptest.NewRecorder(), r, &dst); err == nil {
		t.Fatal("expected unknown field rejection")
	}
}
