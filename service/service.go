package service

import (
	"context"

	gen "github.com/Adambombtastic/grandgame/gen/game"
)

type (
	service struct {
	}
)

func New() gen.Service {
	return &service{}
}

func (s *service) ListParticipants(ctx context.Context) (res []*gen.Participant, err error) {
	return []*gen.Participant{
		{
			ID:        1,
			Name:      "Adam",
			Email:     "foo@bar.com",
			Phone:     "1234567890",
			Kind:      "superadmin",
			Role:      "King",
			Backstory: "The king of the land",
			Gold:      1000,
			Favor:     100,
		},
	}, nil
}

func (s *service) ListAdvantages(ctx context.Context) (res []*gen.Advantage, err error) {
	return []*gen.Advantage{
		{
			ID:          1,
			Name:        "Advantage 1",
			Description: "The first advantage",
		},
	}, nil
}

func (s *service) ListCompetitionEventKinds(ctx context.Context) (res []*gen.CompetitionEventDescription, err error) {
	return []*gen.CompetitionEventDescription{
		{
			ID:          1,
			Name:        "Competition Event 1",
			Description: "The first competition event",
		},
	}, nil

}
