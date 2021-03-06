package main

import "fmt"

func main() {
	//starts with new line ends with new line
	fmt.Print("hello world", 1, 2)

	//variadic function similar to print
	fmt.Println("hello world")

	//same as println but it returns
	x := fmt.Sprintln("hello sprint")
	fmt.Println(x)

	//string interpolation
	// fmt.Printf() //used for interpolation

	// %v for the value
	// fmt.Sprintf() is similar to fmt.Printf() but it returns to a standard output

	// struct
	s := struct {
		name string
		age  int
	}{"John", 26}

	// channel
	c := make(chan int)

	// map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Printf("%v \n", true)           // boolean values //true
	fmt.Printf("%v \n", 132)            // signed and unsigned integers //132
	fmt.Printf("%v \n", 198.454)        // floating point numbers //198.454
	fmt.Printf("%v \n", 3+7i)           // complex numbers //(3+7i)
	fmt.Printf("%v \n", "Hello World!") // strings //Hello World!
	fmt.Printf("%v \n", []int{1, 2, 3}) // slices and arrays //[1 2 3]
	fmt.Printf("%v \n", m)              // maps //map[one:1 two:2]
	fmt.Printf("%v \n", s)              // structs //{John 26}
	fmt.Printf("%v \n", c)              // channels //0xc000062060
	fmt.Printf("%v \n", &s)             // pointers //&{John 26}

	//can use %#v to format value
	//can use %+v to get fields name from struct
	//%T is used to get type of the variable
	// %d converts hexa decimal to base10 format
	//%b to convert into binary format
	// %x hexadecimal format # adds leading 0 to hex value like 0 and %X for capital letters
	// %o octal format # adds leading 0 like 011 to octal value 11
	//%f formatting digits (%w.pf (w for width space and p for precision)) if we want default width we can skip w
	//%e scientific notation //can be used when a floating point number is too large to display
	//%g format floating point number to its significant digit
	//%s used for string format
	// %q used for escaped string format
	// %p pointer address format
	// if we want to have our custom value in %v of our custom type then we must implement string method
	// if we want to spik % we can use %%

}
