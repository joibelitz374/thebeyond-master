package utils

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func RunPeriodic(ctx context.Context, interval time.Duration, name string, logger *zap.Logger, fn func(context.Context) error, immediate bool) {
	go func() {
		if immediate {
			runJob(ctx, interval, name, logger, fn)
		}

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				logger.Info("Periodic task stopped", zap.String("name", name))
				return
			case <-ticker.C:
				logger.Debug("Starting periodic job", zap.String("name", name))
				runJob(ctx, interval, name, logger, fn)
			}
		}
	}()
}

func runJob(ctx context.Context, timeout time.Duration, name string, logger *zap.Logger, fn func(context.Context) error) error {
	jobCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := fn(jobCtx)
	if err != nil {
		logger.Error("Periodic job failed", zap.String("name", name), zap.Error(err))
		return err
	}
	logger.Info("Periodic job completed", zap.String("name", name))
	return nil
}
