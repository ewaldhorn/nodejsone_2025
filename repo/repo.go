package repo

import "errors"

// ----------------------------------------------------------------------------
type DataRepo struct {
	People []Person
}

// ----------------------------------------------------------------------------
func NewDataRepository() *DataRepo {
	return &DataRepo{
		People: make([]Person, 0),
	}
}

// ----------------------------------------------------------------------------
func (dr *DataRepo) PeopleCount() int {
	return len(dr.People)
}

// ----------------------------------------------------------------------------
func (dr *DataRepo) GetPersonAt(position int) (*Person, error) {
	if position < 0 || position >= dr.PeopleCount() {
		return nil, errors.New("invalid position specified")
	}

	return nil, errors.New("no person at that position")
}
