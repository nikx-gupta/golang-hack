package main

import (
	"fmt"
	"unicode"
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23"
	fmt.Println("String length: ", len(sLiteral))
	for i := 0; i < len(sLiteral); i++ {
		fmt.Println("Char at Index: ", i, " is :", sLiteral[i])
	}

	const sUniCode = "\u2318\u2317\u1319"
	// Print Unicode Bytes
	fmt.Printf("% x \n", sUniCode)
	fmt.Println("Length: ", len(sUniCode))

	for i := 0; i < len(sLiteral); i++ {
		if unicode.IsPrint(rune(sLiteral[i])) {
			fmt.Printf("%c \n", sLiteral[i])
		}else{
			fmt.Println("Not a unicode char: ", sLiteral[i])
		}
	}
	//fmt.Printf("x: %x with Length: %d", sLiteral, len(sLiteral))
	//
	//for i := 0; i < len(sLiteral); i++ {
	//	fmt.Printf("%x \n", sLiteral[i])
	//}
	//
	//fmt.Printf("q: %q\n", sLiteral)
	//fmt.Printf("q: %+q\n", sLiteral)

}
