package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pythagoreSynchroniser/models"
)

// SendInvoiceToFNE envoie une facture à l'API FNE et renvoie les métadonnées obtenues.
func SendInvoiceToFNE(invoice models.FNEInvoiceRequest, token string) (string, string, error) {
	url := os.Getenv("FNE_API_URL")
	if url == "" {
		return "", "", fmt.Errorf("FNE_API_URL n'est pas défini")
	}

	if token == "" {
		token = os.Getenv("FNE_API_TOKEN")
		if token == "" {
			return "", "", fmt.Errorf("FNE_API_TOKEN manquant")
		}
	}

	body, err := json.Marshal(invoice)
	if err != nil {
		return "", "", fmt.Errorf("JSON marshal failed: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", "", fmt.Errorf("API error: %s", string(respBody))
	}

	var response struct {
		Reference string `json:"reference"`
		Token     string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", "", fmt.Errorf("failed to parse response: %w", err)
	}

	fmt.Printf("Facture certifiée : ref=%s, QR token=%s\n", response.Reference, response.Token)
	return response.Reference, response.Token, nil
}
