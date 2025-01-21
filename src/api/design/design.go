package design

import (
	dsl "goa.design/goa/v3/dsl"
)

// ForwarderAPI defines the endpoint & associated server
var _ = dsl.API("ForwarderAPI", func() {
	dsl.Title("Forwarder API")
	dsl.Description("The forwarder api forwards warnings to the appropriate channel")
	dsl.Server("ForwarderServer", func() {
		dsl.Description("The forwarder server")
		dsl.Services("ForwarderService")

		dsl.Host("development", func() {
			dsl.Description("Development hosts.")
			dsl.URI("http://localhost:80/forward")
			dsl.URI("grpc://localhost:8080")
		})
	})
})

// ForwarderService defines the service & associated method(s)
var _ = dsl.Service("ForwarderService", func() {
	dsl.Description("The forwarder service forwards warnings to the appropriate channel")
	dsl.Method("forward", func() {
		dsl.Description("Forwards a warning to the appropriate channel")
		dsl.Payload(Message)

		dsl.HTTP(func() {
			dsl.POST("/forward")
			dsl.Response(dsl.StatusCreated)
		})

		dsl.GRPC(func() {
			dsl.Response(dsl.CodeOK)
		})
	})
})

// Message defines the payload for the ForwarderService's forward method.
// It contains the type, name & description of the event.
var Message = dsl.Type("Message", func() {
	dsl.Field(1, "type", dsl.String, func() {
		dsl.Description("The type of message")
	})
	dsl.Field(2, "name", dsl.String, func() {
		dsl.Description("The name of the event")
	})
	dsl.Field(3, "description", dsl.String, func() {
		dsl.Description("The description of the event")
	})
	// TODO: Add more fields as needed

	dsl.Required("type", "name", "description")
})
