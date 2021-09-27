// Illustrate Printing commands in GoLang
// package should be named main so that we can run from command line
package main

import "fmt"

func main() {
	fmt.Print("Print -  This is Simple Print ", "without", " formatting ", ". Multiple arguments can be passed ", "to the function without need of concatenation of string \n")
	fmt.Println("Println - Same Features as Print. Adds new line in the end of the print without using new line feed")

	fmt.Printf("Printf - %s format with arguments like Float: %f, Numeric: %d inserting/interpolation between the strings\n", "enables", 3.0, 11)

	// Printing Binary
	fmt.Printf("Printf - Binary: 4 => %b, 255 => %b \n",4, 255 )

	// Printing Hexadecimal
	fmt.Printf("Printf - Hexa: 255=> %x, 1060 => %#x \n",16, 17)

	// Repeating Arguments
	fmt.Printf("Printf - : %d %d, Reformat First Arg: %#[1]x, Reformat Second Arg: %[2]b \n",10, 11)

	// Pass Struct
	//TODO: Structs
	fmt.Printf("")

}
