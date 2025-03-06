package dlist_test

import (
	"github.com/shults/ds/dlist"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList_Append(t *testing.T) {
	lst := dlist.New[int]()
	require.True(t, lst.Empty())
	lst.Append(1)
	lst.Append(2)
	lst.Append(3)
	require.False(t, lst.Empty())

	var items []int

	for !lst.Empty() {
		item, err := lst.Head()
		require.Nil(t, err)
		items = append(items, item)
	}

	require.Equal(t, []int{1, 2, 3}, items)
	require.True(t, lst.Empty())
}

func TestList_Prepend(t *testing.T) {
	lst := dlist.New[int]()
	require.True(t, lst.Empty())
	lst.Prepend(1)
	lst.Prepend(2)
	lst.Prepend(3)
	require.False(t, lst.Empty())

	var items []int

	for !lst.Empty() {
		item, err := lst.Head()
		require.Nil(t, err)
		items = append(items, item)
	}

	require.Equal(t, []int{3, 2, 1}, items)
	require.True(t, lst.Empty())
}

func TestList_Common(t *testing.T) {
	lst := dlist.New[int]()
	require.True(t, lst.Empty())
	lst.Append(1)
	require.Equal(t, 1, lst.Size())
	lst.Prepend(2)
	require.Equal(t, 2, lst.Size())

	tail, err := lst.Tail()
	require.Nil(t, err)
	require.Equal(t, 1, tail)
	require.Equal(t, 1, lst.Size())

	head, err := lst.Head()
	require.Nil(t, err)
	require.Equal(t, 2, head)
	require.Equal(t, 0, lst.Size())
}
