package badge

import (
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasttemplate"
	"io"
	"hits/api/utils"
)

var badgeSVG = utils.Trim(`
<svg
	xmlns="http://www.w3.org/2000/svg"
	xmlns:xlink="http://www.w3.org/1999/xlink" width="{{rectWidth}}" height="20" role="img" aria-label="{{label}}: {{count}} hits">

	<title>{{label}}: {{count}} hits</title>
	<linearGradient id="s" x2="0" y2="100%">
		<stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
		<stop offset="1" stop-opacity=".1"/>
	</linearGradient>
	<clipPath id="r">
		<rect width="100" height="20" rx="{{border}}" fill="#fff"/>
	</clipPath>
	<g clip-path="url(#r)">
		<rect width="38" height="20" fill="{{leftBgColor}}"/>
		<rect x="38" width="62" height="20" fill="{{rightBgColor}}"/>
		<rect width="{{rectWidth}}" height="20" fill="url(#s)"/>
	</g>
	<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" font-size="110">
		<text x="195" y="140" transform="scale(.1)" fill="{{color}}">{{label}}</text>
		<text x="695" y="140" transform="scale(.1)" fill="{{color}}">{{count}}</text>
	</g>
</svg>
`)

func GenerateBadge(label, count, color, leftBgColor, rightBgColor, border string) (badge []byte) {
	const rectWidth = "100"
	var tmpl = fasttemplate.New(badgeSVG, "{{", "}}")
	buf := bytebufferpool.Get()

	if border == "square" {
		border = "0"
	} else {
		border = "3"
	}

	if label == "" {
		label = "hits"
	}

	_, _ = tmpl.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "label":
			return buf.WriteString(label)
		case "count":
			return buf.WriteString(count)
		case "color":
			return buf.WriteString(color)
		case "leftBgColor":
			return buf.WriteString(leftBgColor)
		case "rightBgColor":
			return buf.WriteString(rightBgColor)
		case "rectWidth":
			return buf.WriteString(rectWidth)
		case "border": 
			return buf.WriteString(border)
		}
		return 0, nil
	})
	badge = buf.Bytes()
	bytebufferpool.Put(buf)
	return badge
}
