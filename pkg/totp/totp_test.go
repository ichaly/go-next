package totp

import (
	"bufio"
	"encoding/base32"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"os"
	"testing"
	"time"
)

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

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
