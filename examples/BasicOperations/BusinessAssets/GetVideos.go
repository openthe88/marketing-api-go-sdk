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

type VideosGetExample struct {
	TAds          *ads.SDKClient
	AccessToken   string
	AccountId     int64
	VideosGetOpts *api.VideosGetOpts
}

func (e *VideosGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.VideosGetOpts = &api.VideosGetOpts{

		Filtering: optional.NewInterface([]model.FilteringStruct{model.FilteringStruct{
			Field:    "media_id",
			Operator: "EQUALS",
			Values:   []string{"YOUR VIDEO ID"},
		}}),

		Fields: optional.NewInterface([]string{"video_id", "width", "height", "video_frames", "video_fps", "video_codec", "video_bit_rate", "audio_codec", "audio_bit_rate", "file_size", "type", "signature", "system_status", "description", "preview_url", "created_time", "last_modified_time", "video_profile_name", "audio_sample_rate", "max_keyframe_interval", "min_keyframe_interval", "sample_aspect_ratio", "audio_profile_name", "scan_type", "image_duration_millisecond", "audio_duration_millisecond", "source_type"}),
	}
}

func (e *VideosGetExample) RunExample() (interface{}, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Videos().Get(ctx, e.AccountId, e.VideosGetOpts)
}

func main() {
	e := &VideosGetExample{}
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
