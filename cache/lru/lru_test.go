package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Stringer struct {
	str string
}

func NewStringer(str string) *Stringer {
	return &Stringer{
		str: str,
	}
}

func (s *Stringer) Len() int {
	return len(s.str)
}

type TestKV struct {
	key   string
	value string
}

func TestLruCache(t *testing.T) {
	t.Run("test multiple set", func(t *testing.T) {
		cache := NewLruCache(0)
		testPara := []*TestKV{
			{
				key:   "foo",
				value: "bar",
			}, {
				key:   "foo2",
				value: "bar2",
			},
		}
		for _, kv := range testPara {
			cache.Set(kv.key, kv.value)
		}
		for _, kv := range testPara {
			value := cache.Get(kv.key)
			assert.Equal(t, value, kv.value)
			t.Logf("key:%s exp:%s real:%s", kv.key, kv.value, value)
		}
	})
}
