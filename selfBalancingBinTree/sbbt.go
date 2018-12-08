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
	var tmp *tree.Sucker
	for _, v := range coco {
		adder(warehouse, tmp, v)
		tmp = nil
	}
}

func adder(warehouse []*tree.Sucker, goods *tree.Sucker, value int) {
	// TODO: error handling
	var (
		pos tree.Born
		up  *tree.Sucker
	)
	vCheck := func(current *tree.Sucker) tree.Born {
		switch {// complex logical operation goes here

		}
		return tree.NULL
	}
	warehouse[0].Travel(vCheck)
	goods.InitNode(value, up, pos)
}

func balanceCheck(warehouse []*tree.Sucker) {
	// TODO: error handling
}
