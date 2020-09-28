package controller

import (
	"github.com/sachinmaharana/appsoperator/pkg/controller/appsoperator"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, appsoperator.Add)
}
