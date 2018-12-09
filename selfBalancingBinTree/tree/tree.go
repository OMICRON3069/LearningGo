package tree

// define a simple tree struct
type Sucker struct {
	Value                  int
	Parent, LChild, RChild *Sucker
}

/*
	I'm considering to add a variable to log the element's position in array/slice
	which can help me to delete/release a node in tree
 */

// this thing are kind like enum in C/C++
type Born int

const (
	LEFT  Born = 0
	RIGHT Born = 1
	NULL  Born = 9
)

// plug this node into it's Parent's arm
func (s *Sucker) InitNode(num int, up *Sucker, local Born) {
	s.Value = num
	s.Parent = up
	// as mentioned in golang tour
	// use switch to replace else if
	switch {
	case local == LEFT:
		up.LChild = s
	case local == RIGHT:
		up.RChild = s
		// TODO : finish error message after learned error handling
	}
	// set it to empty
	s.RChild, s.LChild = nil, nil
}

// use recursive method Travel all node
// hand function can
func (s *Sucker) Travel(hand func(*Sucker) Born) {
	chick := hand(s)
	switch {
	case chick == LEFT:
		// left, then right
		if s.LChild.isNotNil() {
			s.LChild.Travel(hand)
		}
		if s.RChild.isNotNil() {
			s.RChild.Travel(hand)
		}
	case chick == RIGHT:
		if s.RChild.isNotNil() {
			s.RChild.Travel(hand)
		}
		if s.LChild.isNotNil() {
			s.LChild.Travel(hand)
		}
	case chick == NULL:
		// do nothing, continue recursive travel

		// TODO: I need an error handler to fill "default", might same as above
	}
}

// this method can help me to stop traveling while reach the end
func (s *Sucker) isNotNil() bool {
	return s != nil
}
