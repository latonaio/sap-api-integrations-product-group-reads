package responses

type ProductGroup struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaterialGroup      string `json:"MaterialGroup"`
			AuthorizationGroup string `json:"AuthorizationGroup"`
			ToProductGroupText struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Text"`
		} `json:"results"`
	} `json:"d"`
}
