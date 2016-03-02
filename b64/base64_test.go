package b64

import "testing"

func isByteSliceEqual(lp, rp []byte) bool {
	if len(lp) != len(rp) {
		return false
	}

	for i := range lp {
		if lp[i] != rp[i] {
			return false
		}
	}

	return true
}

func TestEncode(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("你好，世界..."), []byte("5L2g5aW977yM5LiW55WMLi4u")},
	}

	for _, test := range tests {
		got := Encode(test.input)
		if !isByteSliceEqual(got, test.want) {
			t.Errorf("Encode(%q) == %q, want %q", test.input, got, test.want)
		}
	}
}

func TestDecode(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("5L2g5aW977yM5LiW55WMLi4u"), []byte("你好，世界...")},
	}

	for _, test := range tests {
		got, err := Decode(test.input)
		if err != nil || !isByteSliceEqual(got, test.want) {
			t.Errorf("Decode(%q) == %q, want %q", test.input, got, test.want)
		}
	}
}

func TestStringEncodeDecode(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"你好，世界...", "5L2g5aW977yM5LiW55WMLi4u"},
	}

	for _, test := range tests {
		got := EncodeString(test.input)
		if got != test.want {
			t.Errorf("EncodeString(%q) == %q, want %q", test.input, got, test.want)
		}
	}

	for _, test := range tests {
		got, err := DecodeString(test.want)
		if err != nil || got != test.input {
			t.Errorf("DecodeString(%q) == %q, want %q", test.want, got, test.input)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode([]byte("你好，世界..."))
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode([]byte("5L2g5aW977yM5LiW55WMLi4u"))
	}
}

func BenchmarkEncodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeString("你好，世界...")
	}
}

func BenchmarkDecodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecodeString("5L2g5aW977yM5LiW55WMLi4u")
	}
}
