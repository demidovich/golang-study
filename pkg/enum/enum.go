package enum

import "github.com/pkg/errors"

type Enum struct {
	Code string
}

type Enumerable interface {
	String() string
}

func (enum Enum) String() string {
	return enum.Code
}

func NewEnumFromString(str string, allValues *[]Enumerable) (Enumerable, error) {
	for _, value := range *allValues {
		if value.String() == str {
			return value, nil
		}
	}
	return Enum{}, errors.Errorf("Unknown '%s' status.", str)
}
