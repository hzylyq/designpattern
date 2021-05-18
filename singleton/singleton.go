package singleton

type Singleton struct {

}

func (s *Singleton) getInstance() *Singleton {
	return s
}
