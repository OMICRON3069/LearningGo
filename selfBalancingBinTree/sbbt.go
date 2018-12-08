package main

import (
	"selfBalancingBinTree/tree"
)

func main() {
	// the warehouse will play a major role in this program
	var warehouse []*tree.Sucker
	start(warehouse, 3, 2, 1, 4, 5, 6, 7, 10, 9, 8)
}

func start(warehouse []*tree.Sucker, coco ...int) {
	tmp := new(tree.Sucker)

	for _, v := range coco {
		if warehouse == nil {
			//fmt.Println("nil warehouse")
			warehouse = append(warehouse,make([]*tree.Sucker,1)...)
			initWarehouse(warehouse, v, tmp)
			tmp = nil
			continue
		}
		adder(warehouse, tmp, v)
		tmp = nil
	}
}

func initWarehouse(warehouse []*tree.Sucker, value int, goods *tree.Sucker) {
	// due to limit of init method
	// must give the value to variable manually
	goods.Parent, goods.LChild, goods.RChild, goods.Value = nil, nil, nil, value

	warehouse [0]=goods

	// TODO: error handling
}

func adder(warehouse []*tree.Sucker, goods *tree.Sucker, value int) {
	// TODO: error handling
	var (
		pos tree.Born
		up  *tree.Sucker
	)
	vCheck := func(current *tree.Sucker) tree.Born {
		switch { // complex logical operation goes here
		case value < current.Value:
			if current.LChild == nil {
				up = current
				return tree.NULL
			} else {
				return tree.LEFT
			}
		case value > current.Value:
			if current.RChild == nil {
				up = current
				return tree.NULL
			} else {
				return tree.RIGHT
			}
		}
		return tree.NULL
	}
	warehouse[0].Travel(vCheck)
	goods.InitNode(value, up, pos)

	// TODO: add new node to slice
}

/*
func balanceCheck(warehouse []*tree.Sucker) {
	// TODO: error handling
}
*/
