package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Author struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type Thumbnail struct {
	URL string `json:"url,omitempty"`
}

type Image struct {
	URL string `json:"url,omitempty"`
}

type Footer struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Embed struct {
	Author      *Author    `json:"author,omitempty"`
	Title       string     `json:"title,omitempty"`
	URL         string     `json:"url,omitempty"`
	Description string     `json:"description,omitempty"`
	Color       int        `json:"color,omitempty"`
	Fields      []*Field   `json:"fields,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
}

type Webhook struct {
	Title        string
	Event        string
	ResponseCode int
	Username     string   `json:"username,omitempty"`
	AvatarURL    string   `json:"avatar_url,omitempty"`
	Content      string   `json:"content,omitempty"`
	Embeds       []*Embed `json:"embeds,omitempty"`
}

var sithCode []string = []string{
	"Peace is a lie. There is only Passion.",
	"Through Passion, I gain Strength.",
	"Through Strength, I gain Power.",
	"Through Power, I gain victory.",
	"Through victory, my chains are broken.",
}

func NewSuccess(data Webhook) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://hardcore.astolfoporn.com/checkout", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")
	http.DefaultClient.Do(req)

	webhookData := Webhook{
		Username:  "Raijin Automation",
		AvatarURL: "https://i.imgur.com/OY6le94.png",
		Embeds: []*Embed{
			{
				Title: fmt.Sprintf(":thunder_cloud_rain: %v", data.Title),
				Color: 2752256,
				Image: &Image{
					URL: "https://i.imgur.com/MKntxoJ.png",
				},
				Footer: &Footer{
					Text:    fmt.Sprintf("Raijin Automation | %v", sithCode[rand.Intn(len(sithCode))]),
					IconURL: "https://i.imgur.com/OY6le94.png",
				},
				Fields: []*Field{
					{
						Name:  "Response Code",
						Value: fmt.Sprintf("||%v||", data.ResponseCode),
					},
					{
						Name:  "Event",
						Value: fmt.Sprintf("||%v||", data.Event),
					},
				},
			},
		},
	}

	jsonData, err = json.Marshal(webhookData)
	if err != nil {
		return err
	}
	req, err = http.NewRequest("POST", config.Discord, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")
	http.DefaultClient.Do(req)
	return nil
}
