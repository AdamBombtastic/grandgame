package design

import (
	. "goa.design/goa/v3/dsl"
)

var Participant = Type("Participant", func() {
	Field(1, "Id", Int, "Unique participant ID")
	Field(2, "Name", String, "Name of participant")
	Field(3, "Email", String, "Email of participant")
	Field(4, "Phone", String, "Phone number of participant")
	Field(5, "Kind", String, "Kind of participant", func() {
		Enum("player", "admin", "superadmin")
	})
	Field(6, "Role", String, "Role of participant", func() {
		Enum("King", "King's Hand", "Archvizier", "Noble", "Mercenary")
	})
	Field(7, "Backstory", String, "Backstory of participant")
	Field(8, "Gold", Int, "Gold of participant")
	Field(9, "Favor", Int, "Influence of participant")

	Required("Id", "Name", "Email", "Phone", "Kind", "Role", "Backstory", "Gold", "Favor")
})
