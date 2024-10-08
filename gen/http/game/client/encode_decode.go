// Code generated by goa v3.18.2, DO NOT EDIT.
//
// game HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/Adambombtastic/grandgame/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	game "github.com/Adambombtastic/grandgame/gen/game"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildListParticipantsRequest instantiates a HTTP request object with method
// and path set to call the "game" service "list_participants" endpoint
func (c *Client) BuildListParticipantsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListParticipantsGamePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("game", "list_participants", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListParticipantsResponse returns a decoder for responses returned by
// the game list_participants endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeListParticipantsResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//   - error: internal error
func DecodeListParticipantsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListParticipantsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_participants", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateParticipantResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_participants", err)
			}
			res := NewListParticipantsParticipantOK(body)
			return res, nil
		case http.StatusNotFound:
			var (
				body ListParticipantsNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_participants", err)
			}
			err = ValidateListParticipantsNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_participants", err)
			}
			return nil, NewListParticipantsNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body ListParticipantsInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_participants", err)
			}
			err = ValidateListParticipantsInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_participants", err)
			}
			return nil, NewListParticipantsInternalError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("game", "list_participants", resp.StatusCode, string(body))
		}
	}
}

// BuildListAdvantagesRequest instantiates a HTTP request object with method
// and path set to call the "game" service "list_advantages" endpoint
func (c *Client) BuildListAdvantagesRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListAdvantagesGamePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("game", "list_advantages", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListAdvantagesResponse returns a decoder for responses returned by the
// game list_advantages endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeListAdvantagesResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//   - error: internal error
func DecodeListAdvantagesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListAdvantagesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_advantages", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateAdvantageResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_advantages", err)
			}
			res := NewListAdvantagesAdvantageOK(body)
			return res, nil
		case http.StatusNotFound:
			var (
				body ListAdvantagesNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_advantages", err)
			}
			err = ValidateListAdvantagesNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_advantages", err)
			}
			return nil, NewListAdvantagesNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body ListAdvantagesInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_advantages", err)
			}
			err = ValidateListAdvantagesInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_advantages", err)
			}
			return nil, NewListAdvantagesInternalError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("game", "list_advantages", resp.StatusCode, string(body))
		}
	}
}

// BuildListCompetitionEventKindsRequest instantiates a HTTP request object
// with method and path set to call the "game" service
// "list_competition_event_kinds" endpoint
func (c *Client) BuildListCompetitionEventKindsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListCompetitionEventKindsGamePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("game", "list_competition_event_kinds", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListCompetitionEventKindsResponse returns a decoder for responses
// returned by the game list_competition_event_kinds endpoint. restoreBody
// controls whether the response body should be restored after having been read.
// DecodeListCompetitionEventKindsResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//   - error: internal error
func DecodeListCompetitionEventKindsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListCompetitionEventKindsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_competition_event_kinds", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateCompetitionEventDescriptionResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_competition_event_kinds", err)
			}
			res := NewListCompetitionEventKindsCompetitionEventDescriptionOK(body)
			return res, nil
		case http.StatusNotFound:
			var (
				body ListCompetitionEventKindsNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_competition_event_kinds", err)
			}
			err = ValidateListCompetitionEventKindsNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_competition_event_kinds", err)
			}
			return nil, NewListCompetitionEventKindsNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body ListCompetitionEventKindsInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("game", "list_competition_event_kinds", err)
			}
			err = ValidateListCompetitionEventKindsInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("game", "list_competition_event_kinds", err)
			}
			return nil, NewListCompetitionEventKindsInternalError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("game", "list_competition_event_kinds", resp.StatusCode, string(body))
		}
	}
}

// unmarshalParticipantResponseToGameParticipant builds a value of type
// *game.Participant from a value of type *ParticipantResponse.
func unmarshalParticipantResponseToGameParticipant(v *ParticipantResponse) *game.Participant {
	res := &game.Participant{
		ID:        *v.ID,
		Name:      *v.Name,
		Email:     *v.Email,
		Phone:     *v.Phone,
		Kind:      *v.Kind,
		Role:      *v.Role,
		Backstory: *v.Backstory,
		Gold:      *v.Gold,
		Favor:     *v.Favor,
	}

	return res
}

// unmarshalAdvantageResponseToGameAdvantage builds a value of type
// *game.Advantage from a value of type *AdvantageResponse.
func unmarshalAdvantageResponseToGameAdvantage(v *AdvantageResponse) *game.Advantage {
	res := &game.Advantage{
		ID:          *v.ID,
		Name:        *v.Name,
		Description: *v.Description,
		Tier:        *v.Tier,
	}
	res.Events = make([]string, len(v.Events))
	for i, val := range v.Events {
		res.Events[i] = val
	}

	return res
}

// unmarshalCompetitionEventDescriptionResponseToGameCompetitionEventDescription
// builds a value of type *game.CompetitionEventDescription from a value of
// type *CompetitionEventDescriptionResponse.
func unmarshalCompetitionEventDescriptionResponseToGameCompetitionEventDescription(v *CompetitionEventDescriptionResponse) *game.CompetitionEventDescription {
	res := &game.CompetitionEventDescription{
		ID:          *v.ID,
		Name:        *v.Name,
		Description: *v.Description,
	}

	return res
}
