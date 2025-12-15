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
	errorsvar "github.com/quickpowered/thebeyond-master/pkg/errors"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
)

func (h subscription) Default(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	subID := c.Params("sub_id")
	if subID == "r" {
		return c.Redirect().To(c.Query("url"))
	}

	account, err := h.accountService.GetByKeyID(ctx, subID)
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

	subscription := "#profile-title: base64:8J+SmyBCZXlvbmQgU2VjdXJl" +
		"\n#profile-update-interval: 12" +
		"\n#profile-web-page-url: https://t.me/beyondsecurenews" +
		"\n#support-url: https://t.me/beyondsecurenews?direct" +
		"\n#notification-subs-expire: 1" +
		"\n#hide-settings: true" +
		fmt.Sprintf("\n#subscription-userinfo: expire=%v", expire)

	lists := []xmdto.ListType{xmdto.NodeTypeBlacklist}
	if account.WhitelistExpiresAt != nil {
		lists = append(lists, xmdto.NodeTypeWhitelist)
	}

	for _, list := range lists {
		nodes, err := h.xraymanagerRepo.GetNodes(ctx, xmdto.ClusterID(account.ClusterID), list)
		if err != nil {
			if errors.Is(err, errorsvar.ErrListNodesNotFound) {
				continue
			}

			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		for _, node := range nodes {
			subscription += "\n" + utils.GenerateVLESSURI(node, account.KeyID, node.PublicKey, account.ShortID)
		}
	}

	return c.SendString(subscription)
}

func (h subscription) Smart(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	if region := c.Params("region"); region != "ru" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Region not found",
		})
	}

	account, err := h.accountService.GetByKeyID(ctx, c.Params("sub_id"))
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

	c.Set("Content-Type", "application/json")
	c.Set("Profile-Title", "base64:8J+SmyBCZXlvbmQgU2VjdXJl")
	c.Set("Profile-Update-Interval", "12")
	c.Set("Profile-Web-Page-URL", "https://t.me/beyondsecurenews")
	c.Set("Support-URL", "https://t.me/beyondsecurenews?direct")
	c.Set("Notification-Subs-Expire", "1")
	c.Set("Hide-Settings", "true")
	c.Set("Subscription-Userinfo", fmt.Sprintf("expire=%v", expire))

	lists := []xmdto.ListType{xmdto.NodeTypeBlacklist}
	if account.WhitelistExpiresAt != nil {
		lists = append(lists, xmdto.NodeTypeWhitelist)
	}

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
			SubjectSelector: []string{"proxy"},
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
					Selector: []string{"proxy"},
					Strategy: dto.Strategy{
						Type: "leastPing",
					},
					FallbackTag: "direct",
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
				// 	Domain:      blockedDomains,
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

	for _, list := range lists {
		nodes, err := h.xraymanagerRepo.GetNodes(ctx, xmdto.ClusterID(account.ClusterID), list)
		if err != nil {
			if errors.Is(err, errorsvar.ErrListNodesNotFound) {
				continue
			}

			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		if len(nodes) == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Nodes not found",
			})
		}

		node := nodes[0]
		clientConfig.Outbounds = []dto.Outbound{
			{
				Tag:      "proxy",
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
			},
			{
				Tag:      "direct",
				Protocol: "freedom",
				Settings: dto.OutboundSettings{},
			},
			{
				Tag:      "block",
				Protocol: "blackhole",
				Settings: dto.OutboundSettings{},
			},
		}

		break
	}

	return c.JSON(clientConfig)
}
