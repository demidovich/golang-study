package robot

import (
	"golang-study/02_robots/model"

	"github.com/google/uuid"
)

type Robot struct {
	id       string
	status   status
	name     string
	model_id string
}

func New(command CreateCommand, model *model.Model) (*Robot, error) {
	err := command.validate()
	if err != nil {
		return nil, err
	}

	item := new(Robot)
	item.id = uuid.Must(uuid.NewRandom()).String()
	item.status = StatusActive
	item.name = command.Name
	item.model_id = model.Id()

	return item, nil
}

func (item *Robot) Id() string {
	return item.id
}

func (item *Robot) EqualTo(other *Robot) bool {
	return item.id != "" && item.id == other.id
}

func (item *Robot) Activate() {
	item.status = StatusActive
}

func (item *Robot) Repair() {
	item.status = StatusRepairing
}
