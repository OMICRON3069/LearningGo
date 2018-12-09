package main

import (
	"fmt"
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
			warehouse = append(warehouse, tmp)
			initWarehouse(warehouse, v, tmp)
			tmp = nil
			continue
		}
		tmp = new(tree.Sucker)
		// I don'y know why the cap() of slice will not follow my order to extend
		// so the ugly workaround is append tmp first then change it in function call
		warehouse = append(warehouse, tmp)
		adder(warehouse, tmp, v)
		tmp = nil
	}
	// break point here and check the slice
	fmt.Println(warehouse)
}

// an ugly workaround for so call slice are reference type of array
func initWarehouse(warehouse []*tree.Sucker, value int, goods *tree.Sucker) {
	// due to limit of init method
	// must give the value to variable manually
	goods.Parent, goods.LChild, goods.RChild, goods.Value = nil, nil, nil, value

	warehouse[0] = goods

	// TODO: error handling
}

func adder(warehouse []*tree.Sucker, goods *tree.Sucker, value int) {
	// TODO: error handling

	// I must be a dumb ass
	// init pos with NULL to avoid further error
	pos := tree.NULL
	var up *tree.Sucker

	// function start
	vCheck := func(current *tree.Sucker) tree.Born {
		switch { // complex logical operation goes here
		case value < current.Value:
			if current.LChild == nil {
				up = current
				pos = tree.LEFT
				return tree.NULL
			} else {
				return tree.LeftOnly
			}
		case value > current.Value:
			if current.RChild == nil {
				up = current
				pos = tree.RIGHT
				return tree.NULL
			} else {
				return tree.RightOnly
			}
		}
		return tree.NULL
	}

	warehouse[0].Travel(vCheck)
	goods.InitNode(value, up, pos)

	// add good to warehouse,
	//warehouse[cap(warehouse)-1] = goods
}

/*
func balanceCheck(warehouse []*tree.Sucker) {
	// TODO: error handling
}
*/
