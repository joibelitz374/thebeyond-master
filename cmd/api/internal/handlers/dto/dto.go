package dto

type XRayClientConfig struct {
	Log              `json:"log,omitempty"`
	API              `json:"api,omitempty"`
	DNS              `json:"dns,omitempty"`
	Assets           `json:"assets,omitempty"`
	Inbounds         []Inbound  `json:"inbounds,omitempty"`
	Outbounds        []Outbound `json:"outbounds,omitempty"`
	BurstObservatory `json:"burstObservatory,omitempty"`
	Remarks          string `json:"remarks,omitempty"`
	Routing          `json:"routing,omitempty"`
}

type Log struct {
	LogLevel string `json:"loglevel,omitempty"`
	Access   string `json:"access,omitempty"`
	Error    string `json:"error"`
	DNSLog   bool   `json:"dnsLog"`
}

type API struct {
	Listen   string   `json:"listen"`
	Services []string `json:"services"`
	Tag      string   `json:"tag"`
}

type DNS struct {
	Servers       []string `json:"servers,omitempty"`
	QueryStrategy string   `json:"queryStrategy,omitempty"`
	DisableCache  bool     `json:"disableCache"`
}

type Assets struct {
	GeoIPURL   string `json:"geoipUrl,omitempty"`
	GeoSiteURL string `json:"geositeUrl,omitempty"`
}

type Inbound struct {
	Listen   string          `json:"listen,omitempty"`
	Port     int             `json:"port,omitempty"`
	Protocol string          `json:"protocol,omitempty"`
	Settings InboundSettings `json:"settings,omitempty"`
	Sniffing `json:"sniffing,omitempty"`
	Tag      string `json:"tag,omitempty"`
}

type InboundSettings struct {
	Auth string `json:"auth,omitempty"`
	UDP  bool   `json:"udp,omitempty"`
}

type Sniffing struct {
	DestOverride    []string `json:"destOverride,omitempty"`
	Enabled         bool     `json:"enabled,omitempty"`
	ExcludedDomains []string `json:"excludedDomains"`
	MetadataOnly    bool     `json:"metadataOnly"`
	RouteOnly       bool     `json:"routeOnly"`
}

type Outbound struct {
	Tag            string           `json:"tag,omitempty"`
	Protocol       string           `json:"protocol,omitempty"`
	Settings       OutboundSettings `json:"settings,omitempty"`
	StreamSettings *StreamSettings  `json:"streamSettings,omitempty"`
}

type OutboundSettings struct {
	Vnext []Vnext `json:"vnext,omitempty"`
}

type Vnext struct {
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
	Users   []User `json:"users,omitempty"`
}

type User struct {
	ID         string `json:"id,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	Security   string `json:"security,omitempty"`
}

type StreamSettings struct {
	Network         string           `json:"network,omitempty"`
	Security        string           `json:"security,omitempty"`
	RealitySettings *RealitySettings `json:"realitySettings,omitempty"`
	XHTTPSettings   *XHTTPSettings   `json:"xhttpSettings,omitempty"`
}

type RealitySettings struct {
	PublicKey   string `json:"publicKey,omitempty"`
	ServerName  string `json:"serverName,omitempty"`
	ShortID     string `json:"shortId,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Show        bool   `json:"show"`
	SpiderX     string `json:"spiderX"`
}

type XHTTPSettings struct {
	Host string `json:"host"`
	Mode string `json:"mode,omitempty"`
	Path string `json:"path,omitempty"`
}

type BurstObservatory struct {
	SubjectSelector []string   `json:"subjectSelector,omitempty"`
	PingConfig      PingConfig `json:"pingConfig,omitempty"`
}

type PingConfig struct {
	Destination  string `json:"destination,omitempty"`
	Connectivity string `json:"connectivity,omitempty"`
	Interval     string `json:"interval,omitempty"`
	Sampling     int    `json:"sampling,omitempty"`
	Timeout      string `json:"timeout,omitempty"`
}

type Routing struct {
	Balancers      []Balancer `json:"balancers,omitempty"`
	DomainStrategy string     `json:"domainStrategy,omitempty"`
	Rules          []Rule     `json:"rules,omitempty"`
}

type Balancer struct {
	Tag         string   `json:"tag,omitempty"`
	Selector    []string `json:"selector,omitempty"`
	Strategy    Strategy `json:"strategy,omitempty"`
	FallbackTag string   `json:"fallbackTag,omitempty"`
}

type Strategy struct {
	Type string `json:"type,omitempty"`
}

type Rule struct {
	Type        string   `json:"type,omitempty"`
	IP          []string `json:"ip,omitempty"`
	Protocol    []string `json:"protocol,omitempty"`
	Domain      []string `json:"domain,omitempty"`
	InboundTag  []string `json:"inboundTag,omitempty"`
	OutboundTag string   `json:"outboundTag,omitempty"`
	BalancerTag string   `json:"balancerTag,omitempty"`
}
