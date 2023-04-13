package kontext

import (
	"context"
	"testing"
)

func TestKontext(t *testing.T) {
	pair := []struct {
		key   string
		value any
	}{
		{
			key:   "number1",
			value: 1,
		},
		{
			key:   "number2",
			value: "2",
		},
	}

	kontext := context.Background()

	t.Run("run for pair key value", func(t *testing.T) {
		for _, each := range pair {
			kontext, err := Store(kontext, each.key, each.value)
			if err != nil {
				t.Error(err)
				return
			}

			value, err := Retreive[any](kontext, each.key)
			if err != nil {
				t.Error(err)
				return
			}

			if each.value != value {
				t.Errorf("failed to assert input: %v and expectation value: %v", each.value, value)
				return
			}
		}
	})
}
