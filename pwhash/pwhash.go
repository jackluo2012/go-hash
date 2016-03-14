package pwhash

import (
	"crypto/md5"
	"crypto/rand"
	"strings"
)

func Encode64(inp []byte, count int) string {
	const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var outp string
	cur := 0
	for cur < count {
		value := uint(inp[cur])
		cur += 1
		outp += string(itoa64[value&0x3f])
		if cur < count {
			value |= (uint(inp[cur]) << 8)
		}
		outp += string(itoa64[(value>>6)&0x3f])

		if cur >= count {
			break
		}
		cur += 1
		if cur < count {
			value |= (uint(inp[cur]) << 16)
		}
		outp += string(itoa64[(value>>12)&0x3f])
		if cur >= count {
			break
		}
		cur += 1
		outp += string(itoa64[(value>>18)&0x3f])
	}
	return outp
}

func CryptPrivate(pw, setting string) string {
	const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	//setting := "$P$BwQZDcQaNU9zAOF.6MOUdEhz9X68fL1"
	var outp = "*0"
	var count_log2 uint
	count_log2 = uint(strings.Index(itoa64, string(setting[3])))
	if count_log2 < 7 || count_log2 > 30 {
		return outp
	}
	count := 1 << count_log2
	salt := setting[4:12]
	if len(salt) != 8 {
		return outp
	}
	hasher := md5.New()
	hasher.Write([]byte(salt + pw))
	hx := hasher.Sum(nil)
	for count != 0 {
		hasher := md5.New()
		hasher.Write([]byte(string(hx) + pw))
		hx = hasher.Sum(nil)
		count -= 1
	}
	return setting[:12] + Encode64(hx, 16)
}

func PortableHashCheck(pw, storedHash string) bool {
	// pw: password to check (non-hashed), storedHash: Really? do you need an explanation?
	hx := CryptPrivate(pw, storedHash)
	return hx == storedHash
}

func GetRandomBytes(count int) ([]byte, error) {
	b := make([]byte, count)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenSaltPrivate(count int) (string, error) {
	b, err := GetRandomBytes(6)
	output := "$P$B"
	return output + Encode64(b, 6), err
}
