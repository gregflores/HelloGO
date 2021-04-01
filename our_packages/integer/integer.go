package integer

import "fmt"

// Output int types
func Output() {
	var unsignedInteger uint = 4294967295
	var integer int = 2147483647
	var unsignedInteger8 uint8 = 255
	var unsignedInteger16 uint16 = 65535
	var unsignedInteger32 uint32 = 4294967295
	var unsignedInteger64 uint64 = 18446744073709551615
	var signedInteger8 int8 = 127
	var signedInteger16 int16 = 32767
	var signedInteger32 int32 = 2147483647
	var signedInteger64 int64 = 9223372036854775807

	fmt.Println(unsignedInteger)
	fmt.Println(integer)
	fmt.Println(unsignedInteger8)
	fmt.Println(unsignedInteger16)
	fmt.Println(unsignedInteger32)
	fmt.Println(unsignedInteger64)
	fmt.Println(signedInteger8)
	fmt.Println(signedInteger16)
	fmt.Println(signedInteger32)
	fmt.Println(signedInteger64)
}