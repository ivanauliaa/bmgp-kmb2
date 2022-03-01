package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	for _, value := range s.name {
		nameEncode += alphaToCipher[string(value)]
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for _, value := range s.nameEncode {
		nameDecode += cipherToAlpha[string(value)]
	}

	return nameDecode
}

var (
	key           = 4
	alphabets     = "abcdefghijklmnopqrstuvwxyz"
	alphaToCipher = map[string]string{}
	cipherToAlpha = map[string]string{}
)

func init() {
	for i, alphabet := range alphabets {
		index := (i + key) % len(alphabets)
		alphaToCipher[string(alphabet)] = string(alphabets[index])
	}

	for key, value := range alphaToCipher {
		cipherToAlpha[string(value)] = string(key)
	}
}

func main() {
	var menu int

	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
