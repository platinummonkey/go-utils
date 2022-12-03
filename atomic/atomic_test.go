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

	require.False(testBool.CompareAndSwap(false, false))
	require.True(testBool.CompareAndSwap(false, true))
}