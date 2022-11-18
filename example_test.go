package size_test

import (
	"fmt"

	"go.withmatt.com/size"
)

func Example() {
	fmt.Println(1 * size.Gigabyte)
	// Output: 1G
}

// Example demonstrates couting the number of units in a Capacity, divide them.
func ExampleDividingUnits() {
	fmt.Println(int(size.Kilobyte/size.Byte), "bytes in a kilobyte")
	fmt.Println(int(size.Gigabyte/size.Kilobyte), "kilobytes in a gigabyte")
	fmt.Println(int(5*size.Gigabyte/size.Kilobyte), "kilobytes in 5 gigabytes")
	// Output: 1000 bytes in a kilobyte
	// 1000000 kilobytes in a gigabyte
	// 5000000 kilobytes in 5 gigabytes
}

// Example demonstrates formatting a Capacity as a string
func ExampleFormatting() {
	fmt.Println(5 * size.Gigabyte)
	fmt.Println(3 * size.Gibibyte)
	// Output: 5G
	// 3.2G
}

// Example demonstrates
func ExampleParseCapacity() {
	fmt.Println(size.ParseCapacity("1G"))
	// Output: 1G <nil>
}
