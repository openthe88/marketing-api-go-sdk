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
)

type CustomTagFilesGetExample struct {
	TAds                  *ads.SDKClient
	AccessToken           string
	AccountId             int64
	CustomTagFilesGetOpts *api.CustomTagFilesGetOpts
}

func (e *CustomTagFilesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.CustomTagFilesGetOpts = &api.CustomTagFilesGetOpts{

		Fields: optional.NewInterface([]string{"tag_id", "custom_tag_file_id", "name", "user_id_type", "operation_type", "open_app_id", "process_status", "process_code", "error_message", "line_count", "valid_line_count", "user_count", "created_time"}),
	}
}

func (e *CustomTagFilesGetExample) RunExample() (interface{}, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.CustomTagFiles().Get(ctx, e.AccountId, e.CustomTagFilesGetOpts)
}

func main() {
	e := &CustomTagFilesGetExample{}
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
