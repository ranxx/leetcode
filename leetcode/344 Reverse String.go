package main

/*
Example 1:
	Input: ["h","e","l","l","o"]
	Output: ["o","l","l","e","h"]
*/

func ReverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[j], s[i] = s[i], s[j]
		j--
		i++
	}
}

func ReverseStringV2(s []byte) {
	sTemp := make([]byte, len(s))
	for i, v := range s {
		sTemp[i] = v
	}
	for i, sLen := len(s)-1, 0; i >= 0; {
		s[sLen] = sTemp[i]
		i--
		sLen++
	}
}

func ReverseStringV3(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[len(s)-1-i], s[i] = s[i], s[len(s)-1-i]
	}
}
