package robot

import (
	"encoding/json"
)

type Robot struct {
	uuid string
	name string
}

// func NewRobot(name string) Robot {
// 	return Robot{
// 		uuid: uuid.Must(uuid.NewRandom()).String(),
// 		name: name,
// 	}
// }

func (r Robot) ToJson() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"uuid": r.uuid,
		"name": r.name,
	})
}

func (r *Robot) Id() string {
	return r.uuid
}

func (r *Robot) Rename(name string) {
	r.name = name
}
