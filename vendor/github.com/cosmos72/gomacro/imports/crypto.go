// this file was generated by gomacro command: import _b "crypto"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto"
	"io"
)

// reflection: allow interpreted code to import "crypto"
func init() {
	Packages["crypto"] = Package{
	Binds: map[string]Value{
		"BLAKE2b_256":	ValueOf(crypto.BLAKE2b_256),
		"BLAKE2b_384":	ValueOf(crypto.BLAKE2b_384),
		"BLAKE2b_512":	ValueOf(crypto.BLAKE2b_512),
		"BLAKE2s_256":	ValueOf(crypto.BLAKE2s_256),
		"MD4":	ValueOf(crypto.MD4),
		"MD5":	ValueOf(crypto.MD5),
		"MD5SHA1":	ValueOf(crypto.MD5SHA1),
		"RIPEMD160":	ValueOf(crypto.RIPEMD160),
		"RegisterHash":	ValueOf(crypto.RegisterHash),
		"SHA1":	ValueOf(crypto.SHA1),
		"SHA224":	ValueOf(crypto.SHA224),
		"SHA256":	ValueOf(crypto.SHA256),
		"SHA384":	ValueOf(crypto.SHA384),
		"SHA3_224":	ValueOf(crypto.SHA3_224),
		"SHA3_256":	ValueOf(crypto.SHA3_256),
		"SHA3_384":	ValueOf(crypto.SHA3_384),
		"SHA3_512":	ValueOf(crypto.SHA3_512),
		"SHA512":	ValueOf(crypto.SHA512),
		"SHA512_224":	ValueOf(crypto.SHA512_224),
		"SHA512_256":	ValueOf(crypto.SHA512_256),
	}, Types: map[string]Type{
		"Decrypter":	TypeOf((*crypto.Decrypter)(nil)).Elem(),
		"DecrypterOpts":	TypeOf((*crypto.DecrypterOpts)(nil)).Elem(),
		"Hash":	TypeOf((*crypto.Hash)(nil)).Elem(),
		"PrivateKey":	TypeOf((*crypto.PrivateKey)(nil)).Elem(),
		"PublicKey":	TypeOf((*crypto.PublicKey)(nil)).Elem(),
		"Signer":	TypeOf((*crypto.Signer)(nil)).Elem(),
		"SignerOpts":	TypeOf((*crypto.SignerOpts)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"Decrypter":	TypeOf((*P_crypto_Decrypter)(nil)).Elem(),
		"Signer":	TypeOf((*P_crypto_Signer)(nil)).Elem(),
		"SignerOpts":	TypeOf((*P_crypto_SignerOpts)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for crypto.Decrypter ---------------
type P_crypto_Decrypter struct {
	Object	interface{}
	Decrypt_	func(_proxy_obj_ interface{}, rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
	Public_	func(interface{}) crypto.PublicKey
}
func (P *P_crypto_Decrypter) Decrypt(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
	return P.Decrypt_(P.Object, rand, msg, opts)
}
func (P *P_crypto_Decrypter) Public() crypto.PublicKey {
	return P.Public_(P.Object)
}

// --------------- proxy for crypto.Signer ---------------
type P_crypto_Signer struct {
	Object	interface{}
	Public_	func(interface{}) crypto.PublicKey
	Sign_	func(_proxy_obj_ interface{}, rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error)
}
func (P *P_crypto_Signer) Public() crypto.PublicKey {
	return P.Public_(P.Object)
}
func (P *P_crypto_Signer) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	return P.Sign_(P.Object, rand, digest, opts)
}

// --------------- proxy for crypto.SignerOpts ---------------
type P_crypto_SignerOpts struct {
	Object	interface{}
	HashFunc_	func(interface{}) crypto.Hash
}
func (P *P_crypto_SignerOpts) HashFunc() crypto.Hash {
	return P.HashFunc_(P.Object)
}
