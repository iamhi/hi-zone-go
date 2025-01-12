package ollamaclient_response

import ollamaclient_dto "github.com/iamhi/hi-zone-go/ollamaclient/dto"

type ChatResponse struct {
	Model      string                          `json:"model"`
	CreatedAt  string                          `json:"created_at"`
	Message    ollamaclient_dto.ChatMessageDto `json:"message"`
	DoneReason string                          `json:"done_reason"`
	Done       bool                            `json:"done"`
}
