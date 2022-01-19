package main

import (
	"fmt"
	"golang-study/04_error_handling_stack_trace/internal/robot"
	"golang-study/04_error_handling_stack_trace/pkg/app_errors"
)

func main() {
	myHandler()
}

func myHandler() {
	robot := robot.NewRobot()
	err := insert(robot)
	if err != nil {
		logError(err)
	}
}

func insert(r robot.Robot) error {
	err := dbQuery(r)
	if err != nil {
		return err
	}
	return nil
}

func dbQuery(r robot.Robot) error {
	return app_errors.NewAppError("Unable commit transaction for robot " + r.Id())
}

func logError(err error) {
	fmt.Println(err)
}
