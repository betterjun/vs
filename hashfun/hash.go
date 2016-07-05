package hashfun

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/crc32"
)

// Define the supported hash methods.
const (
	MD5 = iota
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	HASH_END
)

type hx func() hash.Hash

var hxfun []hx

func init() {
	hxfun = make([]hx, HASH_END)
	hxfun[MD5] = md5.New
	hxfun[SHA1] = sha1.New
	hxfun[SHA224] = sha256.New224
	hxfun[SHA256] = sha256.New
	hxfun[SHA384] = sha512.New384
	hxfun[SHA512] = sha512.New
}

// Get the hash string in bytes.
func Hash(src []byte, method int) []byte {
	if method < MD5 || method >= HASH_END {
		return nil
	}

	h := hxfun[method]()
	signature := make([]byte, h.Size()*2)
	h.Write(src)

	hex.Encode(signature, h.Sum(nil))
	return signature
}

// Get the hash string.
func HashString(src string, method int) string {
	if method < MD5 || method >= HASH_END {
		return ""
	}

	h := hxfun[method]()
	signature := make([]byte, h.Size()*2)
	h.Write([]byte(src))

	hex.Encode(signature, h.Sum(nil))
	return string(signature)
}

// Get the IEEE crc32 value.
func GetCRC32(src []byte) []byte {
	h := crc32.NewIEEE()
	signature := make([]byte, h.Size()*2)
	h.Write([]byte(src))

	hex.Encode(signature, h.Sum(nil))
	return signature
}

func HmacSha1(data, key []byte) []byte {
	h := hxfun[SHA1]
	hm := hmac.New(h, key)
	hm.Write(data)
	return hm.Sum(nil)
}

func HmacSha1String(data, key string) string {
	h := hxfun[SHA1]
	hm := hmac.New(h, []byte(key))
	signature := make([]byte, hm.Size()*2)
	hm.Write([]byte(data))
	hex.Encode(signature, hm.Sum(nil))
	return string(signature)
}
