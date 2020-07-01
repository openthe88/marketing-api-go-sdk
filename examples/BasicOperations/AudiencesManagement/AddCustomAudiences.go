/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type CustomAudiencesAddExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.CustomAudiencesAddRequest
}

func (e *CustomAudiencesAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.CustomAudiencesAddRequest{
		AccountId:   int64(0),
		Name:        "SDK人群5ede252a51ce5",
		Description: "SDK sample",
		Type_:       model.AudienceType_CUSTOMER_FILE,
	}
}

func (e *CustomAudiencesAddExample) RunExample() (interface{}, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.CustomAudiences().Add(ctx, e.Data)
}

func main() {
	e := &CustomAudiencesAddExample{}
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
