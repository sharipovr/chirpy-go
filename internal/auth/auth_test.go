package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	fakeUUID := uuid.New()
	timeDuration, _ := time.ParseDuration("1m")
	madeJWT, err := MakeJWT(fakeUUID, "signing-key", timeDuration)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ValidateJWT(madeJWT, "wrong_key")
	if err == nil {
		t.Errorf("Wrong secret. %v", err)
	}

	validUUID, err := ValidateJWT(madeJWT, "signing-key")
	if err != nil {
		t.Fatal(err)
	}
	if fakeUUID != validUUID {
		t.Errorf("Wrong UUID")
	}

	timeDuration, _ = time.ParseDuration("-1m")
	madeJWT, err = MakeJWT(fakeUUID, "signing-key", timeDuration)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ValidateJWT(madeJWT, "signing-key")
	if err == nil {
		t.Errorf("expected expired token to be rejected")
	}
}
