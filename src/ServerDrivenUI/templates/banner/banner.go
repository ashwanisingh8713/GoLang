package banner

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/imgscale"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/views"
)

func Banner(listChildren []interface{}) []interface{} {
	var imgViewValue = views.ImageValue{
		ImageUrl: "https://images.all-free-download.com/images/graphicwebp/rosa_red_beautiful_girl_219903.webp",
		Width:    500, Height: 240, ImageScale: imgscale.ImgScale.String(imgscale.FillBounds),
		BackgroundColor: properties.ColorDummy(),
		BorderColor:     properties.BorderColorDummy(),
	}
	var imgView = views.ImageView(imgViewValue)
	listChildren = append(listChildren, imgView)
	return listChildren
}
