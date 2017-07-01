package helper

import (
	"github.com/EdlinOrg/prominentcolor"
	"image"
)

func ExtractColoer(orgimg image.Image) string {

	res := prominentcolor.Kmeans(orgimg)
	if len(res) > 0 {
		return "#" + res[0].AsString()
	}

	return "#EDF3EB"
}
