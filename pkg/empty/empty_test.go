package empty_test

import (
	"testing"

	"github.com/YuukiHayashi0510/todo-app/pkg/empty"
)

func TestIs(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		tests := map[string]struct {
			value string
			want  bool
		}{
			"empty":     {"", true},
			"non-empty": {"hello", false},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				if got := empty.Is(tt.value); got != tt.want {
					t.Errorf("Is(%q) = %v, want %v", tt.value, got, tt.want)
				}
			})
		}
	})

	t.Run("numbers", func(t *testing.T) {
		tests := map[string]struct {
			value interface{}
			want  bool
		}{
			"zero int":         {0, true},
			"non-zero int":     {1, false},
			"zero int8":        {int8(0), true},
			"zero int16":       {int16(0), true},
			"zero int32":       {int32(0), true},
			"zero int64":       {int64(0), true},
			"zero uint":        {uint(0), true},
			"zero uint8":       {uint8(0), true},
			"zero uint16":      {uint16(0), true},
			"zero uint32":      {uint32(0), true},
			"zero uint64":      {uint64(0), true},
			"zero float32":     {float32(0), true},
			"zero float64":     {float64(0), true},
			"non-zero float32": {float32(1.1), false},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				if got := empty.Is(tt.value); got != tt.want {
					t.Errorf("Is(%v) = %v, want %v", tt.value, got, tt.want)
				}
			})
		}
	})

	t.Run("slices", func(t *testing.T) {
		tests := map[string]struct {
			value []int
			want  bool
		}{
			"nil":       {nil, true},
			"empty":     {[]int{}, true},
			"non-empty": {[]int{1}, false},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				if got := empty.Is(tt.value); got != tt.want {
					t.Errorf("Is(%v) = %v, want %v", tt.value, got, tt.want)
				}
			})
		}
	})

	t.Run("maps", func(t *testing.T) {
		tests := map[string]struct {
			value map[string]int
			want  bool
		}{
			"nil":       {nil, true},
			"empty":     {map[string]int{}, true},
			"non-empty": {map[string]int{"a": 1}, false},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				if got := empty.Is(tt.value); got != tt.want {
					t.Errorf("Is(%v) = %v, want %v", tt.value, got, tt.want)
				}
			})
		}
	})

	t.Run("structs", func(t *testing.T) {
		type testStruct struct{}
		ptr := &testStruct{}
		tests := map[string]struct {
			value interface{}
			want  bool
		}{
			"nil struct pointer":     {(*testStruct)(nil), true},
			"non-nil struct pointer": {ptr, false},
			"empty struct":           {testStruct{}, false},
			"empty interface":        {interface{}(nil), true},
			"non-empty interface":    {interface{}("hello"), false},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				if got := empty.Is(tt.value); got != tt.want {
					t.Errorf("Is(%v) = %v, want %v", tt.value, got, tt.want)
				}
			})
		}
	})
}

func TestAny(t *testing.T) {
	tests := map[string]struct {
		values []interface{}
		want   bool
	}{
		"all empty": {
			values: []interface{}{"", 0, []int(nil), map[string]int(nil)},
			want:   true,
		},
		"some empty": {
			values: []interface{}{"hello", 0, []int{1}},
			want:   true,
		},
		"none empty": {
			values: []interface{}{"hello", 1, []int{1}, map[string]int{"a": 1}},
			want:   false,
		},
		"empty values": {
			values: []interface{}{},
			want:   false,
		},
		"nil value": {
			values: nil,
			want:   false,
		},
		"mixed types": {
			values: []interface{}{
				"",               // empty string
				0,                // zero int
				[]int{},          // empty slice
				map[string]int{}, // empty map
				(*struct{})(nil), // nil pointer
				interface{}(nil), // nil interface
				make(chan int),   // empty channel
			},
			want: true,
		},
		"all non-empty mixed types": {
			values: []interface{}{
				"hello",                // non-empty string
				1,                      // non-zero int
				[]int{1},               // non-empty slice
				map[string]int{"a": 1}, // non-empty map
				&struct{}{},            // non-nil pointer
				interface{}("hello"),   // non-nil interface
			},
			want: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := empty.Any(tt.values...); got != tt.want {
				t.Errorf("Any(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}
