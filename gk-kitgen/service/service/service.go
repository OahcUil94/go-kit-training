package service

import (
	"context"
	"errors"
)

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}
type Service struct {
}

func (s Service) PostProfile(ctx context.Context, p Profile) error {
	panic(errors.New("not implemented"))
}
