package dto

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/quickpowered/thebeyond-master/pkg/xraynode/v1/v1connect"
)

type Region string
type GroupType string

const (
	GroupNodesGeneral  GroupType = "general"
	GroupNodesFallback GroupType = "fallback"
)

const (
	RegionRussia                    Region = "ru"
	RegionRussiaWhitelists          Region = "ruwh"
	RegionEuropeAntiAgeVerification Region = "euav"
)

type Groups map[Region]map[GroupType][]*Node

type Node struct {
	Conn        v1connect.NodeServiceClient `mapstructure:"-" yaml:"-"`
	Name        string                      `mapstructure:"name" yaml:"name"`
	CountryCode string                      `mapstructure:"country_code" yaml:"country_code"`
	Flag        string                      `mapstructure:"flag" yaml:"flag"`
	IP          string                      `mapstructure:"ip" yaml:"ip"`
	Port        int                         `mapstructure:"port" yaml:"port"`
	SNI         string                      `mapstructure:"sni" yaml:"sni"`
	PublicKey   string                      `mapstructure:"public_key" yaml:"public_key"`
}

func (n *Node) NewGRPCClient(hostname, authSecret string) {
	authInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if authSecret != "" {
				req.Header().Set("Authorization", "Bearer "+authSecret)
			}
			return next(ctx, req)
		}
	})

	n.Conn = v1connect.NewNodeServiceClient(
		http.DefaultClient, hostname,
		connect.WithInterceptors(authInterceptor),
	)
}

func (n *Node) Client() v1connect.NodeServiceClient {
	return n.Conn
}
