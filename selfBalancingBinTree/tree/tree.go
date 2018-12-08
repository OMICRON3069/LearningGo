package tree

// define a simple tree struct
type Sucker struct {
	value                  int
	parent, lChild, rChild *Sucker
}

// this thing are kind like enum in C/C++
type born int

const (
	LEFT  born = 0
	RIGHT born = 1
)

// plug this node into it's parent's arm
func (s *Sucker) InitNode(num int, up *Sucker, local born) {
	s.value = num
	s.parent = up
	// as mentioned in golang tour
	// use switch to replace else if
	switch {
	case local == LEFT:
		up.lChild = s
	case local == RIGHT:
		up.rChild = s
		// TODO : finish error message after learned error handling
	}
	// set it to empty
	s.rChild, s.lChild = nil, nil
}

// use recursive method Travel all node
// hand function can
func (s *Sucker) Travel(hand func(*Sucker) born) {
	chick := hand(s)
	switch {
	case chick == LEFT:
		// left, then right
		if s.lChild.isNotNil() {
			s.lChild.Travel(hand)
		}
		if s.rChild.isNotNil() {
			s.rChild.Travel(hand)
		}
	case chick == RIGHT:
		if s.rChild.isNotNil() {
			s.rChild.Travel(hand)
		}
		if s.lChild.isNotNil() {
			s.lChild.Travel(hand)
		}
		// TODO: I need an error handler to fill "default", might same as above
	}
}

// this method can help me to stop traveling while reach the end
func (s *Sucker) isNotNil() bool {
	return s != nil
}
