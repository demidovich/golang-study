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
	error := s.repository.Add(item)
	return item, error
}

func (s *RobotService) Rename(item Robot, name string) error {
	item.name = name
	return s.repository.Add(item)
}
