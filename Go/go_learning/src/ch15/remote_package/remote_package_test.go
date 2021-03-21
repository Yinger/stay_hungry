package remote

// go get -u github.com/easierway/concurrent_map
import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

/*
=== RUN   TestConcurrentMap
    remote_package_test.go:17: 10 true
*/
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}