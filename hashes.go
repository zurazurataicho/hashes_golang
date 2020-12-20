package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"log"
	"os"
)

func main() {
	fp, err := os.OpenFile("hashes.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Cannot create text file")
	}
	defer fp.Close()

	message := []byte("This is a test program for hashing algorithm libraries in Go.")

	// MD5
	byteMd5 := md5.Sum(message)
	strMd5 := fmt.Sprintf("%x\n", byteMd5)
	fp.WriteString(strMd5)

	// SHA1
	byteSha1 := sha1.Sum(message)
	strSha1 := fmt.Sprintf("%x\n", byteSha1)
	fp.WriteString(strSha1)

	// SHA256
	byteSha256 := sha256.Sum256(message)
	strSha256 := fmt.Sprintf("%x\n", byteSha256)
	fp.WriteString(strSha256)

	// SHA512
	byteSha512 := sha512.Sum512(message)
	strSha512 := fmt.Sprintf("%x\n", byteSha512)
	fp.WriteString(strSha512)
}
