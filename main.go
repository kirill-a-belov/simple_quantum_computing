package main

import (
	"./quantum"
	"fmt"
)

func main() {
	// Do this simulations couple times more
	for i := 0; i < 5; i++ {
		QbitsPrimitives()
	}
}

func QbitsPrimitives() {
	// 1.Get new Qbit in superposition
	q1 := quantum.NewQbit()

	// 2.Measure it - Qbit collapse.
	fmt.Println(q1.Measure())

	// 3.Its state newer changed without our intervention
	fmt.Println(q1)

	// 4.Let's use X-gate (NOT)
	quantum.GateX(q1)
	fmt.Println(q1)

	// 5.Let's back Qbit to superposition by using H-gate (Hadamard)
	quantum.GateH(q1)

	// 6.And measure it again
	fmt.Println(q1)

	// 7.Add another Qbit
	q2 := quantum.NewQbit()

	// 8. Make a Bell state (q1 has fix state - see p.6)
	quantum.GateS(q1, q2)
	bState := []*quantum.Qbit{q1, q2}

	// 9. Let's out from Bell state by measuring second Qbit
	bState[1].Measure()

	// 10. All Qbits in array has states, show them
	fmt.Println(bState)

	// 11. Swap them
	quantum.GateS(bState[0], bState[1])
	fmt.Println(bState)

	// 12. Use PL and set first Qbit
	quantum.Set0Sate(bState[0])
	fmt.Println(bState)
}
