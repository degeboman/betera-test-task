package handlers

import "testing"

func TestHandler_validateApodDate(t *testing.T) {
	data := []struct {
		name     string
		date     string
		expected bool
	}{
		{
			name:     "correct",
			date:     "2023-10-21",
			expected: true,
		},
		{
			name:     "not correct date layout1",
			date:     "2023-13-12",
			expected: false,
		},
		{
			name:     "not correct date layout2",
			date:     "2023",
			expected: false,
		},
		{
			name:     "not correct date layout3",
			date:     "2023-02-0 10:34",
			expected: false,
		},
		{
			name:     "not correct date1",
			date:     "2023-08-32",
			expected: false,
		},
		{
			name:     "not correct date2",
			date:     "2023-02-32",
			expected: false,
		},
		{
			name:     "not correct date2",
			date:     "2023-13-02",
			expected: false,
		},
		{
			name:     "not correct date3",
			date:     "2023-130-02",
			expected: false,
		},
		{
			name:     "future date",
			date:     "2024-07-02",
			expected: false,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			if result := validateApodDate(d.date); result != d.expected {
				t.Errorf("Expected %v, got %v", d.expected, result)
			}
		})
	}
}
