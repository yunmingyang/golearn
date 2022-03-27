package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)


func main() {
	// rune could forcibly transform to string
	// len() will return number of bytes, not rune number
	s := "蝈蝈s果果"
	fmt.Println(len(s))
	fmt.Printf("%b, %c\n", s[0], s[6])

	// // panic out of range
	// c := s[len(s)]
	// fmt.Println(c)

	s1 := "hello, world!"
	// Not include right side of the range
	fmt.Println(s1[0:5])
	fmt.Println(s1[:5])
	fmt.Println(s1[7:])
	fmt.Println(s1[:])
	fmt.Println("goodbye " + s1[:5])

	// // string can not be modified
	// s1[0] = "1"

	var s2 string
	s2 += "`"
	fmt.Printf("%s\n", s2)

	var s3 string = "asdd`"
	fmt.Printf("%s\n", s3)

	const GoUsage = `Go is a tool for managing Go source code.

	Usage:
		go command [arguments]`
	fmt.Println(GoUsage)

	fmt.Printf("%s %s %s %s\n", "世界", "\xe4\xb8\x96\xe7\x95\x8c", "\u4e16\u754c", "\U00004e16\U0000754c")
	// Though 4e16 is 2 bytes, but after transformation, it is 19990 in decimal, so according UTF-8, it is expressed by 3 bytes
	fmt.Println(len("\u4e16"))
	fmt.Println(len("\U00004e16"))

	s4 := "hello, 世界"
	fmt.Println(len(s4))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "hello, 世界" {
		fmt.Printf("%d\t%q\t%c\n", i, r, r)
	}

	// // invalid UTF-8 character, use \uFFFD, it is usually a black hexagon with a question mark in it
	// fmt.Println(string(0234444))

	fmt.Println(basename1("a/b/c.go"))
	fmt.Println(basename1("c.d.go"))
	fmt.Println(basename1("abc"))

	fmt.Println(basename2("a/b/c.go"))
	fmt.Println(basename2("c.d.go"))
	fmt.Println(basename2("abc"))

	fmt.Println(comma("12345"))

	fmt.Println(intstoString([]int{1, 2, 3}))

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))

	x1, errx := strconv.Atoi("123")
	y1, erry := strconv.ParseInt("123", 10, 64) // The third parameters is used as indicating size of integer, 0 is int
	fmt.Println(x1, errx, y1, erry)
}


// func hasPrefix(s, prefix string) bool {
// 	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
// }

// func hasSuffix(s, suffix string) bool {
// 	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
// }

// func contains(s, substr string) bool {
// 	for i := 0; i < len(s); i ++ {
// 		if hasPrefix(s[i:], substr){
// 			return true
// 		}
// 	}
// 	return false
// }

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1: ]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func intstoString(values []int) string {
	var buf bytes.Buffer

	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}