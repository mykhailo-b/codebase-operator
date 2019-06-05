package controller

import (
	"codebase-operator/pkg/controller/codebase"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, codebase.Add)
}