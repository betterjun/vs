package hashfun

import (
	"bytes"
	"testing"
)

func TestHash(t *testing.T) {
	src := []byte("你好，世界")

	var tests = []struct {
		input int
		want  []byte
	}{
		{MD5, []byte("dbefd3ada018615b35588a01e216ae6e")},
		{SHA1, []byte("3becb03b015ed48050611c8d7afe4b88f70d5a20")},
		{SHA224, []byte("9a65a12818b8e6ac357cee9337565337f55bda8a45b0c1bfb9f4403c")},
		{SHA256, []byte("46932f1e6ea5216e77f58b1908d72ec9322ed129318c6d4bd4450b5eaab9d7e7")},
		{SHA384, []byte("fbea16d8be2993f2cda1ef9fc055f53f0fa23f1e1dc4a57a7548c36227c3ef0491484fcf1e30c5d1ff17441a5ce89a11")},
		{SHA512, []byte("45a6e3fe78af4a3326da9bf8c3407bca5fef80b334c046d20544b0b28be6c761718cfaf5b752eaa89849b83a4d4e5f6df4908e195cd8c159181e78971910db13")},
	}

	for _, test := range tests {
		got := Hash(src, test.input)
		if !bytes.Equal(got, test.want) {
			t.Errorf("Hash(%q, %q) == %q, want %q", src, hashName[test.input], got, test.want)
		} else {
			t.Logf("Hash(%q, %q) passed", src, hashName[test.input])
		}
	}
}

type testHashData struct {
	input string
	want  string
}

var hashName = []string{"MD5", "SHA1", "SHA224", "SHA256", "SHA384", "SHA512"}

func hashTest(hash int, tests []testHashData, t *testing.T) {
	for _, test := range tests {
		got := HashString(test.input, hash)
		if got != test.want {
			t.Errorf("HashString(%q, %s) == %q, want %q", test.input, hashName[hash], got, test.want)
		} else {
			t.Logf("HashString(%q, %s) passed", test.input, hashName[hash])
		}
	}
}

func TestMD5(t *testing.T) {
	var tests = []testHashData{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"abc", "900150983cd24fb0d6963f7d28e17f72"},
		{"你好，世界", "dbefd3ada018615b35588a01e216ae6e"},
	}
	hashTest(MD5, tests, t)
}

func TestSHA1(t *testing.T) {
	var tests = []testHashData{
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"abc", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"你好，世界", "3becb03b015ed48050611c8d7afe4b88f70d5a20"},
	}
	hashTest(SHA1, tests, t)
}

func TestSHA224(t *testing.T) {
	var tests = []testHashData{
		{"", "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f"},
		{"abc", "23097d223405d8228642a477bda255b32aadbce4bda0b3f7e36c9da7"},
		{"你好，世界", "9a65a12818b8e6ac357cee9337565337f55bda8a45b0c1bfb9f4403c"},
	}
	hashTest(SHA224, tests, t)
}

func TestSHA256(t *testing.T) {
	var tests = []testHashData{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{"你好，世界", "46932f1e6ea5216e77f58b1908d72ec9322ed129318c6d4bd4450b5eaab9d7e7"},
	}
	hashTest(SHA256, tests, t)
}

func TestSHA384(t *testing.T) {
	var tests = []testHashData{
		{"", "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b"},
		{"abc", "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7"},
		{"你好，世界", "fbea16d8be2993f2cda1ef9fc055f53f0fa23f1e1dc4a57a7548c36227c3ef0491484fcf1e30c5d1ff17441a5ce89a11"},
	}
	hashTest(SHA384, tests, t)
}

func TestSHA512(t *testing.T) {
	var tests = []testHashData{
		{"", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		{"abc", "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"},
		{"你好，世界", "45a6e3fe78af4a3326da9bf8c3407bca5fef80b334c046d20544b0b28be6c761718cfaf5b752eaa89849b83a4d4e5f6df4908e195cd8c159181e78971910db13"},
	}
	hashTest(SHA512, tests, t)
}

func TestCRC32(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte(""), []byte("00000000")},
		{[]byte("abc"), []byte("352441c2")},
		{[]byte("你好，世界"), []byte("acf5da54")},
	}

	for _, test := range tests {
		got := GetCRC32(test.input)
		if !bytes.Equal(got, test.want) {
			t.Errorf("GetCRC32(%q) == %q, want %q", test.input, got, test.want)
		} else {
			t.Logf("GetCRC32(%q) passed", test.input)
		}
	}
}
