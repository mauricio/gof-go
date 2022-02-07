package gof_go

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type Racer struct {
	Position int
	Name     string
}

func TestSortRacers(t *testing.T) {
	racers := []*Racer{
		{
			Position: 10,
			Name:     "Alonso",
		},
		{
			Position: 2,
			Name:     "Verstappen",
		},
		{
			Position: 1,
			Name:     "Hamilton",
		},
		{
			Position: 12,
			Name:     "Vettel",
		},
	}

	sort.SliceStable(racers, func(i, j int) bool {
		return racers[i].Position < racers[j].Position
	})

	assert.Equal(t, []*Racer{
		{
			Position: 1,
			Name:     "Hamilton",
		},
		{
			Position: 2,
			Name:     "Verstappen",
		},
		{
			Position: 10,
			Name:     "Alonso",
		},
		{
			Position: 12,
			Name:     "Vettel",
		},
	}, racers)
}
