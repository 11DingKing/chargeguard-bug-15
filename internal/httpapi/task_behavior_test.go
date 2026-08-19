package httpapi

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	rr := httptest.NewRecorder()
	TaskHTTPHandler(rr, httptest.NewRequest("GET", "/task?tenant=runan&status=open", nil))
	body := rr.Body.String()
	if strings.Contains(body, "neighbor") || !strings.Contains(body, "\"total\":1") {
		t.Fatalf("body=%s", body)
	}
}
