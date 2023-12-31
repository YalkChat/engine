package server

import (
	"github.com/AleRosmo/engine/models/db"
)

func (s *serverImpl) GetUserByUsername(username string) (*db.User, error) {
	return s.db.GetUserByUsername(username)
}
