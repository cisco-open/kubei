package models

import "testing"

//nolint:staticcheck
func TestLoader(t *testing.T) {
	tests := []struct {
		Name  string
		Left  *[]Tag
		Right *[]Tag

		Expected map[string]string
	}{
		{
			Name:     "Both nil",
			Left:     nil,
			Right:    nil,
			Expected: nil,
		},
		{
			Name: "Left nil",
			Left: nil,
			Right: &[]Tag{
				{
					Key:   "key1",
					Value: "right",
				},
				{
					Key:   "key2",
					Value: "right",
				},
			},
			Expected: map[string]string{
				"key1": "right",
				"key2": "right",
			},
		},
		{
			Name: "Right nil",
			Left: &[]Tag{
				{
					Key:   "key1",
					Value: "left",
				},
				{
					Key:   "key2",
					Value: "left",
				},
			},
			Right: nil,
			Expected: map[string]string{
				"key1": "left",
				"key2": "left",
			},
		},
		{
			Name: "Merge",
			Left: &[]Tag{
				{
					Key:   "key1",
					Value: "left",
				},
				{
					Key:   "key2",
					Value: "left",
				},
			},
			Right: &[]Tag{
				{
					Key:   "key2",
					Value: "right",
				},
				{
					Key:   "key3",
					Value: "right",
				},
			},
			Expected: map[string]string{
				"key1": "left",
				"key2": "right",
				"key3": "right",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			merged := MergeTags(test.Left, test.Right)

			if merged == nil && test.Expected == nil {
				return
			}

			if merged == nil && test.Expected != nil {
				t.Errorf("unexpected result: left=%v, expected=%v", test.Left, test.Expected)
			}

			m := map[string]string{}
			for _, tag := range *merged {
				m[tag.Key] = tag.Value
			}

			if len(m) != len(test.Expected) {
				t.Errorf("unexpected length of the result: actual=%d, expected=%d", len(m), len(test.Expected))
			}

			for k, v := range test.Expected {
				vv, ok := m[k]
				if !ok {
					t.Errorf("missing key from merged result. key=%s", k)
				}

				if v != vv {
					t.Errorf("mismatch in values for key. key=%s expected=%s actual=%s", k, v, vv)
				}
			}
		})
	}
}
