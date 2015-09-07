package main

import (
	"fmt"
	"io/ioutil"
)


// Crypt function, takes a byte array and a number
func crypt(a []byte, num int) string {
		num %= 256						// If num is over 255, reduce it using its reminder
		var b []byte					// Define landing array for en/decrypted characters
		for _, c := range a {			// Cycle trough all characters
		d := int(c)						// make d a integer representative of the byte value
		e := (d + num) % 256			// Add the en/decryption number (Use the reminder to avoid out of range numbers
		b = append(b,byte(e))
	}
	return string(b)
}

// Encrypts a string with num rotation
func encrypt(s string, num int) (res string) {
	a := []byte(s)
	res = crypt(a, num)
	return
}

// Decrypts a string encrypted with num rotation
func decrypt(s string, num int) (res string) {
	a := []byte(s)
	res = crypt(a, -num)
	return
}

// Encrypts a file with num rotation
func encryptFile(path string, num int) {
	a, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File Error: %s\n", err)
	}
	err = ioutil.WriteFile("encrypted.txt", []byte(crypt(a, num)), 0777)
}

// Decrypts a file encrypted with num rotation
func decryptFile(path string, num int) {
	a, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File Error: %s\n", err)
	}
	err = ioutil.WriteFile("decrypted.txt", []byte(crypt(a, -num)), 0777)

}

// Supidly easy bruteforce printer - to be improved
func crack(cipher string, size int) {
	for i := 1; i < size+1; i++ {
		fmt.Println("Key", i, decrypt(cipher, i))
	}
}

func main() {
	// Examples
	fmt.Println(encrypt("This is a test message", 10))
	fmt.Println(decrypt("^rs}*s}*k*~o}~*wo}}kqo", 10))
	crack("^rs}*s}*k*~o}~*wo}}kqo", 20)
	encryptFile("test.txt", 78)
	decryptFile("encrypted.txt", 78)
}