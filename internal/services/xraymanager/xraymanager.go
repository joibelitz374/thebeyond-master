package xraymanager

import (
	"context"
	"strings"
	"time"

	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"go.uber.org/zap"
)

type Interface interface {
	GetGroupedNodes(region dto.Region) (map[dto.GroupType][]*dto.Node, error)
	AddClient(ctx context.Context, region dto.Region, id string, email string) error
	RemoveClient(ctx context.Context, region dto.Region, email string) error
}

type service struct {
	repo   xraymanager.Repository
	logger *zap.Logger
}

func New(repo xraymanager.Repository, logger *zap.Logger) Interface {
	return service{repo, logger}
}

func (s service) GetGroupedNodes(region dto.Region) (map[dto.GroupType][]*dto.Node, error) {
	return s.repo.GetGroupedNodes(region)
}

func isAlreadyExistsErr(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "already exists") || (strings.Contains(msg, "user") && strings.Contains(msg, "already exists"))
}

func (s service) AddClient(ctx context.Context, region dto.Region, id string, email string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var lastErr error
	const maxAttempts = 5

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		s.logger.Info("Attempting to add client", zap.Int("attempt", attempt), zap.String("region", string(region)), zap.String("email", email))

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := s.repo.AddClient(ctx, region, id, email); err == nil {
			s.logger.Info("Successfully added client to xray", zap.String("email", email), zap.String("region", string(region)))
			return nil
		} else {
			lastErr = err
			s.logger.Warn("Failed to add client", zap.Error(err), zap.String("email", email), zap.String("region", string(region)))

			if isAlreadyExistsErr(err) {
				s.logger.Info("Detected existing user, attempting removal", zap.String("email", email), zap.String("region", string(region)))
				if remErr := s.RemoveClient(ctx, region, email); remErr != nil {
					s.logger.Error("Failed to remove existing user", zap.Error(remErr), zap.String("email", email), zap.String("region", string(region)))
				} else {
					s.logger.Info("Successfully removed existing user, waiting before retry", zap.String("email", email), zap.String("region", string(region)))
					select {
					case <-time.After(1 * time.Second):
					case <-ctx.Done():
						return ctx.Err()
					}
				}
			} else {
				s.logger.Error("Non-existence error when adding user", zap.Error(err), zap.String("email", email), zap.String("region", string(region)))
			}
		}

		if attempt < maxAttempts {
			backoff := time.Duration(attempt) * time.Second
			s.logger.Info("Backing off before next attempt", zap.Duration("backoff", backoff))
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	s.logger.Error("All attempts failed to add client", zap.Error(lastErr), zap.String("email", email), zap.String("region", string(region)))
	return lastErr
}

func (s service) RemoveClient(ctx context.Context, region dto.Region, email string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var lastErr error
	const maxAttempts = 5

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		s.logger.Info("Attempting to remove client", zap.Int("attempt", attempt), zap.String("region", string(region)), zap.String("email", email))

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := s.repo.RemoveClient(ctx, region, email); err == nil {
			s.logger.Info("Successfully removed client from xray", zap.String("email", email), zap.String("region", string(region)))
			return nil
		} else {
			lastErr = err
			s.logger.Warn("Failed to remove client", zap.Error(err), zap.String("email", email), zap.String("region", string(region)))
		}

		if attempt < maxAttempts {
			backoff := time.Duration(attempt) * time.Second
			s.logger.Info("Backing off before next remove attempt", zap.Duration("backoff", backoff))
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	s.logger.Error("All attempts failed to remove client", zap.Error(lastErr), zap.String("email", email), zap.String("region", string(region)))
	return lastErr
}
