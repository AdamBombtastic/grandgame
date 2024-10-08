// Code generated by goa v3.18.2, DO NOT EDIT.
//
// game HTTP server types
//
// Command:
// $ goa gen github.com/Adambombtastic/grandgame/design

package server

import (
	game "github.com/Adambombtastic/grandgame/gen/game"
	goa "goa.design/goa/v3/pkg"
)

// ListParticipantsResponseBody is the type of the "game" service
// "list_participants" endpoint HTTP response body.
type ListParticipantsResponseBody []*ParticipantResponse

// ListAdvantagesResponseBody is the type of the "game" service
// "list_advantages" endpoint HTTP response body.
type ListAdvantagesResponseBody []*AdvantageResponse

// ListCompetitionEventKindsResponseBody is the type of the "game" service
// "list_competition_event_kinds" endpoint HTTP response body.
type ListCompetitionEventKindsResponseBody []*CompetitionEventDescriptionResponse

// ListParticipantsNotFoundResponseBody is the type of the "game" service
// "list_participants" endpoint HTTP response body for the "not_found" error.
type ListParticipantsNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListParticipantsInternalErrorResponseBody is the type of the "game" service
// "list_participants" endpoint HTTP response body for the "internal_error"
// error.
type ListParticipantsInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListAdvantagesNotFoundResponseBody is the type of the "game" service
// "list_advantages" endpoint HTTP response body for the "not_found" error.
type ListAdvantagesNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListAdvantagesInternalErrorResponseBody is the type of the "game" service
// "list_advantages" endpoint HTTP response body for the "internal_error" error.
type ListAdvantagesInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListCompetitionEventKindsNotFoundResponseBody is the type of the "game"
// service "list_competition_event_kinds" endpoint HTTP response body for the
// "not_found" error.
type ListCompetitionEventKindsNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListCompetitionEventKindsInternalErrorResponseBody is the type of the "game"
// service "list_competition_event_kinds" endpoint HTTP response body for the
// "internal_error" error.
type ListCompetitionEventKindsInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ParticipantResponse is used to define fields on response body types.
type ParticipantResponse struct {
	// Unique participant ID
	ID int `form:"Id" json:"Id" xml:"Id"`
	// Name of participant
	Name string `form:"Name" json:"Name" xml:"Name"`
	// Email of participant
	Email string `form:"Email" json:"Email" xml:"Email"`
	// Phone number of participant
	Phone string `form:"Phone" json:"Phone" xml:"Phone"`
	// Kind of participant
	Kind string `form:"Kind" json:"Kind" xml:"Kind"`
	// Role of participant
	Role string `form:"Role" json:"Role" xml:"Role"`
	// Backstory of participant
	Backstory string `form:"Backstory" json:"Backstory" xml:"Backstory"`
	// Gold of participant
	Gold int `form:"Gold" json:"Gold" xml:"Gold"`
	// Influence of participant
	Favor int `form:"Favor" json:"Favor" xml:"Favor"`
}

// AdvantageResponse is used to define fields on response body types.
type AdvantageResponse struct {
	// Unique advantage ID
	ID int `form:"Id" json:"Id" xml:"Id"`
	// Name of advantage
	Name string `form:"Name" json:"Name" xml:"Name"`
	// Description of advantage
	Description string `form:"Description" json:"Description" xml:"Description"`
	// Events that allow this advantage
	Events []string `form:"Events" json:"Events" xml:"Events"`
	// Tier of advantage
	Tier string `form:"Tier" json:"Tier" xml:"Tier"`
}

// CompetitionEventDescriptionResponse is used to define fields on response
// body types.
type CompetitionEventDescriptionResponse struct {
	// Unique competition event description ID
	ID int `form:"Id" json:"Id" xml:"Id"`
	// Name of competition event
	Name string `form:"Name" json:"Name" xml:"Name"`
	// Description of competition event
	Description string `form:"Description" json:"Description" xml:"Description"`
}

// NewListParticipantsResponseBody builds the HTTP response body from the
// result of the "list_participants" endpoint of the "game" service.
func NewListParticipantsResponseBody(res []*game.Participant) ListParticipantsResponseBody {
	body := make([]*ParticipantResponse, len(res))
	for i, val := range res {
		body[i] = marshalGameParticipantToParticipantResponse(val)
	}
	return body
}

// NewListAdvantagesResponseBody builds the HTTP response body from the result
// of the "list_advantages" endpoint of the "game" service.
func NewListAdvantagesResponseBody(res []*game.Advantage) ListAdvantagesResponseBody {
	body := make([]*AdvantageResponse, len(res))
	for i, val := range res {
		body[i] = marshalGameAdvantageToAdvantageResponse(val)
	}
	return body
}

// NewListCompetitionEventKindsResponseBody builds the HTTP response body from
// the result of the "list_competition_event_kinds" endpoint of the "game"
// service.
func NewListCompetitionEventKindsResponseBody(res []*game.CompetitionEventDescription) ListCompetitionEventKindsResponseBody {
	body := make([]*CompetitionEventDescriptionResponse, len(res))
	for i, val := range res {
		body[i] = marshalGameCompetitionEventDescriptionToCompetitionEventDescriptionResponse(val)
	}
	return body
}

// NewListParticipantsNotFoundResponseBody builds the HTTP response body from
// the result of the "list_participants" endpoint of the "game" service.
func NewListParticipantsNotFoundResponseBody(res *goa.ServiceError) *ListParticipantsNotFoundResponseBody {
	body := &ListParticipantsNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListParticipantsInternalErrorResponseBody builds the HTTP response body
// from the result of the "list_participants" endpoint of the "game" service.
func NewListParticipantsInternalErrorResponseBody(res *goa.ServiceError) *ListParticipantsInternalErrorResponseBody {
	body := &ListParticipantsInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListAdvantagesNotFoundResponseBody builds the HTTP response body from the
// result of the "list_advantages" endpoint of the "game" service.
func NewListAdvantagesNotFoundResponseBody(res *goa.ServiceError) *ListAdvantagesNotFoundResponseBody {
	body := &ListAdvantagesNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListAdvantagesInternalErrorResponseBody builds the HTTP response body
// from the result of the "list_advantages" endpoint of the "game" service.
func NewListAdvantagesInternalErrorResponseBody(res *goa.ServiceError) *ListAdvantagesInternalErrorResponseBody {
	body := &ListAdvantagesInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListCompetitionEventKindsNotFoundResponseBody builds the HTTP response
// body from the result of the "list_competition_event_kinds" endpoint of the
// "game" service.
func NewListCompetitionEventKindsNotFoundResponseBody(res *goa.ServiceError) *ListCompetitionEventKindsNotFoundResponseBody {
	body := &ListCompetitionEventKindsNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListCompetitionEventKindsInternalErrorResponseBody builds the HTTP
// response body from the result of the "list_competition_event_kinds" endpoint
// of the "game" service.
func NewListCompetitionEventKindsInternalErrorResponseBody(res *goa.ServiceError) *ListCompetitionEventKindsInternalErrorResponseBody {
	body := &ListCompetitionEventKindsInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}
