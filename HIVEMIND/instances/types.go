package instances

// LightBulb * Instance 1: Light Bulb
type LightBulb struct {
	Trigger bool
	Degree  int
	Color   string
}

type Curtain struct {
	IsHorizontal bool
	IsCenterMode bool
	IsLeftOrTop  bool
	Degree       int
}
