package stele

import (
	"github.com/dollarkillerx/stele/internal/server"
)

func NewLocal(storagePath string) (*server.Local, error) {
	return server.NewLocal(storagePath)
}
