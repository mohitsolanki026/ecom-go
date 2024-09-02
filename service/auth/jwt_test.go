package auth

import "testing"

func TestJwt(t *testing.T) {
	token,err := CreateJwtToken(1);
	if err != nil {
		t.Errorf("error while crating jwt: %v", err)
	}

	if token == ""{
		t.Error("expected error not to be empty")
	}
}
