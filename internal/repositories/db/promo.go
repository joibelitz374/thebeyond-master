package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quickpowered/frilly/internal/domain"
	"go.uber.org/zap"
)

type PromoInterface interface {
	Create(ctx context.Context, name string, creator int) error
	Get(ctx context.Context, name string) (promo domain.Promo, err error)
	GetByCreator(ctx context.Context, creator int) (promos []domain.Promo, err error)
	UpgradeLevel(ctx context.Context, name string) error
	CheckBuyer(ctx context.Context, name string, buyerID int) error
	AddBuyer(ctx context.Context, name string, buyerID int) error
	IncreaseClients(ctx context.Context, name string) error
	IncreaseBuyers(ctx context.Context, name string) error
}

type promoDB struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewPromo(pool *pgxpool.Pool, logger *zap.Logger) PromoInterface {
	return promoDB{pool, logger}
}

func (db promoDB) Create(ctx context.Context, name string, creator int) error {
	db.logger.Debug("creating promo",
		zap.String("promo_name", name),
		zap.Int("creator", creator))
	_, err := db.pool.Exec(ctx, INSERT_PROMO_SQL, name, creator)
	return err
}

func (db promoDB) Get(ctx context.Context, name string) (promo domain.Promo, err error) {
	db.logger.Debug("getting promo",
		zap.String("promo_name", name))

	err = db.pool.QueryRow(ctx, GET_PROMO_SQL, name).
		Scan(&promo.Name, &promo.Creator, &promo.Level, &promo.Clients, &promo.Buyers)
	return promo, err
}

func (db promoDB) GetByCreator(ctx context.Context, creator int) (promos []domain.Promo, err error) {
	rows, err := db.pool.Query(ctx, GET_PROMOS_SQL, creator)
	if err != nil {
		return nil, err
	}

	promos = make([]domain.Promo, 0)
	for rows.Next() {
		promo := domain.Promo{}
		if err = rows.Scan(&promo.Name, &promo.Creator, &promo.Level, &promo.Clients, &promo.Buyers); err != nil {
			return nil, err
		}
		promos = append(promos, promo)
	}

	return promos, err
}

func (db promoDB) UpgradeLevel(ctx context.Context, name string) error {
	db.logger.Debug("upgrading level",
		zap.String("promo_name", name))
	_, err := db.pool.Exec(ctx, UPGRADE_LEVEL_SQL, name)
	return err
}

func (db promoDB) CheckBuyer(ctx context.Context, name string, buyerID int) error {
	var fkBuyerID int
	db.logger.Debug("checking buyer",
		zap.String("promo_name", name),
		zap.Int("buyer_id", buyerID))
	return db.pool.QueryRow(ctx, CHECK_PROMO_BUYER_SQL, name, buyerID).Scan(&fkBuyerID)
}

func (db promoDB) AddBuyer(ctx context.Context, name string, buyerID int) error {
	db.logger.Debug("adding buyer",
		zap.String("promo_name", name),
		zap.Int("buyer_id", buyerID))

	_, err := db.pool.Exec(ctx, INSERT_PROMO_BUYER_SQL, name, buyerID)
	return err
}

func (db promoDB) IncreaseClients(ctx context.Context, name string) error {
	db.logger.Debug("increasing clients",
		zap.String("promo_name", name))
	_, err := db.pool.Exec(ctx, INCREASE_PROMO_CLIENTS_SQL, name)
	return err
}

func (db promoDB) IncreaseBuyers(ctx context.Context, name string) error {
	db.logger.Debug("increasing buyers",
		zap.String("promo_name", name))
	_, err := db.pool.Exec(ctx, INCREASE_PROMO_BUYERS_SQL, name)
	return err
}
