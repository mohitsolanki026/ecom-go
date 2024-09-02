package auth

import "testing"

func TestHash(t *testing.T) {
	hash,err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password %v:", err)
	}

	if hash == "password"{
		t.Errorf("hash can't be same as password")
	}

	if hash == ""{
		t.Errorf("hash can't be empty")
	}
}

func TestComparePassword(t *testing.T){
	hash, err := HashPassword("password")
	
	if err != nil {
		t.Errorf("error hashing password %v:", err)
	}
	if !ComparePassword(hash,[]byte("password")){
		t.Error("expected password to match hash")
	}

	if ComparePassword(hash,[]byte("passwordd")){
		t.Error("expected password to not match hash")
	}

}
