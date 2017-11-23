package svganimatednumberlist

import (
	"github.com/matthewmueller/golly/dom2/svgnumberlist"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedNumberList struct
// js:"SVGAnimatedNumberList,omit"
type SVGAnimatedNumberList struct {
}

// AnimVal
func (*SVGAnimatedNumberList) AnimVal() (animVal *svgnumberlist.SVGNumberList) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedNumberList) BaseVal() (baseVal *svgnumberlist.SVGNumberList) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}
