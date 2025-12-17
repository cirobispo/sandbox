package main

import (
	"fmt"

	turning "github.com/cirobispo/sandbox/internal/turning"
	"github.com/cirobispo/sandbox/internal/turning/countingturn"
	turn "github.com/cirobispo/sandbox/internal/turning/turn"
)

func main() {
	fmt.Println("Hello!")
	t := countingturn.New(turning.STA)
	t.AddBeforeChangeEvent(func(side turning.SideTurn) {
		fmt.Println("Before: ", side)
	})

	t.AddAfterChangeEvent(func(side turning.SideTurn) {
		fmt.Println("After: ", side)
	})

	for i := 0; i < 10; i++ {
		t.Execute()
	}

	pT := &(t.Turn)
	fmt.Printf("count: %d\n", turn.GetData[int](pT, "count"))
}
