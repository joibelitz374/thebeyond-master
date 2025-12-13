package subscriptions

import (
	"fmt"
	"strconv"

	"github.com/quickpowered/thebeyond-master/internal/domain"
)

func (s *subscriptions) convertKeys(raw map[string]domain.Subscription) error {
	for k, v := range raw {
		n, err := strconv.Atoi(k)
		if err != nil {
			return fmt.Errorf("invalid subscription key %q: %w", k, err)
		}
		s.Subscriptions[n] = v
	}
	return nil
}
