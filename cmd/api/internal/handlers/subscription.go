package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/thebeyond-master/cmd/api/internal/handlers/dto"
	"github.com/quickpowered/thebeyond-master/cmd/api/pkg/domains"
	xmdto "github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
)

func (h subscription) Default(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	keyID := c.Params("key_id")
	if keyID == "r" {
		return c.Redirect().To(c.Query("url"))
	}

	account, err := h.accountService.GetByKeyID(ctx, keyID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Subscription not found",
			})
		} else {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	expire := int(time.Now().Unix())
	if account.SubscriptionExpiresAt != nil {
		subscriptionExpiresAt := *account.SubscriptionExpiresAt
		expire = int(subscriptionExpiresAt.Unix())
	}

	h.setAppHeaders(c, expire, account.UsedUplink, account.UsedDownlink)

	groupedNodes, err := h.xraymanagerRepo.GetGroupedNodes(xmdto.Region(account.Region))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	nodeByCountryCodes := make(map[string]struct{})

	var subscription string
	for _, nodes := range groupedNodes {
		for _, node := range nodes {
			if _, ok := nodeByCountryCodes[node.CountryCode]; ok {
				continue
			}

			nodeByCountryCodes[node.CountryCode] = struct{}{}
			subscription += utils.GenerateVLESSURI(node, account.KeyID, node.PublicKey, account.ShortID) + "\n"
		}
	}

	return c.SendString(subscription)
}

