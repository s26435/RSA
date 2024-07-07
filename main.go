package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
)

func gcd(a, b, x, y *big.Int) *big.Int {
	if a.Cmp(big.NewInt(0)) == 0 {
		x.SetInt64(0)
		y.SetInt64(1)
		return b
	}
	x1, y1 := new(big.Int), new(big.Int)
	temp := new(big.Int).Mod(b, a)
	g := gcd(temp, a, x1, y1)
	x.Set(y1)
	x.Sub(x, new(big.Int).Mul(new(big.Int).Div(b, a), x1))
	y.Set(x1)
	return g
}

func reverseModulo(a, b *big.Int) (*big.Int, error) {
	x, y := new(big.Int), new(big.Int)
	g := gcd(a, b, x, y)
	if g.Cmp(big.NewInt(1)) != 0 {
		return nil, errors.New("given numbers are not coprime")
	}
	x.Mod(x, b)
	if x.Cmp(big.NewInt(0)) < 0 {
		x.Add(x, b)
	}
	return x, nil
}

func eulerFunction(p, q *big.Int) *big.Int {
	temp1 := new(big.Int).Sub(p, big.NewInt(1))
	temp2 := new(big.Int).Sub(q, big.NewInt(1))
	return new(big.Int).Mul(temp1, temp2)
}

func rsa(base, exp, mod *big.Int) *big.Int {
	return new(big.Int).Exp(base, exp, mod)
}

func encryptToBase64(en []*big.Int) string {
	var combinedBytes []byte
	for _, x := range en {
		xBytes := x.Bytes()
		length := len(xBytes)
		lengthBytes := make([]byte, 4)
		lengthBytes[0] = byte((length >> 24) & 0xff)
		lengthBytes[1] = byte((length >> 16) & 0xff)
		lengthBytes[2] = byte((length >> 8) & 0xff)
		lengthBytes[3] = byte(length & 0xff)
		combinedBytes = append(combinedBytes, lengthBytes...)
		combinedBytes = append(combinedBytes, xBytes...)
	}
	return base64.StdEncoding.EncodeToString(combinedBytes)
}

func genKeys(bitLen int) (*big.Int, *big.Int, *big.Int) {
	p, err := rand.Prime(rand.Reader, bitLen)
	if err != nil {
		log.Fatal(err.Error())
	}
	q, err := rand.Prime(rand.Reader, bitLen)
	if err != nil {
		log.Fatal(err.Error())
	}
	n := new(big.Int).Mul(p, q)
	euler := eulerFunction(p, q)
	e := big.NewInt(2)
	temp1, temp2 := new(big.Int), new(big.Int)
	for gcd(e, euler, temp1, temp2).Cmp(big.NewInt(1)) != 0 {
		e.Add(e, big.NewInt(1))
	}
	d, err := reverseModulo(e, euler)
	if err != nil {
		log.Fatal(err.Error())
	}
	return d, n, e
}

func decryptFromBase64(base64Str string) ([]*big.Int, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	var en []*big.Int
	for i := 0; i < len(decodedBytes); {
		length := int(decodedBytes[i])<<24 | int(decodedBytes[i+1])<<16 | int(decodedBytes[i+2])<<8 | int(decodedBytes[i+3])
		i += 4
		en = append(en, new(big.Int).SetBytes(decodedBytes[i:i+length]))
		i += length
	}

	return en, nil
}

func encryptingRSA(e, n *big.Int, text string) string {
	var en []*big.Int
	for _, x := range text {
		t := new(big.Int).SetInt64(int64(x))
		en = append(en, rsa(t, e, n))
	}
	base64Encoded := encryptToBase64(en)
	return base64Encoded
}

func decryptingRSA(d, n *big.Int, en string) string {
	ta, err := decryptFromBase64(en)
	if err != nil {
		log.Fatal(err.Error())
	}
	var de []rune
	for _, x := range ta {
		de = append(de, rune(rsa(x, d, n).Int64()))
	}
	return string(de)
}

func readInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return input, nil
}

func readFile(filename string) (string, error) {
	var table string
	file, err := os.Open(filename)
	if err != nil {
		return table, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		table += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return table, err
	}
	return table, nil
}

func main() {
	bitLenPtr := flag.Int("len", 256, "Length of prime number used to generate keys")
	fromPtr := flag.Int("input", 0, "Input declaration {0,1,2}")
	messagePtr := flag.String("m", "Hello word", "Message input (if input = 0)")
	filePtr := flag.String("file", "text.txt", "File path when input = 2")
	flag.Parse()
	var text string
	var err error
	switch *fromPtr {
	case 0:
		text = *messagePtr
	case 1:
		text, err = readInput()
		if err != nil {
			log.Fatal(err.Error())
		}
	case 2:
		text, err = readFile(*filePtr)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		log.Fatal("invalid flag value")
	}
	d, n, e := genKeys(*bitLenPtr)
	fmt.Printf("decrypt key: %d\nencrypt key: %d\nmodulo: %d\n", d, e, n)
	en := encryptingRSA(e, n, text)
	fmt.Printf("Encrypted message %s\n", en)
	fmt.Println("Decrypted message:", decryptingRSA(d, n, en))
}
