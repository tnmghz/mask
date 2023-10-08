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
	fmt.Println(string(p))
}
func search(p, k []byte) {
	var n, j int
	for i := n; i < len(p); i++ {
		if p[i] != k[j] {
			j = 0
		}
		if p[i] == k[j] {
			j++
			if j == len(k) { //сравнился ли ключ полностью
				n = i + 1 //индекс, с которого начинается ссылка
				j = 0
				mask(n, p)
			}
		}
	}
}

// mask маскирует ссылку
func mask(n int, p []byte) {
	for i := n; i < len(p); i++ {
		if p[i] == s {
			break
		}
		if p[i] != s {
			p[i] = m
		}
	}
}
