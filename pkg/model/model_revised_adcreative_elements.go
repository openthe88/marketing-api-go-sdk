/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改后的创意元素，填写要求同adcreative_elements，当且仅当 online_enabled = true 时，此字段允许写入
type RevisedAdcreativeElements struct {
	Image             *string                            `json:"image,omitempty"`
	Image2            *string                            `json:"image2,omitempty"`
	Image3            *string                            `json:"image3,omitempty"`
	Title             *string                            `json:"title,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Corporate         *AdcreativeCorporate               `json:"corporate,omitempty"`
	Video             *string                            `json:"video,omitempty"`
	DeepLinkType      *string                            `json:"deep_link_type,omitempty"`
	LinkNameType      LinkNameTypeMp                     `json:"link_name_type,omitempty"`
	ImageList         *[]string                          `json:"image_list,omitempty"`
	ElementStory      *[]AdcreativeElementStoryArrayItem `json:"element_story,omitempty"`
	Url               *string                            `json:"url,omitempty"`
	ButtonText        *string                            `json:"button_text,omitempty"`
	BottomText        *string                            `json:"bottom_text,omitempty"`
	CountdownBegin    *int64                             `json:"countdown_begin,omitempty"`
	Countdown         *int64                             `json:"countdown,omitempty"`
	CountdownPrice    *string                            `json:"countdown_price,omitempty"`
	CountdownTimeType AdCreativeCountdownTimeType        `json:"countdown_time_type,omitempty"`
	MiniProgramId     *string                            `json:"mini_program_id,omitempty"`
	MiniProgramPath   *string                            `json:"mini_program_path,omitempty"`
	Label             *[]CreativeLabel                   `json:"label,omitempty"`
	ProductTags       *[]string                          `json:"product_tags,omitempty"`
	LogoDescription   *string                            `json:"logo_description,omitempty"`
	Logo              *string                            `json:"logo,omitempty"`
	LeftBottomTxt     *string                            `json:"left_bottom_txt,omitempty"`
	AnimationEffect   *string                            `json:"animation_effect,omitempty"`
	Phone             *string                            `json:"phone,omitempty"`
	ShortVideoStruct  *ShortVideoStruct                  `json:"short_video_struct,omitempty"`
	LongVideoStruct   *LongVideoStruct                   `json:"long_video_struct,omitempty"`
	BannerContent     *AdcreativeBannerContent           `json:"banner_content,omitempty"`
	CardContent       *AdcreativeCardContent             `json:"card_content,omitempty"`
	VideoPopupUrl     *string                            `json:"video_popup_url,omitempty"`
	VideoPopupButton  *AdcreativeVideoPopupButton        `json:"video_popup_button,omitempty"`
	ButtonUrl         *string                            `json:"button_url,omitempty"`
	Brand             *AdCreativeBrand                   `json:"brand,omitempty"`
	Caption           *string                            `json:"caption,omitempty"`
	LabelledImg       *AdcreativeLabelledImg             `json:"labelled_img,omitempty"`
	FullScreenImage   *string                            `json:"full_screen_image,omitempty"`
	ZipUrl            *string                            `json:"zip_url,omitempty"`
	EndPage           *AdCreativeEndPage                 `json:"end_page,omitempty"`
	ShopImage         *string                            `json:"shop_image,omitempty"`
	HeadLine          *string                            `json:"head_line,omitempty"`
}
