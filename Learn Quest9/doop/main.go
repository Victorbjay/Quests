package main

import "os"

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	b, ok2 := atoi(os.Args[3])
	if !ok1 || !ok2 {
		return
	}

	switch os.Args[2] {
	case "+":
		if (b > 0 && a > 9223372036854775807-b) || (b < 0 && a < -9223372036854775808-b) {
			return
		}
		os.Stdout.WriteString(itoa(a+b) + "\n")

	case "-":
		if (b < 0 && a > 9223372036854775807+b) || (b > 0 && a < -9223372036854775808+b) {
			return
		}
		os.Stdout.WriteString(itoa(a-b) + "\n")

	case "*":
		if a != 0 && b != 0 && (a*b)/b != a {
			return
		}
		os.Stdout.WriteString(itoa(a*b) + "\n")

	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		os.Stdout.WriteString(itoa(a/b) + "\n")

	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		os.Stdout.WriteString(itoa(a%b) + "\n")
	}
}

func atoi(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}

	sign := int64(1)
	i := 0
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	var n int64
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int64(c-'0')
	}

	return n * sign, true
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	var b []byte
	for n > 0 {
		b = append([]byte{byte(n%10) + '0'}, b...)
		n /= 10
	}

	return sign + string(b)
}
