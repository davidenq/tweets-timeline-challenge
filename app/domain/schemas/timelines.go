package schemas

import (
	"github.com/go-playground/validator/v10"
)

type TimelinesSchema []struct {
	Text             string `json:"text,omitempty"`
	ExtendedEntities struct {
		Media []struct {
			ID            int64  `json:"id,omitempty"`
			IDStr         string `json:"id_str,omitempty"`
			MediaURL      string `json:"media_url,omitempty"`
			MediaURLHTTPS string `json:"media_url_https,omitempty"`
			URL           string `json:"url,omitempty"`
			DisplayURL    string `json:"display_url,omitempty"`
			ExpandedURL   string `json:"expanded_url,omitempty"`
			Type          string `json:"type,omitempty"`
			VideoInfo     struct {
				AspectRatio []int `json:"aspect_ratio,omitempty"`
			} `json:"video_info,omitempty"`
			AdditionalMediaInfo struct {
				SourceUser struct {
					ID                             int64  `json:"id,omitempty"`
					IDStr                          string `json:"id_str,omitempty"`
					Name                           string `json:"name,omitempty"`
					ScreenName                     string `json:"screen_name,omitempty"`
					Description                    string `json:"description,omitempty"`
					ProfileBackgroundImageURL      string `json:"profile_background_image_url,omitempty"`
					ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https,omitempty"`
					ProfileImageURL                string `json:"profile_image_url,omitempty"`
					ProfileImageURLHTTPS           string `json:"profile_image_url_https,omitempty"`
					ProfileBannerURL               string `json:"profile_banner_url,omitempty"`
				} `json:"source_user,omitempty"`
			} `json:"additional_media_info,omitempty"`
		} `json:"media,omitempty"`
	} `json:"extended_entities,omitempty"`
	User struct {
		ID                             int64  `json:"id,omitempty"`
		IDStr                          string `json:"id_str,omitempty"`
		Name                           string `json:"name,omitempty"`
		ScreenName                     string `json:"screen_name,omitempty"`
		Description                    string `json:"description,omitempty"`
		ProfileBackgroundImageURL      string `json:"profile_background_image_url,omitempty"`
		ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https,omitempty"`
		ProfileImageURL                string `json:"profile_image_url,omitempty"`
		ProfileImageURLHTTPS           string `json:"profile_image_url_https,omitempty"`
		ProfileBannerURL               string `json:"profile_banner_url,omitempty"`
	} `json:"user,omitempty"`
}

func (s *TimelinesSchema) Validate(validate validator.Validate) error {
	return nil
}
