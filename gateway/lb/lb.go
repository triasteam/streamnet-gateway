package lb

import (
	"errors"
	"github.com/triasteam/streamnet-gateway/gateway/core"
)
type LoadBalance interface {
	Target() (*core.Target, error)
}
var ErrNoPointer = errors.New("no endpoints available")