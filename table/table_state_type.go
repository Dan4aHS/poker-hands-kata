package table

type State int32

const (
	StateEmpty State = iota
	StateFlop
	StateTurn
	StateRiver
	StateEnd
)

func (s *State) Next() State {
	if *s == StateEnd {
		return StateEmpty
	}
	return State(int32(*s) + 1)
}
