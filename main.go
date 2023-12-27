package main

import (
	"fmt"
	"crypto/ecdh"
	"crypto/rand"
)

// Constants
const info string = "OMEMO X3DH"

// GenCurve returns a Curve implementing X25519 function
// over Curve25519
func GenCurve() ecdh.Curve{
	var curve25519 ecdh.Curve = ecdh.X25519()
	return curve25519
}

// Generates a public key from a random PrivateKey
func GenPublicKey(ellipse ecdh.Curve) {
	//random_private_key, _ := ecdh.Curve.GenerateKey(ellipse, rand.Reader)
	fmt.Println(rand.Reader)
}
/*
// Encode encodes an X25519 public key PK into a byte sequence
func Encode(PK) []byte {

}
*/
func main() {
	GenPublicKey(GenCurve())
}
