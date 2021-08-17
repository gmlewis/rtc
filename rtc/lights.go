package rtc

import "math"

// PointLightT represents a point light.
type PointLightT struct {
	position  Tuple
	intensity Tuple
}

// PointLight returns a point light at the given position (a point Tuple) with
// the provided intensity (a color Tuple).
func PointLight(position Tuple, intensity Tuple) *PointLightT {
	return &PointLightT{position: position, intensity: intensity}
}

// Lighting calculates the lighting on an object and returns the color as a Tuple.
func Lighting(material *MaterialT, object Object, light *PointLightT, point Tuple, eyeVector Tuple, normalVector Tuple, inShadow bool) Tuple {
	color := material.Color
	if material.Pattern != nil {
		color = PatternAt(material.Pattern, object, point)
	}

	effectiveColor := color.HadamardProduct(light.intensity)

	lightV := light.position.Sub(point).Normalize()

	ambient := effectiveColor.MultScalar(material.Ambient)

	if inShadow {
		return ambient
	}

	lightDotNormal := lightV.Dot(normalVector)

	if lightDotNormal < 0 {
		return ambient
	}

	diffuse := effectiveColor.MultScalar(material.Diffuse * lightDotNormal)

	reflectV := lightV.Negate().Reflect(normalVector)
	reflectDotEye := reflectV.Dot(eyeVector)

	specular := Color(0, 0, 0)
	if reflectDotEye > 0 {
		factor := math.Pow(reflectDotEye, material.Shininess)
		specular = light.intensity.MultScalar(material.Specular * factor)
	}

	return ambient.Add(diffuse).Add(specular)
}
