package p0155_min_stack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		arguments  [][]int
		expected   []interface{} // use nil for null/void returns
	}{
		{
			name:       "Example 1",
			operations: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			arguments:  [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			expected:   []interface{}{nil, nil, nil, nil, -3, nil, 0, -2},
		},
		{
			name:       "Mixed Operations",
			operations: []string{"MinStack", "push", "push", "getMin", "push", "getMin", "pop", "getMin"},
			arguments:  [][]int{{}, {2}, {1}, {}, {3}, {}, {}, {}},
			expected:   []interface{}{nil, nil, nil, 1, nil, 1, nil, 1},
		},
		{
			name:       "Same Values",
			operations: []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			arguments:  [][]int{{}, {2}, {2}, {2}, {}, {}, {}, {}, {}},
			expected:   []interface{}{nil, nil, nil, nil, 2, nil, 2, nil, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj MinStack
			for i, op := range tt.operations {
				arg := tt.arguments[i]
				exp := tt.expected[i]

				switch op {
				case "MinStack":
					obj = Constructor()
				case "push":
					obj.Push(arg[0])
				case "pop":
					obj.Pop()
				case "top":
					res := obj.Top()
					if res != exp.(int) {
						t.Errorf("step %d: %s(%v) = %v; want %v", i, op, arg, res, exp)
					}
				case "getMin":
					res := obj.GetMin()
					if res != exp.(int) {
						t.Errorf("step %d: %s(%v) = %v; want %v", i, op, arg, res, exp)
					}
				}
			}
		})
	}
}
