package faceitgo

type (
	GameAssets struct {
		Cover        string `json:"cover"`
		FeaturedImgL string `json:"featured_img_l"`
		FeaturedImgM string `json:"featured_img_m"`
		FeaturedImgS string `json:"featured_img_s"`
		FlagImgIcon  string `json:"flag_img_icon"`
		FlagImgL     string `json:"flag_img_l"`
		FlagImgM     string `json:"flag_img_m"`
		FlagImgS     string `json:"flag_img_s"`
		LandingPage  string `json:"landing_page"`
	}

	Game struct {
		Assets       GameAssets `json:"assets"`
		GameID       string     `json:"game_id"`
		LongLabel    string     `json:"long_label"`
		Order        int        `json:"order"`
		ParentGameID string     `json:"parent_game_id"`
		Platforms    []string   `json:"platforms"`
		Regions      []string   `json:"regions"`
		ShortLabel   string     `json:"short_label"`
	}
)
