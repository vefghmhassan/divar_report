package models

type Divar struct {
	InputSuggestion struct {
		JSONSchema struct {
			Properties struct {
				Districts struct {
					AdditionalProperties bool `json:"additionalProperties"`
					Properties           struct {
						NearVacancies struct {
							Title string `json:"title"`
							Type  string `json:"type"`
						} `json:"near-vacancies"`
						Vacancies struct {
							Items struct {
								Enum      []string `json:"enum"`
								EnumNames []string `json:"enumNames"`
								Type      string   `json:"type"`
							} `json:"items"`
							MinItems int    `json:"minItems"`
							Title    string `json:"title"`
							Type     string `json:"type"`
						} `json:"vacancies"`
					} `json:"properties"`
					Required []string `json:"required"`
					Title    string   `json:"title"`
					Type     string   `json:"type"`
				} `json:"districts"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"json_schema"`
		UISchema struct {
			Districts struct {
				NearVacancies struct {
					UIOptions struct {
						Enabled bool `json:"enabled"`
					} `json:"ui:options"`
				} `json:"near-vacancies"`
				UIField   string `json:"ui:field"`
				UIOptions struct {
					DisplayTextFormat  string `json:"display-text-format"`
					Icon               string `json:"icon"`
					IsAdvanced         bool   `json:"is-advanced"`
					Label              string `json:"label"`
					ParentCategorySlug string `json:"parent-category-slug"`
					PostSetLabel       string `json:"post-set-label"`
					SelectInMap        struct {
						Enabled bool `json:"enabled"`
					} `json:"select_in_map"`
				} `json:"ui:options"`
				Urischema struct {
					Order int `json:"order"`
				} `json:"urischema"`
				Vacancies struct {
					Items struct {
					} `json:"items"`
					UIOptions struct {
						Data struct {
							Children []struct {
								Enum     string   `json:"enum"`
								EnumName string   `json:"enumName"`
								Tags     []string `json:"tags"`
							} `json:"children"`
							Enum     string `json:"enum"`
							EnumName string `json:"enumName"`
						} `json:"data"`
					} `json:"ui:options"`
				} `json:"vacancies"`
			} `json:"districts"`
		} `json:"ui_schema"`
	} `json:"input_suggestion"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	SeoDetails struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Headline    string `json:"headline"`
		MetaRobots  struct {
			Follow bool `json:"follow"`
			Index  bool `json:"index"`
		} `json:"meta_robots"`
		BreadCrumbs []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"bread_crumbs"`
		Next       string        `json:"next"`
		Prev       string        `json:"prev"`
		WebURL     string        `json:"web_url"`
		AndroidURL string        `json:"android_url"`
		Canonical  string        `json:"canonical"`
		MetaTags   []interface{} `json:"meta_tags"`
	} `json:"seo_details"`
	InternalLinkSections interface{} `json:"internal_link_sections"`
	WidgetList           []struct {
		WidgetType string `json:"widget_type"`
		Data       struct {
			Title    string `json:"title"`
			Image    string `json:"image"`
			WebImage []struct {
				Src  string `json:"src"`
				Type string `json:"type"`
			} `json:"web_image"`
			Description     string      `json:"description"`
			HasChat         bool        `json:"has_chat"`
			RedText         string      `json:"red_text"`
			NormalText      string      `json:"normal_text"`
			Token           string      `json:"token"`
			ImageOverlayTag interface{} `json:"image_overlay_tag"`
			ImageTopLeftTag interface{} `json:"image_top_left_tag"`
			Index           int         `json:"index"`
			City            string      `json:"city"`
			District        string      `json:"district"`
			Category        string      `json:"category"`
		} `json:"data"`
	} `json:"widget_list"`
	LastPostDate  int64 `json:"last_post_date"`
	FirstPostDate int64 `json:"first_post_date"`
	WebWidgets    struct {
		Toolbox  []interface{} `json:"toolbox"`
		PostList []struct {
			WidgetType string `json:"widget_type"`
			Data       struct {
				Title    string `json:"title"`
				Image    string `json:"image"`
				WebImage []struct {
					Src  string `json:"src"`
					Type string `json:"type"`
				} `json:"web_image"`
				Description     string      `json:"description"`
				HasChat         bool        `json:"has_chat"`
				RedText         string      `json:"red_text"`
				NormalText      string      `json:"normal_text"`
				Token           string      `json:"token"`
				ImageOverlayTag interface{} `json:"image_overlay_tag"`
				ImageTopLeftTag interface{} `json:"image_top_left_tag"`
				Index           int         `json:"index"`
				City            string      `json:"city"`
				District        string      `json:"district"`
				Category        string      `json:"category"`
			} `json:"data"`
		} `json:"post_list"`
	} `json:"web_widgets"`
	Banners        []interface{} `json:"banners"`
	SuggestionList []interface{} `json:"suggestion_list"`
}

type InputDivar struct {
	JSONSchema struct {
		Category struct {
			Value string `json:"value"`
		} `json:"category"`
	} `json:"json_schema"`
	Page int `json:"page"`
}
