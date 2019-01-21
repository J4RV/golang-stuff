package mem

type abstractMemStore struct {
	lastID   int
	lockChan chan bool
}

func (s *abstractMemStore) nextID() int {
	s.lastID++
	return s.lastID
}

func (s *abstractMemStore) lock() {
	s.lazyLock() <- true
}

func (s *abstractMemStore) release() {
	<-s.lazyLock()
}

func (s *abstractMemStore) lazyLock() chan bool {
	if s.lockChan == nil {
		s.lockChan = make(chan bool, 1)
	}
	return s.lockChan
}
