package main

import "fmt"

func pertanyaan(o int) {
	var pertanyaan [5]string
	pertanyaan[0] = "Saya merasa mampu mengatasi tekanan dalam kehidupan sehari-hari : "
	pertanyaan[1] = "Saya sering merasa cemas atau khawatir tanpa alasan yang jelas : "
	pertanyaan[2] = "Saya merasa memiliki seseorang untuk diajak bicara ketika saya sedang tertekan : "
	pertanyaan[3] = "Saya mampu tidur dengan nyenyak dan cukup setiap malam : "
	pertanyaan[4] = "Saya merasa puas dengan kehidupan saya saat ini : "
	fmt.Printf("%s\n", pertanyaan[o])
}
