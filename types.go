package main

type ButtonConfig struct {
	PaddingX        int    `json:"padding_x"`
	PaddingY        int    `json:"padding_y"`
	FontName        string `json:"font"`
	Text            string `json:"text"`
	FontSize        int    `json:"font_size"`
	TextColor       string `json:"text_color"`
	BackgroundColor string `json:"background_color"`
}
