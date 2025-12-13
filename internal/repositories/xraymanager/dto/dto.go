package dto

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/quickpowered/thebeyond-master/pkg/xraynode/v1/v1connect"
)

type ClusterID int
type ListType string

const (
	NodeTypeBlacklist ListType = "blacklists"
	NodeTypeWhitelist ListType = "whitelists"
)

type Cluster map[ListType][]*Node

type Node struct {
	conn      v1connect.NodeServiceClient
	Name      string `mapstructure:"name" yaml:"name"`
	Flag      string `mapstructure:"flag" yaml:"flag"`
	IP        string `mapstructure:"ip" yaml:"ip"`
	Port      int    `mapstructure:"port" yaml:"port"`
	SNI       string `mapstructure:"sni" yaml:"sni"`
	PublicKey string `mapstructure:"pbk" yaml:"pbk"`
}

func (n *Node) NewGRPCAPI(hostname, authSecret string) {
	authInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if authSecret != "" {
				req.Header().Set("Authorization", "Bearer "+authSecret)
			}
			return next(ctx, req)
		}
	})

	n.conn = v1connect.NewNodeServiceClient(
		http.DefaultClient,
		hostname,
		connect.WithInterceptors(authInterceptor),
	)
}

func (n *Node) Client() v1connect.NodeServiceClient {
	return n.conn
}
