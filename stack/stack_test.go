package stack_test

import (
	"github.com/shults/ds/stack"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack(t *testing.T) {
	st := stack.New[int]()
	require.True(t, st.Empty())
	require.Equal(t, 0, st.Size())
	st.Push(1)
	require.False(t, st.Empty())
	require.Equal(t, 1, st.Size())
	st.Push(2)
	require.False(t, st.Empty())
	require.Equal(t, 2, st.Size())
	res, err := st.Pop()
	require.Nil(t, err)
	require.Equal(t, 2, res)
	require.Equal(t, 1, st.Size())
	require.Equal(t, false, st.Empty())
	res, err = st.Pop()
	require.Nil(t, err)
	require.Equal(t, 1, res)
	require.Equal(t, 0, st.Size())
	require.Equal(t, true, st.Empty())
}
