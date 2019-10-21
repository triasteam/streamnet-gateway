package lb

import (
	"errors"
	"../core"
)
type LoadBalance interface {
	Target() (*core.Target, error)
}
var ErrNoPointer = errors.New("no endpoints available")