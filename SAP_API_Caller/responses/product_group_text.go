package responses

type ProductGroupText struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaterialGroup     string `json:"MaterialGroup"`
			Language          string `json:"Language"`
			MaterialGroupName string `json:"MaterialGroupName"`
			MaterialGroupText string `json:"MaterialGroupText"`
		} `json:"results"`
	} `json:"d"`
}