func (h subscription) Smart(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	keyID := c.Params("key_id")
	if keyID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Key ID is required",
		})
	}

	account, err := h.accountService.GetByKeyID(ctx, keyID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Account not found",
			})
		} else {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	region := c.Params("region")
	if region == "" {
		region = account.Region
	}

	if region != "ru" && region != "ruwh" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Region not found",
		})
	}

	expire := int(time.Now().Unix())
	if account.SubscriptionExpiresAt != nil {
		subscriptionExpiresAt := *account.SubscriptionExpiresAt
		expire = int(subscriptionExpiresAt.Unix())
	}

	c.Set("Content-Type", "application/json")
	h.setAppHeaders(c, expire, account.UsedUplink, account.UsedDownlink)

	clientConfig := dto.XRayClientConfig{
		Log: dto.Log{
			LogLevel: "warning",
			Access:   "none",
			Error:    "",
			DNSLog:   false,
		},
		API: dto.API{
			Listen:   "127.0.0.1:10085",
			Services: []string{"HandlerService", "StatsService"},
			Tag:      "api",
		},
		DNS: dto.DNS{
			Servers: []string{
				"https://8.8.8.8/dns-query",
				"https://1.1.1.1/dns-query",
				"208.67.222.222",
				"208.67.220.220",
				"https+local://doh.opendns.com/dns-query",
				"https+local://cloudflare-dns.com/dns-query",
			},
			QueryStrategy: "UseIPv4",
			DisableCache:  false,
		},
		Assets: dto.Assets{
			GeoIPURL:   "https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geoip.dat",
			GeoSiteURL: "https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geosite.dat",
		},
		Inbounds: []dto.Inbound{
			{
				Listen:   "127.0.0.1",
				Port:     10808,
				Protocol: "socks",
				Settings: dto.InboundSettings{
					Auth: "noauth",
					UDP:  true,
				},
				Sniffing: dto.Sniffing{
					DestOverride:    []string{"http", "tls", "quic"},
					Enabled:         true,
					ExcludedDomains: []string{},
					MetadataOnly:    false,
					RouteOnly:       false,
				},
				Tag: "socks-in",
			},
			{
				Listen:   "127.0.0.1",
				Port:     10809,
				Protocol: "http",
				Settings: dto.InboundSettings{
					Auth: "noauth",
					UDP:  true,
				},
				Sniffing: dto.Sniffing{
					DestOverride:    []string{"http", "tls", "quic"},
					Enabled:         true,
					ExcludedDomains: []string{},
					MetadataOnly:    false,
					RouteOnly:       false,
				},
				Tag: "http",
			},
			{
				Listen:   "127.0.0.1",
				Port:     10820,
				Protocol: "socks",
				Settings: dto.InboundSettings{
					Auth: "noauth",
					UDP:  true,
				},
				Sniffing: dto.Sniffing{
					DestOverride:    []string{"http", "tls", "quic"},
					Enabled:         true,
					ExcludedDomains: []string{},
					MetadataOnly:    false,
					RouteOnly:       false,
				},
				Tag: "socks-direct",
			},
		},
		Outbounds: []dto.Outbound{},
		BurstObservatory: dto.BurstObservatory{
			SubjectSelector: []string{"general-proxy", "fallback-proxy"},
			PingConfig: dto.PingConfig{
				Destination:  "https://www.google.com/generate_204",
				Connectivity: "",
				Interval:     "180s",
				Sampling:     5,
				Timeout:      "5s",
			},
		},
		Remarks: "🧠 SMART (BEST)",
		Routing: dto.Routing{
			Balancers: []dto.Balancer{
				{
					Tag:      "balancer",
					Selector: []string{"general-proxy"},
					Strategy: dto.Strategy{
						Type: "leastPing",
					},
					FallbackTag: "fallback-proxy",
				},
			},
			DomainStrategy: "IPIfNonMatch",
			Rules: []dto.Rule{
				{
					Type: "field",
					IP: []string{
						"8.8.8.8",
						"1.1.1.1",
						"208.67.222.222",
						"208.67.220.220",
					},
					OutboundTag: "direct",
				},
				{
					Type: "field",
					Domain: []string{
						"cloudflare-dns.com",
						"doh.opendns.com",
						"www.google.com",
						"connectivitycheck.gstatic.com",
					},
					OutboundTag: "direct",
				},
				{
					Type:        "field",
					IP:          []string{"geoip:private"},
					OutboundTag: "block",
				},
				{
					InboundTag:  []string{"socks-direct"},
					OutboundTag: "direct",
				},
				// {
				// 	Type:        "field",
				// 	Domain:      domains.Blocked,
				// 	BalancerTag: "balancer",
				// },
				{
					Type:        "field",
					Protocol:    []string{"bittorrent"},
					OutboundTag: "direct",
				},
				{
					Type:        "field",
					Domain:      domains.RU,
					OutboundTag: "direct",
				},
				{
					Type:        "field",
					IP:          []string{"0.0.0.0/0", "::/0"},
					BalancerTag: "balancer",
				},
			},
		},
	}

	groupedNodes, err := h.xraymanagerRepo.GetGroupedNodes(xmdto.Region(region))
	if err != nil {
		fmt.Println("groupedNodes error is", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if len(groupedNodes) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Nodes not found",
		})
	}

	for groupName, nodes := range groupedNodes {
		for _, node := range nodes {
			clientConfig.Outbounds = append(clientConfig.Outbounds, dto.Outbound{
				Tag:      string(groupName) + "-proxy",
				Protocol: "vless",
				Settings: dto.OutboundSettings{
					Vnext: []dto.Vnext{
						{
							Address: node.IP,
							Port:    node.Port,
							Users: []dto.User{
								{
									ID:         account.KeyID,
									Encryption: "none",
									Security:   "auto",
								},
							},
						},
					},
				},
				StreamSettings: &dto.StreamSettings{
					Network:  "xhttp",
					Security: "reality",
					RealitySettings: &dto.RealitySettings{
						PublicKey:   node.PublicKey,
						ServerName:  node.SNI,
						ShortID:     account.ShortID,
						Fingerprint: "chrome",
						Show:        false,
						SpiderX:     "",
					},
					XHTTPSettings: &dto.XHTTPSettings{
						Host: "",
						Mode: "auto",
						Path: "/",
					},
				},
			})
		}
	}

	clientConfig.Outbounds = append(clientConfig.Outbounds, dto.Outbound{
		Tag:      "direct",
		Protocol: "freedom",
		Settings: dto.OutboundSettings{},
	})

	clientConfig.Outbounds = append(clientConfig.Outbounds, dto.Outbound{
		Tag:      "block",
		Protocol: "blackhole",
		Settings: dto.OutboundSettings{},
	})

	return c.JSON(clientConfig)
}

const DecimalGB uint64 = 1_000_000_000

const (
	GiB uint64 = 1 << 30
	TiB uint64 = 1 << 40
)

var limits = map[int]uint64{
	10:   10 * GiB,
	150:  150 * GiB,
	400:  400 * GiB,
	1000: 1 * TiB,
}

func (h subscription) setAppHeaders(c fiber.Ctx, expire int, usedUplink, usedDownlink int64) {
	c.Set("Profile-Title", "base64:8J+SmyBUSEUgQkVZT05E")
	c.Set("Profile-Update-Interval", "12")
	c.Set("Profile-Web-Page-URL", "https://t.me/beyondsecurenews")
	c.Set("Support-URL", "https://t.me/beyondsecurenews?direct")
	c.Set("Notification-Subs-Expire", "1")
	c.Set("Hide-Settings", "true")
	c.Set("Subscription-Userinfo", fmt.Sprintf("upload=%d; download=%d; total=%d; expire=%v", usedUplink, usedDownlink, limits[10], expire))
}
