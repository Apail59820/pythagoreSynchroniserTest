package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pythagoreSynchroniser/models"
	"testing"
)

func TestSendInvoiceToFNE_NoURL(t *testing.T) {
	t.Setenv("FNE_API_URL", "")
	_, _, err := SendInvoiceToFNE(models.FNEInvoiceRequest{}, "")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSendInvoiceToFNE_NoToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer srv.Close()
	t.Setenv("FNE_API_URL", srv.URL)
	t.Setenv("FNE_API_TOKEN", "")
	_, _, err := SendInvoiceToFNE(models.FNEInvoiceRequest{}, "")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSendInvoiceToFNE_Non200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}))
	defer srv.Close()
	t.Setenv("FNE_API_URL", srv.URL)
	t.Setenv("FNE_API_TOKEN", "tok")
	_, _, err := SendInvoiceToFNE(models.FNEInvoiceRequest{}, "")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSendInvoiceToFNE_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req models.FNEInvoiceRequest
		json.NewDecoder(r.Body).Decode(&req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"reference": "ref", "token": "tok"})
	}))
	defer srv.Close()
	t.Setenv("FNE_API_URL", srv.URL)
	t.Setenv("FNE_API_TOKEN", "tok")
	ref, tok, err := SendInvoiceToFNE(models.FNEInvoiceRequest{InvoiceType: "A"}, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ref != "ref" || tok != "tok" {
		t.Fatalf("unexpected response %s %s", ref, tok)
	}
}
