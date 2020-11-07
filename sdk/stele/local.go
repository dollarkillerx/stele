package stele

import (
	"github.com/dollarkillerx/stele/pkg/stele"
)

func NewLocal(storagePath string) (*stele.Local, error) {
	return stele.NewLocal(storagePath)
}
