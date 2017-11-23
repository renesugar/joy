package path2d

import "github.com/matthewmueller/golly/js"

// New fn
func New(path *Path2d) *Path2d {
	js.Rewrite("Path2D")
	return &Path2d{}
}

// Path2d struct
// js:"Path2d,omit"
type Path2d struct {
}

// Arc
func (*Path2D) Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	js.Rewrite("$<.Arc($1, $2, $3, $4, $5, $6)", x, y, radius, startAngle, endAngle, anticlockwise)
}

// ArcTo
func (*Path2D) ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	js.Rewrite("$<.ArcTo($1, $2, $3, $4, $5)", x1, y1, x2, y2, radius)
}

// BezierCurveTo
func (*Path2D) BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32) {
	js.Rewrite("$<.BezierCurveTo($1, $2, $3, $4, $5, $6)", cp1x, cp1y, cp2x, cp2y, x, y)
}

// ClosePath
func (*Path2D) ClosePath() {
	js.Rewrite("$<.ClosePath()")
}

// Ellipse
func (*Path2D) Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	js.Rewrite("$<.Ellipse($1, $2, $3, $4, $5, $6, $7, $8)", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

// LineTo
func (*Path2D) LineTo(x float32, y float32) {
	js.Rewrite("$<.LineTo($1, $2)", x, y)
}

// MoveTo
func (*Path2D) MoveTo(x float32, y float32) {
	js.Rewrite("$<.MoveTo($1, $2)", x, y)
}

// QuadraticCurveTo
func (*Path2D) QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32) {
	js.Rewrite("$<.QuadraticCurveTo($1, $2, $3, $4)", cpx, cpy, x, y)
}

// Rect
func (*Path2D) Rect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.Rect($1, $2, $3, $4)", x, y, w, h)
}
