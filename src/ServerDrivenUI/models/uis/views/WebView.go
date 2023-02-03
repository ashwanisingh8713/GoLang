package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewtype"
)

type webView struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Value    WebViewValue    `json:"value,omitempty"`
}

func WebView(modifier properties.Modifier, value WebViewValue) *webView {
	webView := new(webView)
	webView.Type = viewtype.WEBVIEW.String()
	webView.Modifier = modifier
	webView.Value = value
	return webView
}

type WebViewValue struct {
	URL string `json:"url"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
	BackgroundColor properties.Color  `json:"backgroundColor"`
}