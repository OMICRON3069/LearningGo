package tree

// define a simple tree struct
type sucker struct {
	value                  int
	parent, lChild, rChild *sucker
}

// this thing are kind like enum in C/C++
type Born int

const (
	LEFT  Born = 0
	RIGHT Born = 1
)

// plug this node into it's parent's arm
func (s *sucker) initNode(num int, up *sucker, local Born) {
	s.value = num
	s.parent = up
	// as mentioned in golang tutor
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

// use recursive method travel all node
// hand function can
func (s *sucker) travel(hand func(*sucker) Born) {
	chick := hand(s)
	switch {
	case chick == LEFT:
		// left, then right
		if s.lChild.isNotNil() {
			s.lChild.travel(hand)
		}
		if s.rChild.isNotNil() {
			s.rChild.travel(hand)
		}
	case chick == RIGHT:
		if s.rChild.isNotNil() {
			s.rChild.travel(hand)
		}
		if s.lChild.isNotNil() {
			s.lChild.travel(hand)
		}
		// TODO: I need an error handler to fill "default", might same as above
	}
}

// this method can help me to stop traveling while reach the end
func (s *sucker) isNotNil() bool {
	return s != nil
}
