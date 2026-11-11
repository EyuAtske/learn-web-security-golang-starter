package cart

import "testing"

func TestParseQuantityAcceptsCanonicalValues(t *testing.T) {
	testCases := []struct {
		value   string
		minimum int64
		want    int64
	}{
		{value: "0", minimum: 0, want: 0},
		{value: "1", minimum: 1, want: 1},
		{value: "99", minimum: 1, want: 99},
	}
	for _, testCase := range testCases {
		quantity, valid := parseQuantity(testCase.value, testCase.minimum)
		if !valid || quantity != testCase.want {
			t.Errorf("parseQuantity(%q, %d) = (%d, %t), want (%d, true)", testCase.value, testCase.minimum, quantity, valid, testCase.want)
		}
	}
}

func TestParseQuantityRejectsInvalidValues(t *testing.T) {
	testCases := []struct {
		value   string
		minimum int64
	}{
		{value: "0", minimum: 1},
		{value: "00", minimum: 0},
		{value: "01", minimum: 0},
		{value: "-1", minimum: 0},
		{value: "+1", minimum: 0},
		{value: "1.0", minimum: 0},
		{value: " 1", minimum: 0},
		{value: "100", minimum: 0},
		{value: "", minimum: 0},
	}
	for _, testCase := range testCases {
		if _, valid := parseQuantity(testCase.value, testCase.minimum); valid {
			t.Errorf("parseQuantity(%q, %d) accepted an invalid quantity", testCase.value, testCase.minimum)
		}
	}
}
