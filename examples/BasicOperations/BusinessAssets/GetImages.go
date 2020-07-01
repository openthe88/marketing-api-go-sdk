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

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type ImagesGetExample struct {
	TAds          *ads.SDKClient
	AccessToken   string
	AccountId     int64
	ImagesGetOpts *api.ImagesGetOpts
}

func (e *ImagesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.ImagesGetOpts = &api.ImagesGetOpts{

		Filtering: optional.NewInterface([]model.FilteringStruct{model.FilteringStruct{
			Field:    "image_id",
			Operator: "EQUALS",
			Values:   []string{"YOUR IMAGE ID"},
		}}),

		Fields: optional.NewInterface([]string{"image_id", "width", "height", "file_size", "type", "signature", "source_signature", "preview_url", "source_type", "created_time", "last_modified_time"}),
	}
}

func (e *ImagesGetExample) RunExample() (interface{}, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Images().Get(ctx, e.AccountId, e.ImagesGetOpts)
}

func main() {
	e := &ImagesGetExample{}
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
