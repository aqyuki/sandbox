package design

import (
	//lint:ignore ST1001 This is the goa design package
	. "goa.design/goa/v3/dsl"
)

var TodoCreatePayload = Type("TodoCreatePayload", func() {
	Attribute("title", String)
	Attribute("desc", String)
})

var _ = API("todo", func() {
	Title("ToDo Service")
	Description("Service for managing ToDo lists")
	Server("Develop", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
	Server("Production", func() {
		Host("example.com", func() {
			URI("http://example.com")
		})
	})
})

var _ = Service("todo", func() {
	Description("The todo service manages todo lists")

	Method("create", func() {
		Payload(TodoCreatePayload)
		Result(Int)
		HTTP(func() {
			POST("/task/create")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
