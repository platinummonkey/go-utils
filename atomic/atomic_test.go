package atomic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtomic(t *testing.T) {
	testBool := NewAtomic(false)
	require.False(t, testBool.Load())
	
	testBool.Store(true)
	require.True(t, testBool.Load())
	
	require.True(t, testBool.Swap(false))
	require.False(t, testBool.Load())

	require.False(t, testBool.CompareAndSwap(true, false))
	require.True(t, testBool.CompareAndSwap(false, true))
	require.True(t, testBool.Load())
}