package model

import (
	"github.com/google/uuid"
)

type Model struct {
	id   string
	name string
}

func New(command CreateCommand) (*Model, error) {
	err := command.validate()
	if err != nil {
		return nil, err
	}

	item := new(Model)
	item.id = uuid.Must(uuid.NewRandom()).String()
	item.name = command.Name

	return item, nil
}

func (item *Model) Id() string {
	return item.id
}

func (item *Model) EqualTo(other *Model) bool {
	return item.id != "" && item.id == other.id
}
