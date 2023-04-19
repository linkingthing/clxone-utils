package traceroute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// run this test with root
func TestTraceroute(t *testing.T) {
	destAddr := "8.8.8.8"
	results, err := Traceroute(destAddr, &TraceConfig{
		FirstTTL: 1,
		MaxTTL:   30,
		Debug:    true,
		WaitSec:  1,
	})
	assert.NoError(t, err)
	t.Logf("result:%+v", results)
}
