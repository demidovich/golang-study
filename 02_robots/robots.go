package main

import (
	"fmt"
	"log"

	"golang-study/02_robots/model"
	"golang-study/02_robots/robot"
)

func main() {
	robot1 := makeRobot("R2", "Василий")
	robot2 := makeRobot("D2", "Николай")

	robot1.Repair()

	fmt.Println("Robot1: ", robot1)
	fmt.Println("Robot2: ", robot2)
	fmt.Println("Robot2 equal to Robot1: ", robot2.EqualTo(robot1))

	robot3 := robot1

	fmt.Println("Robot3 equal to Robot1: ", robot3.EqualTo(robot1))
}

func makeRobot(modelName string, robotName string) *robot.Robot {
	model, err := model.New(
		model.CreateCommand{
			Name: "R2",
		},
	)

	if err != nil {
		log.Fatalln(err.Error())
	}

	robot, err := robot.New(
		robot.CreateCommand{
			Name: "Василий",
		},
		model,
	)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return robot
}
