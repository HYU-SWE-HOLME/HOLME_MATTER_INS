package hivemind

type Status int8

type InstanceResponse struct {
	Ok    Status
	error error
}

type RESTResponse struct {
	Result []InstanceResponse
}

type InstanceFrame struct {
	User    string            `json:"user"`
	Request []InstanceRequest `json:"request"`
}

type InstanceRequest struct {
	InstType int    `json:"instance"`
	Payload  string `json:"payload"`
}

type HandleFrameFunction func(string) InstanceResponse

const (
	SUCCESSFUL Status = iota
	NO_DEVICE
	ERROR
)
