package robot

import "github.com/go-playground/validator/v10"

type CreateCommand struct {
	Name string `validate:"required,min=2,max=255"`
}

func (c *CreateCommand) validate() error {
	return validator.New().Struct(c)
}
