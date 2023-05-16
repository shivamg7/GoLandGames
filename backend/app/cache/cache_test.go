package cache

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestNewCache(t *testing.T) {
	lc := NewCache()
	_, ok := lc.areaCodeGameModeCounter[12]

	// Assert that ok is false
	assert.Assert(t, !ok)

	lc.mu.Lock()
	lc.mu.Unlock()
}

func TestLocalCache_GetCounter_NoEntry(t *testing.T) {
	lc := NewCache()

	resp, err := lc.GetCounter(123)
	assert.Assert(t, err != nil)
	assert.Assert(t, resp == nil)
}

func TestLocalCache_GetCounter_FoundEntry(t *testing.T) {
	lc := NewCache()

	lc.UpdateCounter(123, "mode")

	resp, err := lc.GetCounter(123)
	assert.Assert(t, err == nil)
	assert.Assert(t, resp != nil)
}
