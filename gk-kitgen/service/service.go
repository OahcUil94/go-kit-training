package profilesvc

import "context"

type Service interface {
	PostProfile(ctx context.Context, p Profile) error
	// ...
}

type Profile struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	// ...
}
