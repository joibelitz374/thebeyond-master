package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/sqlc"
	"go.uber.org/zap"
)

type PromoInterface interface {
	Create(ctx context.Context, name string, creator int) error
	Get(ctx context.Context, name string) (domain.Promo, error)
	GetByCreator(ctx context.Context, creator int) ([]domain.Promo, error)
	UpgradeLevel(ctx context.Context, name string) error
	IncreaseClients(ctx context.Context, name string) error
}

type promoDB struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
	logger  *zap.Logger
}

func NewPromo(pool *pgxpool.Pool, logger *zap.Logger) PromoInterface {
	return promoDB{pool, sqlc.New(pool), logger}
}

func (db promoDB) Create(ctx context.Context, name string, creator int) error {
	db.logger.Debug("creating promo",
		zap.String("promo_name", name),
		zap.Int("creator", creator))

	return db.queries.CreatePromo(ctx, sqlc.CreatePromoParams{
		Name:    name,
		Creator: int32(creator),
	})
}

func (db promoDB) Get(ctx context.Context, name string) (domain.Promo, error) {
	db.logger.Debug("getting promo",
		zap.String("promo_name", name))

	promo, err := db.queries.GetPromo(ctx, name)
	return domain.Promo{
		Name:    promo.Name,
		Creator: int(promo.Creator),
		Level:   int(promo.Level),
		Clients: int(promo.Clients),
	}, err
}

func (db promoDB) GetByCreator(ctx context.Context, creator int) ([]domain.Promo, error) {
	rows, err := db.queries.GetPromosByCreator(ctx, int32(creator))
	if err != nil {
		return nil, err
	}

	promos := make([]domain.Promo, len(rows))
	for i, row := range rows {
		promos[i] = domain.Promo{
			Name:    row.Name,
			Creator: int(row.Creator),
			Level:   int(row.Level),
			Clients: int(row.Clients),
		}
	}

	return promos, err
}

func (db promoDB) UpgradeLevel(ctx context.Context, name string) error {
	db.logger.Debug("upgrading level",
		zap.String("promo_name", name))
	return db.queries.UpgradePromoLevel(ctx, name)
}

func (db promoDB) IncreaseClients(ctx context.Context, name string) error {
	db.logger.Debug("increasing clients",
		zap.String("promo_name", name))
	return db.queries.IncreasePromoClients(ctx, name)
}
