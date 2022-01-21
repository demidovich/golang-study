package robot

import "github.com/google/uuid"

type RobotService struct {
	repository RobotRepository
}

func NewRobotService(repository RobotRepository) *RobotService {
	return &RobotService{
		repository: repository,
	}
}

func (s *RobotService) NewRobot(name string) (Robot, error) {
	item := Robot{
		uuid: uuid.Must(uuid.NewRandom()).String(),
		name: name,
	}
	return item, s.repository.Add(item)
}

func (s *RobotService) Rename(item Robot, name string) error {
	item.name = name
	return s.repository.Add(item)
}
