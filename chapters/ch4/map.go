package ch4

import "fmt"

func Map() {
	var m map[rune]int
	m = map[rune]int{}
	m = make(map[rune]int)
	m = make(map[rune]int, 0)
	fmt.Printf("%v\n", m)

	//addition
	m['中'] = 1
	m['国'] += 1
	fmt.Printf("%#v\n", m)

	//existence testing
	wordCount, ok := m['你']
	fmt.Printf("%#v, %v\n", wordCount, ok)
	wordCount, ok = m['中']
	fmt.Printf("%#v, %v\n", wordCount, ok)

	//不在乎拿到的value是key不存在时初始化的零值, 还是key存在时value就是零值.
	wordCount = m['好']
	fmt.Printf("%#v\n", wordCount)

	//deletion
	delete(m, '中')
	fmt.Printf("%#v\n", m)

	//iteration
	for key, value := range m {
		fmt.Printf("%#v, %v\n", key, value)
	}

}
