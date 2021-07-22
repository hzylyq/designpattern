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

func (s *support) SetNext(next *support) *support {
	s.Next = next

	return next
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

type noSupportManager struct{}

func (n *noSupportManager) resolve(t trouble) bool {
	return false
}

func NewNoSupport(name string) *support {
	return &support{
		Manager: &noSupportManager{},
		Name:    name,
	}
}

type limitSupport struct {
	limit int
}

func (l *limitSupport) resolve(t trouble) bool {
	if t.Number() < l.limit {
		return true
	}
	return false
}

func NewLimitSupport(name string, limit int) *support {
	return &support{
		Manager: &limitSupport{limit: limit},
		Name:    name,
	}
}

type oddSupport struct{}

func (o *oddSupport) resolve(t trouble) bool {
	if t.Number()%2 == 1 {
		return true
	}
	return false
}

func NewOddSupport(name string) *support {
	return &support{
		Manager: &oddSupport{},
		Name:    name,
	}
}

type specialSupport struct {
	Number int
}

func (s *specialSupport) resolve(t trouble) bool {
	if t.Number() == s.Number {
		return true
	}
	return false
}

func NewSpecialSupport(name string, number int) *support {
	return &support{
		Name: name,
		Manager: &specialSupport{
			Number: number,
		},
	}
}
