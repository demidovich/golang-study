package robot

import "github.com/google/uuid"

type Robot struct {
	id string
}

func NewRobot() Robot {
	return Robot{
		id: uuid.Must(uuid.NewRandom()).String(),
	}
}

func (r Robot) Id() string {
	return r.id
}
