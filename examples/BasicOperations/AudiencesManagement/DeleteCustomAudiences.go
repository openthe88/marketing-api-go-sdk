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

type CustomAudiencesDeleteExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.CustomAudiencesDeleteRequest
}

func (e *CustomAudiencesDeleteExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.CustomAudiencesDeleteRequest{
		AccountId:  int64(0),
		AudienceId: int64(0),
	}
}

func (e *CustomAudiencesDeleteExample) RunExample() (interface{}, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.CustomAudiences().Delete(ctx, e.Data)
}

func main() {
	e := &CustomAudiencesDeleteExample{}
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
