package main

type ButtonConfig struct {
	PaddingX        int     `json:"padding_x"`
	PaddingY        int     `json:"padding_y"`
	Text            string  `json:"text"`
	FontScale       float32 `json:"font_scale"`
	TextColor       string  `json:"text_color"`
	BackgroundColor string  `json:"background_color"`
}
