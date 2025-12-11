package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type LivroSimplificado struct {
	Titulo    string   `json:"titulo"`
	Autores   []string `json:"autores"`
	Descricao string   `json:"descricao"`
	Imagem    string   `json:"imagem"` 
}

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Title       string   `json:"title"`
			Authors     []string `json:"authors"`
			Description string   `json:"description"`
			ImageLinks  struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pesquisando livros...")

	query := r.URL.Query().Get("q")
	if query == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"erro": "O parâmetro de busca 'q' é obrigatório"})
		return
	}

	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, googleURL, nil)

	resp, err := http.DefaultClient.Do(gReq)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"erro": err.Error()})
		return
	}
	defer resp.Body.Close()

	
	var googleResp GoogleBooksResponse
	if err := json.NewDecoder(resp.Body).Decode(&googleResp); err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"erro": "Erro ao processar resposta do Google"})
		return
	}

	var livros []LivroSimplificado

	for _, item := range googleResp.Items {
		livro := LivroSimplificado{
			Titulo:    item.VolumeInfo.Title,
			Autores:   item.VolumeInfo.Authors,
			Descricao: item.VolumeInfo.Description,
			Imagem:    item.VolumeInfo.ImageLinks.Thumbnail,
		}
		livros = append(livros, livro)
	}
	writeJson(w, http.StatusOK, livros)
}

func writeJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}