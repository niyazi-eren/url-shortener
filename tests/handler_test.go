package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/handler"
	"url-shortener/model/request"
	"url-shortener/model/response"
	"url-shortener/repository"
	"url-shortener/tests/mocks"
)

func TestUrlIsShortened(t *testing.T) {
	req := makePostRequest(t, "http://example.com")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	fakeKVStore := mocks.NewFakeKVStore()
	appHandler := handler.NewAppHandler(repository.NewAppRepository(fakeKVStore))

	// Call the Post handler
	appHandler.Post(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v", http.StatusOK, rr.Code)
	}

	var resp = &response.Response{}
	err := json.Unmarshal(rr.Body.Bytes(), resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Check url is shortened
	if resp.Key == "" {
		t.Errorf("Expected key")
	}
}

func TestBadUrl(t *testing.T) {
	req := makePostRequest(t, "htt")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create an instance of AppHandler
	fakeKVStore := mocks.NewFakeKVStore()
	appHandler := handler.NewAppHandler(repository.NewAppRepository(fakeKVStore))

	// Call the Post handler
	appHandler.Post(rr, req)

	// Check the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v; got %v", http.StatusBadRequest, rr.Code)
	}
}

func TestDeleteUrl(t *testing.T) {
	// Create an instance of AppHandler with a fake KV store
	fakeKVStore := mocks.NewFakeKVStore()
	appHandler := handler.NewAppHandler(repository.NewAppRepository(fakeKVStore))

	// setup
	existingKey := "abcd"
	_, _ = fakeKVStore.Set(existingKey, "lol")

	// Test case: Successfully delete an existing URL
	testDelete(t, appHandler, existingKey, http.StatusOK)

	// Test case: Delete a non-existing URL
	nonExistingKey := "notFound"
	testDelete(t, appHandler, nonExistingKey, http.StatusNotFound)
}

func testDelete(t *testing.T, appHandler *handler.AppHandler, key string, expectedStatusCode int) {
	// Create a DELETE request
	delReq := makeDelRequest(t, key)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the Delete handler
	appHandler.Delete(rr, delReq)

	// Check the response status code
	if rr.Code != expectedStatusCode {
		t.Errorf("Expected status %d; got %d", expectedStatusCode, rr.Code)
	}
}

func makePostRequest(t *testing.T, url string) *http.Request {
	// Create a sample URL request
	req := &request.Request{Url: url}

	// Encode the URL entity as JSON
	payload, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Create a request with the JSON payload
	resp, err := http.NewRequest("POST", "/app", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	return resp
}

func makeDelRequest(t *testing.T, url string) *http.Request {
	reqURL := "/" + url
	resp, err := http.NewRequest("DELETE", reqURL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	return resp
}
