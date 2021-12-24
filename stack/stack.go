type Stack struct {
	store []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.store) == 0
}

func (s *Stack) Size() int {
	return len(s.store)
}

func (s *Stack) Push(item interface{}) {
	s.store = append(s.store, item)
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.store[len(s.store)-1]
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return v
}
