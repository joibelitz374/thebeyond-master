package components

const REGION_COMPONENT = "region"

type RegionComponent struct {
	text map[string][]string
}

func NewRegionComponent() *RegionComponent {
	text := map[string][]string{
		"en": {
			"Choose your region",
			"Region changed to",
		},
		"de": {},
		"nl": {},
		"sv": {},
		"no": {},
		"da": {},
		"es": {},
		"fr": {},
		"pt": {},
		"it": {},
		"ru": {},
		"ua": {},
		"pl": {},
		"cs": {},
		"bg": {},
		"sr": {},
		"hr": {},
		"sk": {},
		"sl": {},
		"lt": {},
		"lv": {},
		"et": {},
		"fi": {},
		"el": {},
		"ro": {},
		"hu": {},
		"ar": {},
		"fa": {},
		"tr": {},
		"he": {},
		"zh": {},
		"ja": {},
		"ko": {},
		"vi": {},
		"th": {},
		"id": {},
		"ms": {},
		"tl": {},
		"hi": {},
		"ur": {},
		"bn": {},
		"ta": {},
		"te": {},
		"mr": {},
	}

	return &RegionComponent{text}
}

func (c *RegionComponent) Text(language string) []string {
	return c.text[language]
}
