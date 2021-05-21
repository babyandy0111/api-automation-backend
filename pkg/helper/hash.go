package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/scrypt"
)

// Sha1Str 將字串轉換為 sha1 hash
func Sha1Str(str string) string {
	h := sha1.New()
	_, _ = h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func ScryptStr(str string) string {
	salt := os.Getenv("SCRYPT_SALT")
	secret := []byte(salt)
	dk, _ := scrypt.Key([]byte(str), secret, 16384, 8, 1, 32)
	return fmt.Sprintf("%x", dk)
}

// Md5Str 將字串轉換為 md5 hash
func Md5Str(str string) (string, error) {
	m := md5.New()
	_, err := io.WriteString(m, str)
	hash := fmt.Sprintf("%x", m.Sum(nil))
	return hash, err
}

func Sha1Hex(str string) string {
	h := sha1.New()
	_, _ = h.Write([]byte(str))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}
