package rtc

// Group creates a group of objects at the origin.
// It implements the Object interface.
func Group(shapes ...Object) *GroupT {
	g := &GroupT{Shape: Shape{transform: M4Identity(), material: Material()}}
	g.AddChild(shapes...)
	return g
}

// AddChild adds shape(s) to a group.
func (g *GroupT) AddChild(shapes ...Object) {
	g.Children = append(g.Children, shapes...)
	for _, shape := range shapes {
		shape.SetParent(g)
	}
}

// GroupT represents a group of objects with its own transformation matrix.
// It implements the Object interface.
type GroupT struct {
	Shape
	Children []Object
}

var _ Object = &GroupT{}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (s *GroupT) LocalIntersect(ray RayT) []IntersectionT {
	return nil
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (s *GroupT) LocalNormalAt(objectPoint Tuple) Tuple {
	return Tuple{}
}
