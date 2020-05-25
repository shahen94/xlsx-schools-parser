package computing

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

var (
	// ErrStateNotFound thrown when state not found by the given ID
	ErrStateNotFound = errors.New("State not found in statesMap")
)

// StateItem represents one State
// of USA
type StateItem struct {
	Name      string
	StateCode string
	Timezone  string
}

// StateMapper holds functionality
// Related to States mapping from stateID to state name
// Load from the given file csv
type StateMapper struct {
	Path      string
	statesMap map[int]*StateItem
}

func (states *StateMapper) Read() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := path.Join(dir, states.Path)
	fileDescriptor, err := os.Open(filePath)
	defer fileDescriptor.Close()

	if err != nil {
		return err
	}

	reader := csv.NewReader(fileDescriptor)
	states.statesMap = make(map[int]*StateItem)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stateID, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}

		fmt.Println(record)
		stateItem := &StateItem{
			Name:      record[1],
			StateCode: record[2],
			Timezone:  record[3],
		}
		states.statesMap[stateID] = stateItem
	}
	return nil
}

// GetStateByID returns state name by the given stateID
func (states *StateMapper) GetStateByID(stateID int) (*StateItem, error) {
	stateName, ok := states.statesMap[stateID]

	if !ok {
		return nil, ErrStateNotFound
	}
	return stateName, nil
}
