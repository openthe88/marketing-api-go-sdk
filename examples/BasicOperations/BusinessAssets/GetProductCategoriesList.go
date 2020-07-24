/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type ProductCategoriesListGetExample struct {
	TAds                         *ads.SDKClient
	AccessToken                  string
	AccountId                    int64
	ProductCatalogId             int64
	ProductCategoriesListGetOpts *api.ProductCategoriesListGetOpts
}

func (e *ProductCategoriesListGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.ProductCatalogId = 789
	e.ProductCategoriesListGetOpts = &api.ProductCategoriesListGetOpts{}
}

func (e *ProductCategoriesListGetExample) RunExample() (model.ProductCategoriesListGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.ProductCategoriesList().Get(ctx, e.AccountId, e.ProductCatalogId, e.ProductCategoriesListGetOpts)
}

func main() {
	e := &ProductCategoriesListGetExample{}
	e.Init()
	response, httpResponse, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Http response:", httpResponse)
}
