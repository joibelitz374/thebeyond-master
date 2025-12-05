package serverlocations

import "github.com/quickpowered/frilly/internal/domain"

type Repository interface {
	Get(shortName string) domain.Location
}

type repository struct {
	locations map[string]domain.Location
}

func New() (Repository, error) {
	return repository{
		locations: map[string]domain.Location{
			"de": {
				Name: "Frankfurt",
				Flag: "🇩🇪",
				IP:   "77.221.157.159",
				Port: 443,
				SNI:  "web.max.ru",
			},
		},
	}, nil
}

func (r repository) Get(shortName string) domain.Location {
	return r.locations[shortName]
}
