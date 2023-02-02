package banner

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/imgscale"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/views"
)

func Banner(listChildren []interface{}) []interface{} {
	//var imgViewValue = properties.Value{ImageUrl: "https://cdn.pixabay.com/photo/2015/10/30/18/04/banner-1014539_1280.jpg",
	var imgViewValue = properties.Value{
		ImageUrl: "https://images.all-free-download.com/images/graphicwebp/rosa_red_beautiful_girl_219903.webp",
		Width:    500, Height: 240, ImageScale: imgscale.ImgScale.String(imgscale.FillBounds)}
	var imgViewModifier = properties.Modifier{BackgroundColor: properties.ColorDummy(), BorderColor: properties.BorderColorDummy()}
	var imgView = views.ImageView(imgViewModifier, imgViewValue)
	listChildren = append(listChildren, imgView)
	return listChildren
}
