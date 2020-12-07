package raytracing

import ()

type Material struct {
	RGB                                                                              RGB
	Ambient, Diffuse, Specular, Shininess, Reflective, Transparency, RefractiveIndex float64
	Pattern                                                                          Pattern
}

func GetMaterial() Material {
	return Material{RGB: RGB{Red: 1, Green: 1, Blue: 1},
		Ambient:         0.1,
		Diffuse:         0.9,
		Specular:        0.9,
		Shininess:       200.0,
		RefractiveIndex: 1.0}
}
