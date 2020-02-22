package controller

import (
	"github.com/grdryn/dummy-owner-test-operator/pkg/controller/clusterscopedowner"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, clusterscopedowner.Add)
}
