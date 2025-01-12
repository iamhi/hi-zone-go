package ollamaclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	ollamaclient_requests "github.com/iamhi/hi-zone-go/ollamaclient/requests"
	ollamaclient_response "github.com/iamhi/hi-zone-go/ollamaclient/responses"
)

const default_embedding_model = "nomic-embed-text"
const default_ollama_url = "http://wildberry.local:11434/api/embed"

func GetEmbedding(text string) ([]float32, error) {
	url := default_ollama_url // TODO: get this from config file
	request_body, _ := json.Marshal(ollamaclient_requests.EmbeddingRequest{
		Model: default_embedding_model,
		Input: []string{text}})

	response, err := http.Post(url, "appllication/json", bytes.NewBuffer(request_body))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("failed to generate embedding") // TODO: make this a custom error
	}

	var embedding ollamaclient_response.EmbeddingResponse

	if err := json.NewDecoder(response.Body).Decode(&embedding); err != nil {
		return nil, err
	}

	return embedding.Embeddings[0], nil
}
