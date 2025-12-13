package xraymanager

import (
	"context"
	"fmt"
	"os"

	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/errors"
	v1 "github.com/quickpowered/thebeyond-master/pkg/xraynode/v1"
	"github.com/spf13/viper"
)

type Repository interface {
	GetNodes(ctx context.Context, clusterID dto.ClusterID, listType dto.ListType) ([]*dto.Node, error)
	AddClient(ctx context.Context, clusterID dto.ClusterID, nodeType dto.ListType, id string, email string) error
	RemoveClient(ctx context.Context, clusterID dto.ClusterID, nodeType dto.ListType, email string) error
}

type service struct {
	clusters map[dto.ClusterID]dto.Cluster
}

func New(path string) (Repository, error) {
	vp := viper.New()
	vp.SetConfigFile(path)
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	clusters := map[dto.ClusterID]dto.Cluster{}
	if err := vp.Unmarshal(&clusters); err != nil {
		return nil, err
	}

	for _, cluster := range clusters {
		for _, list := range cluster {
			for _, node := range list {
				node.NewGRPCAPI(fmt.Sprintf("http://%s:%d", node.IP, 2053), os.Getenv("AUTH_SECRET"))
			}
		}
	}

	return service{clusters}, nil
}

func (s service) GetNodes(ctx context.Context, clusterID dto.ClusterID, listType dto.ListType) ([]*dto.Node, error) {
	lists, ok := s.clusters[clusterID]
	if !ok {
		return nil, errors.ErrClusterNotFound
	}

	nodes, ok := lists[listType]
	if !ok {
		return nil, errors.ErrListNodesNotFound
	}

	return nodes, nil
}

func (s service) AddClient(ctx context.Context, clusterID dto.ClusterID, listType dto.ListType, id string, email string) error {
	for _, node := range s.clusters[clusterID][listType] {
		if _, err := node.Client().AddClient(ctx, &v1.AddClientRequest{
			Id:    id,
			Email: email,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s service) RemoveClient(ctx context.Context, clusterID dto.ClusterID, listType dto.ListType, email string) error {
	for _, node := range s.clusters[clusterID][listType] {
		if _, err := node.Client().RemoveClient(ctx, &v1.RemoveClientRequest{
			Email: email,
		}); err != nil {
			return err
		}
	}

	return nil
}
