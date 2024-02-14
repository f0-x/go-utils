package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	fmt.Printf("The edge value of tree %v is %v\n",t, t.Value);
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func () {
		Walk(t1, channel1)
		close(channel1)
	}() 
	
	go func () {
		Walk(t2, channel2)
		close(channel2)
	}() 
	
	

	for i := 0; i <= 10; i++ {
		fmt.Println("The count of loop", i+1);
		if <-channel1 != <-channel2 {
			fmt.Println("Value not equal", <-channel1, <-channel2)
			return false
		}
	}
	return true
}

func main() {
	tree1 := tree.New(3)
	tree2 := tree.New(5)
	fmt.Println("The equality of two trees", Same(tree1, tree2))
}
