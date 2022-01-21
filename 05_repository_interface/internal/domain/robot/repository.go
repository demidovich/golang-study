package robot

type RobotRepository interface {
	Add(robot Robot) error
	// Retrieve(id string) (Robot, error)
	Remove(robot Robot) error
}
