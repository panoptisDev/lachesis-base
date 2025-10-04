package multidb

import "github.com/panoptisDev/lachesis-base/kvdb"

type closableTable struct {
	kvdb.Store
	underlying kvdb.Store
	noDrop     bool
}

// Close leaves underlying database.
func (s *closableTable) Close() error {
	return s.underlying.Close()
}

// Drop whole database.
func (s *closableTable) Drop() {
	if s.noDrop {
		return
	}
	s.underlying.Drop()
}
