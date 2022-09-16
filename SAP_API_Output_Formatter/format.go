package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-product-group-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToProductGroup(raw []byte, l *logger.Logger) ([]ProductGroup, error) {
	pm := &responses.ProductGroup{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductGroup. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	productGroup := make([]ProductGroup, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		productGroup = append(productGroup, ProductGroup{
			MaterialGroup:      data.MaterialGroup,
			AuthorizationGroup: data.AuthorizationGroup,
			ToProductGroupText: data.ToProductGroupText.Deferred.URI,
		})
	}

	return productGroup, nil
}

func ConvertToProductGroupText(raw []byte, l *logger.Logger) ([]ProductGroupText, error) {
	pm := &responses.ProductGroupText{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductGroupText. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	productGroupText := make([]ProductGroupText, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		productGroupText = append(productGroupText, ProductGroupText{
			MaterialGroup:     data.MaterialGroup,
			Language:          data.Language,
			MaterialGroupName: data.MaterialGroupName,
			MaterialGroupText: data.MaterialGroupText,
		})
	}

	return productGroupText, nil
}

func ConvertToToProductGroupText(raw []byte, l *logger.Logger) ([]ToProductGroupText, error) {
	pm := &responses.ToProductGroupText{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToProductGroupText. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toProductGroupText := make([]ToProductGroupText, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toProductGroupText = append(toProductGroupText, ToProductGroupText{
			MaterialGroup:     data.MaterialGroup,
			Language:          data.Language,
			MaterialGroupName: data.MaterialGroupName,
			MaterialGroupText: data.MaterialGroupText,
		})
	}

	return toProductGroupText, nil
}
