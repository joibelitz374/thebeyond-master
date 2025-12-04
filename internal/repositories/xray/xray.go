package xray

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/quickpowered/frilly/internal/repositories/xray/dto"
	handler "github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/protocol"
	cserial "github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/vless"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var configMu sync.Mutex

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
		return fmt.Errorf("runtime API failed: %w", err)
	}

	err := updateConfigJSON(s.configPath, s.inboundTag, func(clients []interface{}) []interface{} {
		for _, c := range clients {
			if cmap, ok := c.(map[string]any); ok {
				if e, ok := cmap["email"].(string); ok && e == email {
					return clients
				}
			}
		}

		return append(clients, dto.Client{
			Email: email,
			Id:    id,
			Level: 0,
		})
	})

	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

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
		fmt.Printf("Warning: runtime user removal failed for %s: %v\n", email, err)
	}

	err := updateConfigJSON(s.configPath, s.inboundTag, func(clients []any) []any {
		var newClients []any

		for _, c := range clients {
			if cmap, ok := c.(map[string]any); ok {
				if e, ok := cmap["email"].(string); ok && e == email {
					continue
				}
			}
			newClients = append(newClients, c)
		}

		return newClients
	})

	if err != nil {
		return fmt.Errorf("failed to save config after user removal: %w", err)
	}

	return nil
}

func updateConfigJSON(path, inboundTag string, modifier func([]interface{}) []interface{}) error {
	configMu.Lock()
	defer configMu.Unlock()

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var cfg map[string]any
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	inbounds, ok := cfg["inbounds"].([]any)
	if !ok {
		return fmt.Errorf("no inbounds found")
	}

	found := false
	for i, ib := range inbounds {
		inboundMap, ok := ib.(map[string]any)
		if !ok {
			continue
		}

		if tag, ok := inboundMap["tag"].(string); ok && tag == inboundTag {
			found = true
			settings, ok := inboundMap["settings"].(map[string]any)
			if !ok {
				settings = make(map[string]any)
				inboundMap["settings"] = settings
			}

			var currentClients []any
			if c, ok := settings["clients"].([]any); ok {
				currentClients = c
			}

			settings["clients"] = modifier(currentClients)
			inbounds[i] = inboundMap
			break
		}
	}

	if !found {
		return fmt.Errorf("inbound not found")
	}

	cfg["inbounds"] = inbounds

	if err := f.Truncate(0); err != nil {
		return err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(cfg); err != nil {
		return err
	}

	return f.Sync()
}
