package main

import (
	"runtime/debug"

	robot "github.com/demidovich/golang-study/03_error_handling/internal/robot"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	myHandler()
}

func myHandler() {
	robot := robot.NewRobot()
	err := insert(robot)
	if err != nil {
		logError("Unable to handle request for robot "+robot.Id(), err)
		// logTraceableError("Unable to handle request for robot "+robot.Id(), err)
		// logrus.WithError(err).Errorf("Unable to handle request for robot %s", robot.Id())
	}
}

func insert(robot robot.Robot) error {
	err := dbQuery(robot)
	if err != nil {
		return errors.Wrapf(err, "Unable to insert robot %s", robot.Id())
	}
	return nil
}

func dbQuery(robot robot.Robot) error {
	return errors.New("Unable commit transaction")
}

func logError(message string, err error) {
	logrus.WithError(err).Error(message)
}

func logTraceableError(message string, err error) {
	logrus.WithError(err).Error(message, string(debug.Stack()))
}
