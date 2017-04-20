package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	var user = User {1,"bert",50}

	assert.NotNil(t, user, "User should not be nill")
	assert.True(t,user.Username == "bert")
}

// BenchmarTest run with go test -bench=.
func BenchmarkTest(b *testing.B) {
	//b.StopTimer()
	// do time consuming things
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		// do something
	}
}
