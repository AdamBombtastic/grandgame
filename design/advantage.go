package design

import (
	. "goa.design/goa/v3/dsl"
)

var Advantage = Type("Advantage", func() {
	Field(1, "Id", Int, "Unique advantage ID")
	Field(2, "Name", String, "Name of advantage")
	Field(3, "Description", String, "Description of advantage")
	Field(4, "Events", ArrayOf(String), "Events that allow this advantage")
	Field(5, "Tier", String, "Tier of advantage", func() {
		Enum("Slight", "Major", "Busted")
	})
	Required("Id", "Name", "Description", "Events", "Tier")
})
