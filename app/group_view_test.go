package app

import (
	"testing"

	"github.com/apereiroc/go-memo/models"
	"github.com/stretchr/testify/assert"
)

func TestGroupView_NextPrevInGroupView(t *testing.T) {
	// create a simple model
	m := model{
		view:          groupView{},
		keys:          groupKeys,
		selectedGroup: 0,
		selectedCmd:   0,
	}

	// create and append 10 groups
	for range 10 {
		g := models.Group{}
		m.groups = append(m.groups, g)
	}

	// check that calling next() increases the index correctly
	counter := 0
	assert.Equal(t, index(counter), m.selectedGroup)
	for range 9 {
		m = m.view.next(m)
		counter++
		assert.Equal(t, index(counter), m.selectedGroup)
	}

	// should go back to the beginning
	m = m.view.next(m)
	counter = 0
	assert.Equal(t, index(counter), m.selectedGroup)

	// now check prev
	m = m.view.prev(m)
	counter = 9
	assert.Equal(t, index(counter), m.selectedGroup)

	for range 9 {
		m = m.view.prev(m)
		counter--
		assert.Equal(t, index(counter), m.selectedGroup)
	}
}
