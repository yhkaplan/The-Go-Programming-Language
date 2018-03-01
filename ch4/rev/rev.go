// reverse reverses a slice of ints in place
package rev

// Reverse 1
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Reverse 2: Use array pointer instead of a slice
//func reverse2(s *[]int) {
//	*s = make([Len(s)])
//	return *newS
//}
