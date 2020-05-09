package raytracing

import ()

type Material struct {
	Color                                                                            Color
	Ambient, Diffuse, Specular, Shininess, Reflective, Transparency, RefractiveIndex float64
	Pattern                                                                          Pattern
}

func GetMaterial() Material {
	return Material{Color: Color{Red: 1, Green: 1, Blue: 1},
		Ambient:         0.1,
		Diffuse:         0.9,
		Specular:        0.9,
		Shininess:       200.0,
		RefractiveIndex: 1.0}
}
