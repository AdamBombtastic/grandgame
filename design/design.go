package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
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
	cors.Origin("*")

	Method("list_participants", func() {
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

	Method("list_advantages", func() {
		Description("Returns all the advantages and their information")
		Result(ArrayOf(Advantage))
		Error("not_found")
		Error("internal_error")
		HTTP(func() {
			GET("/advantages")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
		})
	})

	Method("list_competition_event_kinds", func() {
		Description("Returns all the competition event kinds")
		Result(ArrayOf(CompetitionEventDescription))
		Error("not_found")
		Error("internal_error")
		HTTP(func() {
			GET("/competition_events/kinds")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
		})
	})
})
