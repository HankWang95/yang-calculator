package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Help() {
	fmt.Println("请依次输入变量：Ps ZTD")
}

func PWV(ZWD_res float64) {
	var PWV_res float64
	PWV_res = 0.15 * ZWD_res
	fmt.Println("PWV :", PWV_res)
}

func ZWD(ZTD, ZHD_res float64) float64 {
	var ZWD_res float64
	ZWD_res = ZTD - ZHD_res
	fmt.Println("ZWD :", ZWD_res)
	return ZWD_res
}

func ZHD(Ps float64) float64 {
	var ZHD_res float64
	ZHD_res = 2.2768*(10000+Ps)*0.1/1 - (0.0026*-0.024 - 0.00028*0.181418)
	fmt.Println("ZHD: ", ZHD_res)
	return ZHD_res
}

func Handler(Ps, ZTD float64) {
	PWV(ZWD(ZTD, ZHD(Ps)))
}

func main() {
LABEL:
	Help()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("依次输入：Ps ZTD > ")
		b, _, _ := r.ReadLine()
		line := string(b)

		tokens := strings.Split(line, " ")

		// var (
		// 	P, H, T, ZTD float32
		// )
		if len(tokens) == 2 {
			Ps, err := strconv.ParseFloat(tokens[0], 10)
			ZTD, err := strconv.ParseFloat(tokens[1], 10)
			if err != nil {
				goto LABEL
			}

			Handler(Ps, ZTD)
			goto LABEL
		}

		// if handler, ok := handlers[tokens[0]]; ok {
		// 	ret := handler(tokens)
		// 	if ret != 0 {
		// 		break
		// 	}
		// } else {
		// 	fmt.Println("Unknown command:", tokens[0])
		// }
	}
}
