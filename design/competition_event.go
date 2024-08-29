package design

import (
	. "goa.design/goa/v3/dsl"
)

var CompetitionEventDescription = Type("CompetitionEventDescription", func() {
	Field(1, "Id", Int, "Unique competition event description ID")
	Field(2, "Name", String, "Name of competition event")
	Field(3, "Description", String, "Description of competition event")

	Required("Id", "Name", "Description")
})

var CompetitionEvent = Type("CompetitionEvent", func() {
	Field(1, "Id", Int, "Unique competition event ID")
	Field(2, "Name", String, "Name of competition event")
	Field(3, "Description", String, "Description of competition event")
	Field(4, "Participants", ArrayOf(Int), "Participants in the competition event")
	Field(5, "Status", String, "Status of the competition event", func() {
		Enum("Open", "Closed", "Running", "Finished")
	})
	Field(6, "Winner", Int, "Winner of the competition event")
	Field(7, "Reward", CompetitionReward, "Reward for the competition event")

	Required("Id", "Name", "Description", "Participants", "Status", "Reward")
})

var CompetitionReward = Type("CompetitionReward", func() {
	Field(1, "Favor", Int, "Favor reward for the competition event")
	Field(2, "Gold", Int, "Gold reward for the competition event")
	Required("Favor", "Gold")
})
