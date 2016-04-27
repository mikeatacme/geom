/*
Package geom holds geometry objects and functions to operate on them.
They can be encoded and decoded by other packages in this repository.
*/
package geom

import "github.com/ctessum/geom/proj"

// Geom is an interface for generic geometry types.
type Geom interface {
	Bounds() *Bounds
	Within(Polygonal) bool
	Similar(Geom, float64) bool
	Transform(proj.Transformer) (Geom, error)
}

// Linear is an interface for types that are linear in nature.
type Linear interface {
	Geom
	Length() float64
	//Clip(Polygonal) Linear
	//Intersection(Linear) MultiPoint
	Simplify(tolerance float64) Geom
}

// Polygonal is an interface for types that are polygonal in nature.
type Polygonal interface {
	Geom
	Polygons() []Polygon
	Intersection(Polygonal) Polygon
	Union(Polygonal) Polygon
	XOr(Polygonal) Polygon
	Difference(Polygonal) Polygon
	Area() float64
	Simplify(tolerance float64) Geom
	FixOrientation()
	Centroid() Point
}

// PointLike is an interface for types that are pointlike in nature.
type PointLike interface {
	Geom
	Points() []Point
	On(l Linear, tolerance float64) bool
}
