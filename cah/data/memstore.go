package data

type abstractMemStore struct {
	lastID int
}

func (s *abstractMemStore) nextID() int {
	s.lastID++
	return s.lastID
}
