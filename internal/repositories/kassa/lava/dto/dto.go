package dto

type AvailableTariffs struct {
	Data        []AvailableTariffsData `json:"data"`
	Error       `json:"error"`
	Status      int  `json:"status"`
	StatusCheck bool `json:"status_check"`
}

type AvailableTariffsData struct {
	Currency           string  `json:"currency"`
	DiscountFromAmount float64 `json:"discount_from_amount"`
	DiscountPercent    float64 `json:"discount_percent"`
	FixCommission      float64 `json:"fix_commission"`
	MaxAmount          float64 `json:"max_amount"`
	MinAmount          float64 `json:"min_amount"`
	Percent            float64 `json:"percent"`
	ServiceID          string  `json:"service_id"`
	ServiceName        string  `json:"service_name"`
	ShopPercent        float64 `json:"shop_percent"`
	Status             int     `json:"status"`
	UserPercent        float64 `json:"user_percent"`
}

type Invoice struct {
	Data        InvoiceData `json:"data"`
	Error       `json:"error"`
	Status      int  `json:"status"`
	StatusCheck bool `json:"status_check"`
}

type InvoiceData struct {
	ID             string   `json:"id"`
	Amount         int      `json:"amount"`
	Expired        string   `json:"expired"`
	Status         int      `json:"status"`
	ShopID         string   `json:"shop_id"`
	URL            string   `json:"url"`
	Comment        string   `json:"comment"`
	FailURL        string   `json:"fail_url"`
	SuccessURL     string   `json:"success_url"`
	HookURL        string   `json:"hook_url"`
	CustomFields   string   `json:"custom_fields"`
	MerchantName   string   `json:"merchantName"`
	ExcludeService []string `json:"exclude_service"`
	IncludeService []string `json:"include_service"`
}

type Error struct {
	IncludeService []string `json:"includeService.0"`
}
