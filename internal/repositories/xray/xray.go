package xray

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/quickpowered/frilly/internal/repositories/xray/dto"
	handler "github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/protocol"
	cserial "github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/vless"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Interface interface {
	AddUser(email, id string) error
	RemoveUser(email string) error
}

type client struct {
	client     handler.HandlerServiceClient
	inboundTag string
	configPath string
}

func New(apiAddr, inboundTag, configPath string) (Interface, error) {
	conn, err := grpc.NewClient(apiAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithIdleTimeout(10*time.Second))
	if err != nil {
		return nil, fmt.Errorf("dial api: %w", err)
	}

	return client{handler.NewHandlerServiceClient(conn), inboundTag, configPath}, nil
}

func (s client) AddUser(email, id string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if _, err := s.client.AlterInbound(ctx, &handler.AlterInboundRequest{
		Tag: s.inboundTag,
		Operation: cserial.ToTypedMessage(&handler.AddUserOperation{
			User: &protocol.User{
				Email: email,
				Level: 0,
				Account: cserial.ToTypedMessage(&vless.Account{
					Id: id,
				}),
			},
		}),
	}); err != nil {
		return fmt.Errorf("AlterInbound failed: %w", err)
	}

	rewriteConfig(s.configPath, s.inboundTag, email, id)

	return nil
}

func (s client) RemoveUser(email string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if _, err := s.client.AlterInbound(ctx, &handler.AlterInboundRequest{
		Tag: s.inboundTag,
		Operation: cserial.ToTypedMessage(&handler.RemoveUserOperation{
			Email: email,
		}),
	}); err != nil {
		return fmt.Errorf("AlterInbound failed: %w", err)
	}

	removeFromConfig(s.configPath, s.inboundTag, email)

	return nil
}

func rewriteConfig(path, inboundTag, email, id string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	cfg := dto.Config{}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return err
	}

	for i := range cfg.Inbounds {
		if cfg.Inbounds[i].Tag == inboundTag {
			clients := cfg.Inbounds[i].Settings.Clients
			clients = append(clients, dto.Client{
				Email: email,
				Id:    id,
				Level: 2,
			})
			cfg.Inbounds[i].Settings.Clients = clients
		}
	}

	tmpFile := path + ".tmp"

	data, _ := json.MarshalIndent(cfg, "", "  ")
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmpFile, path)
}

func removeFromConfig(path, inboundTag, email string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	cfg := dto.Config{}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return err
	}

	for i := range cfg.Inbounds {
		if cfg.Inbounds[i].Tag == inboundTag {
			clients := cfg.Inbounds[i].Settings.Clients
			var newClients []dto.Client
			for _, client := range clients {
				if client.Email != email {
					newClients = append(newClients, client)
				}
			}
			cfg.Inbounds[i].Settings.Clients = newClients
		}
	}

	tmpFile := path + ".tmp"

	data, _ := json.MarshalIndent(cfg, "", "  ")
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmpFile, path)
}
