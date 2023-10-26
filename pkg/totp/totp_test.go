package totp

import (
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"testing"
	"time"
)

func TestTotp(t *testing.T) {
	secret := base32.StdEncoding.EncodeToString([]byte("15210203617"))
	opt := totp.ValidateOpts{
		Skew:      1,
		Period:    5 * 60,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	}
	code, err := totp.GenerateCodeCustom(secret, time.Now(), opt)
	if err != nil {
		panic(err)
	}
	valid, _ := totp.ValidateCustom(code, secret, time.Now(), opt)
	if valid {
		t.Log("Valid code!")
	} else {
		t.Log("Invalid code!")
	}
	t.Log(code)
	t.Log(secret)
}
