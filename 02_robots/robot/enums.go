package robot

type status string

const (
	StatusActive    status = "active"
	StatusStandBy   status = "stand_by"
	StatusRepairing status = "repairing"
	StatusDisabled  status = "disabled"
)
