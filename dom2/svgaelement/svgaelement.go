package svgaelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGAElement struct
// js:"SVGAElement,omit"
type SVGAElement struct {
	window.SVGGraphicsElement
}

// Href
func (*SVGAElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Href")
	return href
}

// Target
func (*SVGAElement) Target() (target *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Target")
	return target
}
