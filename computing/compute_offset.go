package computing

import (
	"actogram.net/go-xlsx-parser/plt4m"
)

// ComputeOffset calculates offset for the given item
func ComputeOffset(states *StateMapper) plt4m.IterateCallback {
	return func(index int, item *plt4m.SchoolItem) {
		stateID := item.GetStateID()
		state, err := states.GetStateByID(stateID)

		if err != nil {
			panic(err)
		}
		item.AddOffset(state.Timezone)
		if len(state.Timezone) > 2 {
			item.Mark()
		}
	}
}
