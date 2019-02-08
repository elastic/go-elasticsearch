// +build !integration

package esapi

import (
	"testing"
)

func TestAPIHelpers(t *testing.T) {
	t.Run("BoolPtr", func(t *testing.T) {
		var v *bool

		v = BoolPtr(false)
		if v == nil || *v != false {
			t.Errorf("Expected false, got: %v", v)
		}

		v = BoolPtr(true)
		if v == nil || *v != true {
			t.Errorf("Expected true, got: %v", v)
		}
	})

	t.Run("IntPtr", func(t *testing.T) {
		var v *int

		v = IntPtr(0)
		if v == nil || *v != 0 {
			t.Errorf("Expected 0, got: %v", v)
		}
	})
}
