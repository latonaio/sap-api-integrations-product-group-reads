package sap_api_output_formatter

type ProductGroup struct {
	MaterialGroup      string `json:"MaterialGroup"`
	AuthorizationGroup string `json:"AuthorizationGroup"`
	ToProductGroupText string `json:"to_Text"`
}

type ProductGroupText struct {
	MaterialGroup     string `json:"MaterialGroup"`
	Language          string `json:"Language"`
	MaterialGroupName string `json:"MaterialGroupName"`
	MaterialGroupText string `json:"MaterialGroupText"`
}

type ToProductGroupText struct {
	MaterialGroup     string `json:"MaterialGroup"`
	Language          string `json:"Language"`
	MaterialGroupName string `json:"MaterialGroupName"`
	MaterialGroupText string `json:"MaterialGroupText"`
}
