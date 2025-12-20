package xraymanager

import (
	"context"
	"fmt"
	"os"

	"github.com/go-viper/mapstructure/v2"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/errors"
	v1 "github.com/quickpowered/thebeyond-master/pkg/xraynode/v1"
	"github.com/spf13/viper"
)

type Repository interface {
	GetGroupedNodes(region dto.Region) (map[dto.GroupType][]*dto.Node, error)
	AddClient(ctx context.Context, region dto.Region, id string, email string) error
	RemoveClient(ctx context.Context, region dto.Region, email string) error
}

type repository struct {
	groups dto.Groups
}

func New(path string) (Repository, error) {
	vp := viper.New()
	vp.SetConfigFile(path)
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var rawGroups map[dto.Region]map[dto.GroupType][]map[string]any
	if err := vp.Unmarshal(&rawGroups); err != nil {
		return nil, err
	}

	nodeByIP := make(map[string]*dto.Node)
	groups := make(dto.Groups)

	for region, nodeTypes := range rawGroups {
		if groups[region] == nil {
			groups[region] = make(map[dto.GroupType][]*dto.Node)
		}

		for nodeType, nodeList := range nodeTypes {
			for _, rawNode := range nodeList {
				var node dto.Node
				if err := mapstructure.Decode(rawNode, &node); err != nil {
					return nil, fmt.Errorf("decode node: %w", err)
				}

				key := node.IP
				existingNode, exists := nodeByIP[key]
				if exists {
					groups[region][nodeType] = append(groups[region][nodeType], existingNode)
					continue
				}

				node.NewGRPCClient(fmt.Sprintf("http://%s:%d", node.IP, 2053), os.Getenv("AUTH_SECRET"))

				nodeByIP[key] = &node
				groups[region][nodeType] = append(groups[region][nodeType], &node)
			}
		}
	}

	return repository{groups}, nil
}

func (r repository) GetGroupedNodes(region dto.Region) (map[dto.GroupType][]*dto.Node, error) {
	group, ok := r.groups[region]
	if !ok {
		return nil, errors.ErrRegionNotFound
	}

	return group, nil
}

func (r repository) AddClient(ctx context.Context, region dto.Region, id string, email string) error {
	addedToNodes := map[string]struct{}{}
	for _, nodes := range r.groups[region] {
		for _, node := range nodes {
			if _, ok := addedToNodes[node.IP]; !ok {
				if _, err := node.Client().AddClient(ctx, &v1.AddClientRequest{
					Id:    id,
					Email: email,
				}); err != nil {
					return err
				}
				addedToNodes[node.IP] = struct{}{}
			}
		}
	}

	return nil
}

func (r repository) RemoveClient(ctx context.Context, region dto.Region, email string) error {
	addedToNodes := map[string]struct{}{}
	for _, nodes := range r.groups[region] {
		for _, node := range nodes {
			if _, ok := addedToNodes[node.IP]; !ok {
				if _, err := node.Client().RemoveClient(ctx, &v1.RemoveClientRequest{
					Email: email,
				}); err != nil {
					return err
				}
				addedToNodes[node.IP] = struct{}{}
			}
		}
	}

	return nil
}
