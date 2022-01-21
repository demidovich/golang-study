package main

import (
	"golang-study/05_repository_interface/internal/domain/robot"
	"golang-study/05_repository_interface/internal/repository"
)

func main() {
	repository := repository.NewRobotFileRepository("./storage")
	robotService := robot.NewRobotService(&repository)

	robot, _ := robotService.NewRobot("Иван")

	robotService.Rename(robot, "Василий")
}
