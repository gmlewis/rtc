package yaml

import (
	"log"

	"github.com/gmlewis/rtc/rtc"
)

// AddToWorld adds the yaml data to the RTC world as is (no added groups).
func (y *YAMLFile) AddToWorld(w *rtc.WorldT) {
	for _, item := range y.Items {
		if item.Define != nil {
			continue
		}

		switch *item.Add {
		case "camera":
			continue
		case "light":
			y.addLight(&item, w)
		case "plane":
			y.addPlane(&item, w)
		case "sphere":
			y.addSphere(&item, w)
		case "cube":
			y.addCube(&item, w)
		default:
			log.Printf("unknown YAML item: %v", item)
		}
	}
}

func (y *YAMLFile) addLight(item *Item, w *rtc.WorldT) {
	position := rtc.Point(item.At[0], item.At[1], item.At[2])
	intensity := rtc.Color(item.Intensity[0], item.Intensity[1], item.Intensity[2])
	light := rtc.PointLight(position, intensity)
	w.Lights = append(w.Lights, light)
}

func (y *YAMLFile) addPlane(item *Item, w *rtc.WorldT) {
	object := rtc.Plane()
	y.addMaterial(item, object)
	y.setTransform(item, object)
	w.Objects = append(w.Objects, object)
}

func (y *YAMLFile) addSphere(item *Item, w *rtc.WorldT) {
	object := rtc.Sphere()
	y.addMaterial(item, object)
	y.setTransform(item, object)
	w.Objects = append(w.Objects, object)
}

func (y *YAMLFile) addCube(item *Item, w *rtc.WorldT) {
	object := rtc.Cube()
	y.addMaterial(item, object)
	y.setTransform(item, object)
	w.Objects = append(w.Objects, object)
}

func (y *YAMLFile) addMaterial(item *Item, o rtc.Object) {
	material := y.getMaterial(item)
	o.SetMaterial(material)
}

func (y *YAMLFile) setTransform(item *Item, o rtc.Object) {
	transform := y.getTransform(item)
	o.SetTransform(transform)
}

func (y *YAMLFile) getMaterial(item *Item) rtc.MaterialT {
	material := rtc.Material()
	if item.Extend != nil {
		material = y.getMaterialByName(*item.Extend)
	}
	m := item.Material
	if m == nil {
		return material
	}
	if m.NamedItem != nil {
		return y.getMaterialByName(*m.NamedItem)
	}
	if m.Ambient != nil {
		material.Ambient = *m.Ambient
	}
	if len(m.Color) == 3 {
		material.Color = rtc.Color(m.Color[0], m.Color[1], m.Color[2])
	}
	if m.Diffuse != nil {
		material.Diffuse = *m.Diffuse
	}
	if m.Reflective != nil {
		material.Reflective = *m.Reflective
	}
	if m.RefractiveIndex != nil {
		material.RefractiveIndex = *m.RefractiveIndex
	}
	if m.Shininess != nil {
		material.Shininess = *m.Shininess
	}
	if m.Specular != nil {
		material.Specular = *m.Specular
	}
	if m.Transparency != nil {
		material.Transparency = *m.Transparency
	}
	return material
}

func (y *YAMLFile) getMaterialByName(name string) rtc.MaterialT {
	item, ok := y.DefinedItems[name]
	if !ok {
		log.Printf("Unknown material named %q, ignoring.", name)
		return rtc.Material()
	}
	return y.getMaterial(item)
}

func (y *YAMLFile) getTransform(item *Item) rtc.M4 {
	transform := rtc.M4Identity()
	for _, t := range item.Transform {
		if t.NamedItem != nil {
			xfm := y.getTransformByName(*t.NamedItem)
			transform = xfm.Mult(transform)
			continue
		}

		if t.Type == nil {
			log.Printf("expected transform type, got %#v", *t)
			continue
		}

		var m4 rtc.M4
		switch v := *t.Type; v {
		case "translate":
			m4 = rtc.Translation(t.Args[0], t.Args[1], t.Args[2])
		case "scale":
			m4 = rtc.Scaling(t.Args[0], t.Args[1], t.Args[2])
		case "rotate-x":
			m4 = rtc.RotationX(t.Args[0])
		case "rotate-y":
			m4 = rtc.RotationY(t.Args[0])
		case "rotate-z":
			m4 = rtc.RotationZ(t.Args[0])
		default:
			log.Printf("unhandled transform: %v", v)
		}
		transform = m4.Mult(transform)
	}
	return transform
}

func (y *YAMLFile) getTransformByName(name string) rtc.M4 {
	item, ok := y.DefinedItems[name]
	if !ok {
		log.Printf("Unknown Transform named %q, ignoring.", name)
		return rtc.M4Identity()
	}
	return y.getTransform(item)
}
