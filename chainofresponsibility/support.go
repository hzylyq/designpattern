package chainofresponsibility

type Re

type support struct {

	Next *support
}

func (s *support) SetNext(next *support) {
	s.Next = next
}

func (s *support) Support(t trouble) {


	if s.Next != nil {
		s.Next.Support(t)
	}
}
