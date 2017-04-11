package util

func Reverse(bytes string) string {
	strLength := len(bytes)
	reversed := make([]byte, strLength, strLength)
	for i := 0; i < strLength; i++ {
		reversed[i] = bytes[strLength-i-1]
	}
	return string(reversed)
}
