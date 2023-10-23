package instances

// LightBulb * Instance 1: Light Bulb
type LightBulb struct {
	Trigger bool
	Degree  int
	Color   string
}

// Curtain * Instance 2: Curtain
type Curtain struct {
	IsHorizontal bool
	IsCenterMode bool
	IsLeftOrTop  bool
	Degree       int
}

// Aircon * Instance 3: Airconditioner
type Aircon struct {
	Trigger     bool
	Temperature int
	WindDegree  int
}

// Refrigerator * Instance 4: Refrigerator
type Refrigerator struct {
	Trigger bool
}

// WaterDispenser * Instance 5: Water Dispenser
type WaterDispenser struct {
	TriggerReminder bool
	TriggerWater    bool
}

// Television * Instance 6: Television
type Television struct {
	Trigger bool
}

// Soundbar * Instance 7: Soundbar
type Soundbar struct {
	Trigger bool
}

// MassageChair * Instance 8: Massage Chair
type MassageChair struct {
	Trigger bool
}

// AiSpeaker * Instance 9: AI Speaker
type AiSpeaker struct {
	Trigger bool
}
