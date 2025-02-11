package main

// ----------------------------------------------------------------------------
// byteToMegabyte converts bytes to megabytes by dividing by 1024 twice
// to get from bytes -> kilobytes -> megabytes
// Parameters:
//   - b: number of bytes to convert as uint64
//
// Returns:
//   - uint64 representing the equivalent number of megabytes
func byteToMegabyte(b uint64) uint64 {
	return b / 1024 / 1024
}
