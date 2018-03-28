/*
In quantum computing, a qubit or quantum bit (sometimes qbit) is a unit of quantum information - the quantum analogue
of the classical binary bit. A qubit is a two-state quantum-mechanical system, such as the polarization
of a single photon: here the two states are vertical polarization and horizontal polarization.
In a classical system, a bit would have to be in one state or the other. However, quantum mechanics allows the qubit
to be in a superposition of both states at the same time, a property that is fundamental to quantum computing.
*/

package quantum

import (
	"fmt"
	"math"
	"math/rand"
)

type Qbit struct {
	state0 float64 //a
	state1 float64 //b

	wasMeasured bool
	stateFix    int
}

func (q *Qbit) Measure() int {
	if q.wasMeasured {
		return q.stateFix
	}

	x := rand.Float64()
	if math.Abs(x-q.state0) > math.Abs(x-q.state1) {
		q.stateFix = 0
	} else {
		q.stateFix = 1
	}
	q.wasMeasured = true

	return q.stateFix
}

func (q *Qbit) String() string {
	q.Measure()

	return fmt.Sprint(q.stateFix)
}

func NewQbit() *Qbit {
	var q = &Qbit{}

	prop := rand.Float64()
	q.state0 = prop
	q.state1 = 1 - prop

	return q
}
