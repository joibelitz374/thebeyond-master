package utils

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func RunPeriodic(ctx context.Context, interval time.Duration, name string, logger *zap.Logger, fn func(context.Context) error) {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				jobCtx, cancel := context.WithTimeout(ctx, interval)
				if err := fn(jobCtx); err != nil {
					logger.Error(name, zap.Error(err))
				}
				cancel()
			}
		}
	}()
}
