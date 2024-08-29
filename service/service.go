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

func (s *service) Participants(ctx context.Context) (res []*gen.Participant, err error) {
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
