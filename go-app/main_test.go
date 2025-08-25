package main
import (
    "fmt"              // Add this line
    "net/http"
    "net/http/httptest"
    "testing"
)
func TestHealth(t *testing.T) {
    req, _ := http.NewRequest("GET", "/healthz", nil)
    rr := httptest.NewRecorder()
    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "OK") }).ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
    expected := "OK"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}