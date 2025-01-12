package ollamaclient_requests

import ollamaclient_dto "github.com/iamhi/hi-zone-go/ollamaclient/dto"

type ChatRequest struct {
	Model    string                            `json:"model"`
	Messages []ollamaclient_dto.ChatMessageDto `json:"messages"`
	Stream   bool                              `json:"stream"`
}
