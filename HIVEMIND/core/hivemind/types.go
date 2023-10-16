package hivemind

type Status int8

type InstanceResponse struct {
	Ok    Status
	error error
}

type RESTResponse struct {
	Ok Status
}

const (
	SUCCESSFUL Status = iota
	NO_DEVICE
	ERROR
)
