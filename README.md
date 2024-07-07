# RSA Encryption/Decryption in Go

This repository contains a Go program that demonstrates RSA encryption and decryption. The program generates RSA keys, encrypts a given text message using the public key, and decrypts it back to the original message using the private key. The encrypted message is also encoded in Base64 format for safe transmission.

## Features

- Generation of RSA key pairs (public and private keys).
- Encryption of a text message using the RSA public key.
- Decryption of the encrypted message using the RSA private key.
- Encoding of the encrypted message in Base64 format.
- Decoding of the Base64 encoded message back to the encrypted form.
- Reading input from the command line, standard input, or a file.

## Dependencies

The program uses the following Go packages:
- `crypto/rand` for generating random prime numbers.
- `encoding/base64` for encoding and decoding the encrypted message.
- `math/big` for handling large integer calculations.
- `bufio` for reading input from files and standard input.
- `flag` for command-line flag parsing.
- `os` for handling file operations.

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

### readInput() (string, error)
Reads input from standard input.

### readFile(filename string) (string, error)
Reads input from a specified file.

## Usage

1. Clone the repository:
```sh
git clone <repository_url>
```

2. Build the program:
```sh
go build -o rsa_program main.go
```

4. Run the program with desired flags:
```sh
./rsa_program -len <bit_length> -input <input_type> -m <message> -file <file_path>
```

### Flags

- `-len` (int): Length of prime numbers used to generate keys. Default is 256.
- `-input` (int): Input source, where 0 indicates command-line message, 1 indicates standard input, and 2 indicates a file. Default is 0.
- `-m` (string): Message to be encrypted if `input` is 0. Default is "Hello world".
- `-file` (string): File path to read the message from if `input` is 2. Default is "text.txt".

### Example Commands

1. Using a command-line message:
   ```sh
   ./rsa_program -len 512 -input 0 -m "Yes? It's fantastic"
   ```

2. Using standard input:
   ```sh
   echo "Yes? It's fantastic" | ./rsa_program -len 512 -input 1
   ```

3. Using a file:
   ```sh
   ./rsa_program -len 512 -input 2 -file "path/to/your/text.txt"
   ```

## Example Output

```
decrypt key: 10123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
encrypt key: 65537
modulo: 12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
Encrypted message: UGVpT21HZXZKa0ttT09jQW9kTXlFb1J2d2dnU3cwUlpFY1dxdGVNR2dFTGZ5UG9BYz0=
Decrypted message: Yes? It's fantastic
```
