package robot

type Robot struct {
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	ModelId string          `json:"name"`
	Status  RobotStatusType `json:"status"`
}
