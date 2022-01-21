package repository

import (
	"fmt"
	"golang-study/05_repository_interface/internal/domain/robot"
	"golang-study/pkg/util"
	"os"
)

type RobotFileRepository struct {
	storageDir string
}

func NewRobotFileRepository(storageDir string) RobotFileRepository {
	if !util.IsDir(storageDir) {
		panic("path '" + storageDir + "' is not a directory")
	}

	return RobotFileRepository{
		storageDir: storageDir,
	}
}

func (repo *RobotFileRepository) Add(robot robot.Robot) error {
	file := repo.entityFile(robot)
	data, err := robot.ToJson()
	if err != nil {
		return err
	}
	if err := os.WriteFile(file, data, 0666); err != nil {
		return err
	} else {
		return nil
	}
}

func (repo *RobotFileRepository) Remove(robot robot.Robot) error {
	file := repo.entityFile(robot)
	if err := os.Remove(file); err != nil {
		return err
	} else {
		return nil
	}
}

func (repo *RobotFileRepository) entityFile(robot robot.Robot) string {
	return fmt.Sprintf("%s/%s.json", repo.storageDir, robot.Id())
}
