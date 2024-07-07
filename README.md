# RSA Encryption/Decryption in Go

This repository contains a Go program that demonstrates RSA encryption and decryption. The program generates RSA keys, encrypts a given text message using the public key, and decrypts it back to the original message using the private key. The encrypted message is also encoded in Base64 format for safe transmission.

## Features

- Generation of RSA key pairs (public and private keys).
- Encryption of a text message using the RSA public key.
- Decryption of the encrypted message using the RSA private key.
- Encoding of the encrypted message in Base64 format.
- Decoding of the Base64 encoded message back to the encrypted form.

## Dependencies

The program uses the following Go packages:
- `crypto/rand` for generating random prime numbers.
- `encoding/base64` for encoding and decoding the encrypted message.
- `math/big` for handling large integer calculations.

## Functions

### gcd(a, b, x, y *big.Int) *big.Int
Computes the greatest common divisor of `a` and `b` using the Extended Euclidean Algorithm. It also finds `x` and `y` such that `a*x + b*y = gcd(a, b)`.

### reverseModulo(a, b *big.Int) (*big.Int, error)
Finds the modular inverse of `a` modulo `b`, if it exists.

### eulerFunction(p, q *big.Int) *big.Int
Calculates Euler's totient function for two prime numbers `p` and `q`.

### rsa(base, exp, mod *big.Int) *big.Int
Performs RSA encryption/decryption by computing `(base^exp) % mod`.

### encryptToBase64(en []*big.Int) string
Encodes a slice of `*big.Int` values into a Base64 string.

### decryptFromBase64(base64Str string) ([]*big.Int, error)
Decodes a Base64 string back into a slice of `*big.Int` values.

### genKeys(bitLen int) (*big.Int, *big.Int, *big.Int)
Generates an RSA key pair (private key, modulus, public key) with the specified bit length.

### encryptingRSA(e, n *big.Int, text string) string
Encrypts a given text message using the RSA public key (`e`, `n`) and returns the Base64 encoded encrypted message.

### decryptingRSA(d, n *big.Int, en string) string
Decrypts a Base64 encoded encrypted message using the RSA private key (`d`, `n`) and returns the original text message.

## Usage

1. Clone the repository:
```sh
git clone https://github.com/s26435/RSA.git
```

2. Run the program:
```sh
go run main.go
```

The program will output the generated RSA keys, the encrypted message, and the decrypted message.

## Example

```
decrypt key: 10123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
encrypt key: 65537
modulo: 12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
Encrypted message: UGVpT21HZXZKa0ttT09jQW9kTXlFb1J2d2dnU3cwUlpFY1dxdGVNR2dFTGZ5UG9BYz0=
Decrypted message: Yes? Its fantastic
```
