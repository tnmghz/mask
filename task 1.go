// принимает на вход строку (сообщение) и маскирует там ВСЕ ссылки, заменяя их на звездочки
package main

import "fmt"

// toSlice1 преобразует полученную строку в срез
func toSlice1(page string, s []byte) {
	for i, x := range page {
		s[i] = byte(x)
	}
	fmt.Println("длина введенной строки:", len(s))
	fmt.Println(string(s))
}

// toSlice2 преобразует полученный ключ в срез
func toSlice2(key string, k []byte) {
	for i, x := range key {
		k[i] = byte(x)
	}
}

// mask маскирует ссылку
func mask(s, k []byte) {
	var n int
	for i, j := 0, 0; i < len(s); {
		if s[i] == k[j] {
			i++
			j++
			if j == len(k)-1 { //сравнился ли ключ полностью
				n = i + 1 //индекс, с которого начинается ссылка
				j = 0
				for i := n; i < len(s); i++ {
					if s[i] != ' ' { //маскирует ссылку до пробела или конца строки
						s[i] = '*'
					} else {
						break
					}
				}
			}
		} else { //следующий элемент не соответствует
			i++
			j = 0
		}
	}
	fmt.Println("длина полученной строки:", len(s))
	fmt.Println(string(s))
}

func main() {
	const page = "Here's http://hehefouls.netHAHAHA my spammy page: http://hehefouls.netHAHAHA see you.http://hehefouls.netHAHAHA"
	const key = "http://"
	s := make([]byte, len(page))
	k := make([]byte, len(key))
	toSlice1(page, s)
	toSlice2(key, k)
	mask(s, k)
}
