package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// program to convert base64 strings to PBKDF2-HMAC-SHA1 hashes
// written by cyclone
// v2023-02-16.1900; initial release
// v2023-03-04.1300; fixed typo

func versionFunc() {
	funcBase64Decode("Q3ljbG9uZSBiYXNlNjQgdG8gUEJLREYyLUhNQUMtU0hBMSBoYXNoIGNvbnZlcnRlciB2MjAyMy0wMy0wNC4xMzAwCg==")
}

// help function
func helpFunc() {
	versionFunc()
	str := "Example Usage:\n" +
		"\n./base64_2_pbkdf2sha1.bin -i base64.txt\n"
	fmt.Println(str)
	os.Exit(0)
}

func main() {
	// parse flag arguments
	inputFilePath := flag.String("i", "", "Path to input file")
	version := flag.Bool("version", false, "Program version:")
	cyclone := flag.Bool("cyclone", false, "")
	help := flag.Bool("help", false, "Prints help:")
	flag.Parse()

	// run sanity checks for -version & -help
	if *version == true {
		versionFunc()
		os.Exit(0)
	} else if *cyclone == true {
		funcBase64Decode("Q29kZWQgYnkgY3ljbG9uZSA7KQo=")
		os.Exit(0)
	} else if *help == true {
		helpFunc()
	}
	// run sanity checks on input (-i)
	if len(*inputFilePath) < 1 {
		fmt.Println("--> missing '-i input file' <--\n")
		helpFunc()
		os.Exit(0)
	}

	// open input file
	inputFile, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// create 10MB buffered reader / writer
	bufferSize := 10 * 1024 * 1024
	reader := bufio.NewReaderSize(inputFile, bufferSize)
	writer := bufio.NewWriterSize(os.Stdout, bufferSize)
	defer writer.Flush()

	// process each line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// decode base64 to bytes
		b, err := base64.StdEncoding.DecodeString(line)
		if err != nil {
			log.Println("Error decoding base64 string:", err)
			continue
		}

		// remove first byte (0x00 used by microsoft for auth)
		b = b[1:]

		// create salt from first 16 bytes
		salt := b[:16]

		// create hash from remaining 32 bytes
		hash := b[16:]

		// encode salt and hash back to base64 strings
		saltB64 := base64.StdEncoding.EncodeToString(salt)
		hashB64 := base64.StdEncoding.EncodeToString(hash)

		// parse hash as "sha1:1000:salt:hash"
		parsedHash := fmt.Sprintf("sha1:1000:%s:%s\n", saltB64, hashB64)

		// Write the result to standard output
		_, err = writer.WriteString(parsedHash)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// base64 decode function used for displaying encoded messages
func funcBase64Decode(line string) {
	str, err := base64.StdEncoding.DecodeString(line)
	if err != nil {
		log.Println("--> Text doesn't appear to be base64 encoded. <--")
		os.Exit(0)
	}
	fmt.Printf("%s\n", str)
}

// end
