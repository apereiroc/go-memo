package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel_CreateInitialModel(t *testing.T) {
	m := NewModel()
	assert.True(t, len(m.groups) > 0)

	for _, g := range m.groups {
		assert.True(t, len(g.cmds) > 0)

		for _, c := range g.cmds {
			assert.NotEqual(t, c.cmd, "")
		}
	}

	assert.Equal(t, m.selectedGroup, index(0))
	assert.Equal(t, m.selectedCmd, index(0))
	assert.Equal(t, m.view, groupView)
	assert.Equal(t, m.keys, groupKeys)
	assert.Equal(t, m.quitWithCmd, false)
}

func TestModel_NextPrevInGroupView(t *testing.T) {
	// create a simple model
	m := model{
		view:          groupView,
		keys:          groupKeys,
		selectedGroup: 0,
		selectedCmd:   0,
	}

	// create and append 10 groups
	for range 10 {
		g := group{}
		m.groups = append(m.groups, g)
	}

	// check that calling next() increases the index correctly
	counter := 0
	assert.Equal(t, index(counter), m.selectedGroup)
	for range 9 {
		m.next()
		counter++
		assert.Equal(t, index(counter), m.selectedGroup)
	}

	// should go back to the beginning
	m.next()
	counter = 0
	assert.Equal(t, index(counter), m.selectedGroup)

	// now check prev
	m.prev()
	counter = 9
	assert.Equal(t, index(counter), m.selectedGroup)

	for range 9 {
		m.prev()
		counter--
		assert.Equal(t, index(counter), m.selectedGroup)
	}
}
