/*
Direct programming operations on Qbits.
*/

package quantum

func Set1Sate(qbit *Qbit) {
	s := qbit.Measure()

	if s != 1 {
		GateX(qbit)
	}
}

func Set0Sate(qbit *Qbit) {
	s := qbit.Measure()

	if s != 0 {
		GateX(qbit)
	}
}

func SetSuperpositionSate(qbit *Qbit) {
	GateH(qbit)
}
