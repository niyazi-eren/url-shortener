package handler

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/blake2b"
	"io"
	"net/http"
	"net/url"
	"os"
	"url-shortener/model/entity"
	"url-shortener/model/request"
	"url-shortener/model/response"
	"url-shortener/repository"
)

var PUBLIC_DNS = "PUBLIC_DNS"

type AppHandler struct {
	Db *repository.AppRepository
}

func NewAppHandler() *AppHandler {
	return &AppHandler{Db: repository.NewAppRepository()}
}

// Post Request Handler - POST /app - shorten url
func (a *AppHandler) Post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading request body", http.StatusInternalServerError)
		return
	}

	data, err := parseJSON(body)
	if err != nil {
		http.Error(w, "error parsing JSON data", http.StatusBadRequest)
		return
	}

	if err := validateURL(data.Url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// store the url entity
	uri, err := a.generateAndSaveUrl(data.Url)
	if err != nil {
		http.Error(w, "error storing data", http.StatusInternalServerError)
		return
	}

	res := generateResponse(uri)
	writeJSONResponse(w, http.StatusOK, res)
}

// Get Request Handler - GET /{shortUrl} - redirect to url if found
func (a *AppHandler) Get(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "shortUrl")
	uri, err := a.Db.Find(key)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, uri.LongUrl, http.StatusFound)
}

// Delete Request Handler - Delete /{shortUrl} - delete url if found
func (a *AppHandler) Delete(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "shortUrl")
	count, err := a.Db.Delete(key)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if count > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		http.NotFound(w, r)
	}
}

// Get /health request handler - returns http.StatusOK
func (a *AppHandler) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *AppHandler) generateAndSaveUrl(url string) (entity.Url, error) {
	hash, _ := hashWithSize([]byte(url), 4)
	hashString := hex.EncodeToString(hash)
	uri := entity.Url{LongUrl: url, Hash: hashString}
	u, err := a.Db.Save(&uri)
	return *u, err
}

// hash the url with a specified size to keep the url short
func hashWithSize(data []byte, size int) ([]byte, error) {
	hash, err := blake2b.New(size, nil)
	if err != nil {
		return nil, err
	}
	_, err = hash.Write(data)
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func validateURL(uri string) error {
	if uri == "" {
		return errors.New("error: the URL is empty")
	}

	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return errors.New("error: invalid URL")
	}
	return nil
}

func generateResponse(url entity.Url) response.Response {
	return response.Response{
		Key:      url.Hash,
		LongUrl:  url.LongUrl,
		ShortUrl: "http://" + os.Getenv(PUBLIC_DNS) + ":8080/" + url.Hash,
	}
}

func parseJSON(body []byte) (request.Request, error) {
	var data request.Request
	err := json.Unmarshal(body, &data)
	return data, err
}

func writeJSONResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
