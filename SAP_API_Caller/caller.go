package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-product-group-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetProductGroup(materialGroup, language, materialGroupName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ProductGroup":
			func() {
				c.ProductGroup(materialGroup)
				wg.Done()
			}()
		case "ProductGroupName":
			func() {
				c.ProductGroupName(language, materialGroupName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ProductGroup(materialGroup string) {
	productGroupData, err := c.callProductGroupSrvAPIRequirementProductGroup("A_ProductGroup", materialGroup)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(productGroupData)
	}

	productGroupNameData, err := c.callToProductGroupName(productGroupData[0].ToProductGroupText)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(productGroupNameData)
	}
	return
}

func (c *SAPAPICaller) callProductGroupSrvAPIRequirementProductGroup(api, materialGroup string) ([]sap_api_output_formatter.ProductGroup, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTGROUP_SRV", api}, "/")
	param := c.getQueryWithProductGroup(map[string]string{}, materialGroup)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProductGroup(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductGroupName(url string) ([]sap_api_output_formatter.ToProductGroupText, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductGroupText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ProductGroupName(language, materialGroupName string) {
	data, err := c.callProductGroupSrvAPIRequirementProductGroupName("A_ProductGroupText", language, materialGroupName)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(data)
	}
	return
}

func (c *SAPAPICaller) callProductGroupSrvAPIRequirementProductGroupName(api, language, materialGroupName string) ([]sap_api_output_formatter.ProductGroupText, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTGROUP_SRV", api}, "/")

	param := c.getQueryWithProductGroupName(map[string]string{}, language, materialGroupName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProductGroupText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithProductGroup(params map[string]string, materialGroup string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("MaterialGroup eq '%s'", materialGroup)
	return params
}

func (c *SAPAPICaller) getQueryWithProductGroupName(params map[string]string, language, materialGroupName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and substringof('%s', MaterialGroupName)", language, materialGroupName)
	return params
}
