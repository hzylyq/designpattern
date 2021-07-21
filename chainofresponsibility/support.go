package chainofresponsibility

import "log"

type Manager interface {
	resolve(t trouble) bool
}

type support struct {
	Manager
	Next *support
	Name string
}

func (s *support) SetNext(next *support) {
	s.Next = next
}

func (s *support) Done(t trouble) {
	log.Println(t.ToString() + " is resolved by " + s.Name)
}

func (s *support) Fail(t trouble) {
	log.Println(t.ToString() + " cannot be resolved.")
}

func (s *support) Support(t trouble) {
	if s.Manager.resolve(t) {
		s.Done(t)
	} else if s.Next != nil {
		s.Next.Support(t)
	} else {
		s.Fail(t)
	}
}
