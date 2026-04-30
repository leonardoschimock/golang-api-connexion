package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bateu aqui")

	query := r.URL.Query().Get("q")

	if query == "" {
		writeJSON(w, http.StatusBadRequest, "Query param 'q' é obrigatório")
		return
	}

	if query == "test" {
		mock := `{
		"items": [
			{
				"volumeInfo": {
					"title": "Titulo Teste",
					"authors": ["Autor Teste"]
				}
			}
		]
	}`

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mock))
		return
	}

	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, googleURL, nil)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, "Erro ao criar request")
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, "Erro ao chamar API externa")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, "Erro ao ler resposta")
		return
	}

	if resp.StatusCode != http.StatusOK {
		writeJSON(w, resp.StatusCode, string(body))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
