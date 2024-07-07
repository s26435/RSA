# RSA Encryption/Decryption Program

This program provides functionality for RSA encryption and decryption. It can generate RSA keys, encrypt messages, and decrypt encrypted messages. The program supports reading input from the command line, standard input, or a file, and it can output the results to the console or a file.

## Program Overview

### Flags

The program supports several flags to customize its behavior:

- `-dec`: Boolean flag to specify if you want to decrypt a message. If not set, the program will encrypt a message.
- `-len`: Integer flag to specify the length of the prime numbers used to generate keys. Default is 256.
- `-input`: Integer flag to specify the input source:
   - `0`: Message input from a string.
   - `1`: Message input from standard input.
   - `2`: Message input from a file.
- `-m`: String flag to specify the message input if `-input` is set to `0`. Default is "Hello world".
- `-file`: String flag to specify the file path when `-input` is set to `2`. Default is "text.txt".
- `-out`: String flag to specify the output file path. If not set, the results will be printed to the console.

### Functions

- **gcd(a, b, x, y)**: Computes the greatest common divisor of `a` and `b` and also finds `x` and `y` such that `a*x + b*y = gcd(a, b)`.
- **reverseModulo(a, b)**: Computes the modular multiplicative inverse of `a` modulo `b`.
- **eulerFunction(p, q)**: Computes Euler's totient function for the product of two primes `p` and `q`.
- **rsa(base, exp, mod)**: Performs modular exponentiation.
- **encryptToBase64(en)**: Encodes the encrypted big integers to a Base64 string.
- **genKeys(bitLen)**: Generates RSA keys of specified bit length.
- **decryptFromBase64(base64Str)**: Decodes a Base64 string to big integers.
- **encryptingRSA(e, n, text)**: Encrypts the text using RSA encryption.
- **decryptingRSA(d, n, en)**: Decrypts the encrypted text using RSA decryption.
- **readInput()**: Reads input from standard input.
- **readFile(filename)**: Reads input from a file.
- **writeToFile(path, text)**: Writes text to a file.

### Main Functions

- **encrypt()**: Handles the encryption process, including reading the input, generating keys, encrypting the message, and outputting the results.
- **decrypt()**: Handles the decryption process, including reading the input, parsing keys and the encrypted message, decrypting the message, and outputting the results.

### How to Use
1. **Clone the repository:**
   ```sh
   git clone https://github.com/s26435/RSA.git
   ```
2. **Build the program:**
   ```sh
   go build -o rsa_program main.go
   ```
3. **Encrypt a Message from a String**

   ```sh
   go run main.go -len 512 -input 0 -m "Your message here"
   ```

4. **Encrypt a Message from Standard Input**

   ```sh
   echo "Your message here" | go run main.go -len 512 -input 1
   ```

5. **Encrypt a Message from a File**

   ```sh
   go run main.go -len 512 -input 2 -file message.txt
   ```

6. **Decrypt a Message from a File**

   ```sh
   go run main.go -dec -input 2 -file encrypted.txt
   ```

### Example

To encrypt a message and save the result to a file:

```sh
go run main.go -len 512 -input 0 -m "Hello, RSA!" -out encrypted.txt
```

To decrypt the previously saved encrypted message:

```sh
go run main.go -dec -input 2 -file encrypted.txt
```

### Notes

- Ensure that the file paths provided with the `-file` and `-out` flags are correct and accessible.
- The program generates new RSA keys each time it runs in encryption mode. Save your keys if you need to use them again for decryption.