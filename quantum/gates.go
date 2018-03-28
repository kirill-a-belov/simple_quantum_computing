/*
In quantum computing and specifically the quantum circuit model of computation,
a quantum logic gate (or simply quantum gate) is a basic quantum circuit operating
on a small number of qubits. They are the building blocks of quantum circuits,
like classical logic gates are for conventional digital circuits.
*/

package quantum

import (
	"math"
)

// Single Qbit gates
func GateX(qbit *Qbit) { //NOT gate
	qbit.state0, qbit.state1 = qbit.state1, qbit.state0

	if qbit.wasMeasured {
		if qbit.stateFix == 1 {
			qbit.stateFix = 0
		} else {
			qbit.stateFix = 1
		}
	}
}

func GateZ(qbit *Qbit) {
	qbit.state1 = -qbit.state1

	if qbit.wasMeasured {
		if qbit.stateFix == 1 {
			qbit.stateFix = 0
		} else {
			qbit.stateFix = 1
		}
	}
}

func GateH(qbit *Qbit) { //Hadamard gate
	sqrt2 := math.Sqrt(2)
	s0 := qbit.state0
	s1 := qbit.state1

	qbit.state0 = (s0 + s1) / sqrt2
	qbit.state1 = (s0 - s1) / sqrt2

	qbit.wasMeasured = false
}

// Multi Qbits gates
func GateCX(mQbit, sQbit *Qbit) { //CNOT gate
	ms := sQbit.Measure()

	if ms == 1 {
		GateX(sQbit)
	}
}

func GateS(fQbit, sQbit *Qbit) { //SWAP gate
	*fQbit, *sQbit = *sQbit, *fQbit
}
