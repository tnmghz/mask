// принимает на вход строку (сообщение) и маскирует там ВСЕ ссылки, заменяя их на звездочки
package main

import "fmt"

const page = "Here's my spammy page: http://hehefouls.netHAHAHA see you."
const key = "http://"
const m, s = '*', ' '

func main() {
	fmt.Println(page)
	p := []byte(page)
	k := []byte(key)
	search(p, k)
}

// mask маскирует ссылку
func search(p, k []byte) {
	var n, j int
	for i := 0; i < len(p); i++ {
		if p[i] == k[j] {
			j++
			if j == len(k)-1 { //сравнился ли ключ полностью
				n = i + 1 //индекс, с которого начинается ссылка
				j = 0
				mask(n, p)
			}
		} else { //следующий элемент не соответствует
			j = 0
		}
	}
	fmt.Println(string(p))
}
func mask(n int, p []byte) {
	for i := n; i < len(p); i++ {
		if p[i] != s { //маскирует ссылку до пробела или конца строки
			p[i] = m
		} else {
			break
		}
	}
}
