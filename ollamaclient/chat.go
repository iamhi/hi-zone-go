package ollamaclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	ollamaclient_dto "github.com/iamhi/hi-zone-go/ollamaclient/dto"
	ollamaclient_errors "github.com/iamhi/hi-zone-go/ollamaclient/errors"
	ollamaclient_requests "github.com/iamhi/hi-zone-go/ollamaclient/requests"
	ollamaclient_response "github.com/iamhi/hi-zone-go/ollamaclient/responses"
)

const application_json_media_type = "application/json"

const model = "qwen3.5-coder:1.5b"

func SendChatMessage(message string, history []ollamaclient_dto.ChatMessageDto) (ollamaclient_dto.ChatMessageDto, ollamaclient_errors.OllamaHandlerError) {
	request := ollamaclient_requests.ChatRequest{
		Model: model,
		Messages: append(history, ollamaclient_dto.ChatMessageDto{
			Role:    "user",
			Content: message,
		}),
		Stream: false,
	}

	request_body, err := json.Marshal(request)

	if err != nil {
		fmt.Printf("Unable to marshal request: %s", err)

		return ollamaclient_dto.ChatMessageDto{}, ollamaclient_errors.OllamaGenericError{}
	}

	response, err := http.Post(
		"http://wildberry.local:11434/api/chat",
		application_json_media_type,
		bytes.NewBuffer(request_body))

	if err != nil {
		fmt.Printf("There was an error getting a response %s", err)

		return ollamaclient_dto.ChatMessageDto{}, ollamaclient_errors.OllamaGenericError{}
	}

	defer response.Body.Close()

	var response_body ollamaclient_response.ChatResponse

	err = json.NewDecoder(response.Body).Decode(&response_body)

	if err != nil {
		fmt.Printf("There was an error decoding the response %s", err)

		return ollamaclient_dto.ChatMessageDto{}, ollamaclient_errors.OllamaGenericError{}
	}

	return response_body.Message, nil
}
