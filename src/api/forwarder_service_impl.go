package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	f "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/gen/forwarder_service"
	c "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/config"
	"net/http"
)

// ForwarderService implements the forwarder service
type ForwarderService struct {
	client *http.Client
	url    string
	chatID string
}

// NewForwarderService creates a new forwarder service
func NewForwarderService(id string) *ForwarderService {
	return &ForwarderService{
		client: &http.Client{},
		url:    "https://api.telegram.org/bot" + c.TELEGRAM_BOT_TOKEN,
		chatID: id,
	}
}

// Forward forwards a notification if it is a warning, otherwise it ignores it
func (s *ForwarderService) Forward(ctx context.Context, m *f.Notification) (err error) {
	// switch case for future extensibility
	switch m.Type {
	case "Warning":
		if err := s.forward(ctx, m); err != nil {
			return err
		}
	default:
		break
	}
	return nil
}

// forwardRequestBody defines the request body for the forward method of the telegram bot api
type forwardRequestBody struct {
	ChatID string `json:"chat_id" required:"true"`
	Text   string `json:"text" required:"true"`
}

// forward forwards a warning to the appropriate channel
func (s *ForwarderService) forward(ctx context.Context, m *f.Notification) (err error) {
	body, err := json.Marshal(
		forwardRequestBody{
			ChatID: s.chatID,
			Text:   "Warning\n" + m.Description,
		},
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		s.url+"/sendMessage",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
