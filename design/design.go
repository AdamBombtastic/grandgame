package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("game", func() {
	Title("The grand-game API")
	Description("This API serves the game")
	Server("game", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("game", func() {
	Description("The game service")
	Method("participants", func() {
		Description("Returns all the participants and their information")
		Result(ArrayOf(Participant))
		Error("not_found")
		Error("internal_error")
		HTTP(func() {
			GET("/participants")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
		})
	})
})
