package utility

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/chrislim1914/xrp-transaction/api/models"
	"golang.org/x/crypto/bcrypt"
)

func RandomStrGenerator(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func GenerateSecretKey(r models.Accounts) string {
	secret := r.Email
	payload, _ := json.Marshal(r)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
