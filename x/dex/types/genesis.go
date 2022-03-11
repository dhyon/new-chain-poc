package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LongBookList:  []LongBook{},
		ShortBookList: []ShortBook{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in longBook
	longBookIdMap := make(map[uint64]bool)
	longBookCount := gs.GetLongBookCount()
	for _, elem := range gs.LongBookList {
		if _, ok := longBookIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for longBook")
		}
		if elem.Id >= longBookCount {
			return fmt.Errorf("longBook id should be lower or equal than the last id")
		}
		longBookIdMap[elem.Id] = true
	}
	// Check for duplicated ID in shortBook
	shortBookIdMap := make(map[uint64]bool)
	shortBookCount := gs.GetShortBookCount()
	for _, elem := range gs.ShortBookList {
		if _, ok := shortBookIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for shortBook")
		}
		if elem.Id >= shortBookCount {
			return fmt.Errorf("shortBook id should be lower or equal than the last id")
		}
		shortBookIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
