package api

import (
	"testing"
)

func TestVerifyKey_Success(t *testing.T) {
	api := &API{
		ScriptRunner: &ScriptRunner{},
	}
	verified, err := api.ScriptRunner.VerifySecretKey("super-secret-key", "super-secret-key")
	if err != nil {
		t.Errorf("error was found : %s", err.Error())
	}
	if !verified {
		t.Error("expected verified but seem not found")
	}
}
