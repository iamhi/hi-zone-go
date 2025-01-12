package ollamaclient_requests

type EmbeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}
