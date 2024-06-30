package main

import (
	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New("user> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		result := rep(line)
		println(result)
	}
}

func rep(in string) string {
	return _print(eval(read(in)))
}

func read(in string) string {
	return in
}

func eval(in string) string {
	return in
}

func _print(in string) string {
	return in
}
