package api

import (
	"context"
	"fmt"
	f "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/gen/forwarder_service"
)

// ForwarderService implements the forwarder service
type ForwarderService struct{}

// NewForwarderService creates a new forwarder service
func NewForwarderService() *ForwarderService {
	return &ForwarderService{}
}

// Forward forwards a notification if it is a warning, otherwise it ignores it
func (s *ForwarderService) Forward(ctx context.Context, m *f.Notification) (err error) {
	// switch case for future extensibility
	switch m.Type {
	case "Warning":
		// TODO: Forward to the appropriate channel
		fmt.Println("sending to channel")
	default:
		//TODO:  ignore
		fmt.Println("ignoring")
	}
	return nil
}

// forward forwards a warning to the appropriate channel
func (s *ForwarderService) forward(ctx context.Context, m *f.Notification) (err error) {
	//TODO: implement
	return nil
}
