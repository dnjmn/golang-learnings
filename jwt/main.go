package main

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

func hell() {
	key, _ := 
	pubKey, _ := 
	badKey := []byte("All your base are belong to key")

	// Test parsePrivateKey
	if _, e := jwt.ParseRSAPrivateKeyFromPEM(key); e != nil {
		t.Errorf("Failed to parse valid private key: %v", e)
	}

	if k, e := jwt.ParseRSAPrivateKeyFromPEM(pubKey); e == nil {
		t.Errorf("Parsed public key as valid private key: %v", k)
	}

	if k, e := jwt.ParseRSAPrivateKeyFromPEM(badKey); e == nil {
		t.Errorf("Parsed invalid key as valid private key: %v", k)
	}

	if _, e := jwt.ParseRSAPrivateKeyFromPEMWithPassword(secureKey, "password"); e != nil {
		t.Errorf("Failed to parse valid private key with password: %v", e)
	}

	if k, e := jwt.ParseRSAPrivateKeyFromPEMWithPassword(secureKey, "123132"); e == nil {
		t.Errorf("Parsed private key with invalid password %v", k)
	}

	// Test parsePublicKey
	if _, e := jwt.ParseRSAPublicKeyFromPEM(pubKey); e != nil {
		t.Errorf("Failed to parse valid public key: %v", e)
	}

	if k, e := jwt.ParseRSAPublicKeyFromPEM(key); e == nil {
		t.Errorf("Parsed private key as valid public key: %v", k)
	}

	if k, e := jwt.ParseRSAPublicKeyFromPEM(badKey); e == nil {
		t.Errorf("Parsed invalid key as valid private key: %v", k)
	}
}
