package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("언어"))
	fmt.Println(HasConsonantSuffix("그래"))
	fmt.Println(HasConsonantSuffix("안녕 디지몬"))
	// Output:
	// false
	// true
	// true
}

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}

	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_modifyBytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}
