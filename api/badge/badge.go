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
	xmlns:xlink="http://www.w3.org/1999/xlink" width="{{rectWidth}}" height="20" role="img" aria-label="{{title}}: {{value}}">
	<title>{{title}}: {{value}}</title>
	<linearGradient id="s" x2="0" y2="100%">
		<stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
		<stop offset="1" stop-opacity=".1"/>
	</linearGradient>
	<clipPath id="r">
		<rect width="100" height="20" rx="3" fill="#fff"/>
	</clipPath>
	<g clip-path="url(#r)">
		<rect width="69" height="20" fill="#555"/>
		<rect x="69" width="31" height="20" fill="{{backgroundColor}}"/>
		<rect width="{{rectWidth}}" height="20" fill="url(#s)"/>
	</g>
	<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">
		<text x="355" y="140" transform="scale(.1)" fill="{{color}}">hits</text>
		<text x="835" y="140" transform="scale(.1)" fill="{{color}}">{{count}}</text>
	</g>
</svg>
`)

func GenerateBadge(count, color, backgroundColor string) (badge []byte) {
	const rectWidth = "100"
	var tmpl = fasttemplate.New(badgeSVG, "{{", "}}")
	buf := bytebufferpool.Get()

	_, _ = tmpl.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "count":
			return buf.WriteString(count)
		case "color":
			return buf.WriteString(color)
		case "backgroundColor":
			return buf.WriteString(backgroundColor)
		}
		return 0, nil
	})
	badge = buf.Bytes()
	bytebufferpool.Put(buf)
	return badge
}
