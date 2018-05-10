package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Help() {
	fmt.Println("请依次输入变量：P H T ZTD")
}

func PWV(Two_res, ZWD_res float64) {
	var PWV_res float64
	PWV_res = Two_res * ZWD_res
	fmt.Println("PWV :", PWV_res)
}

func Two(T float64) float64 {
	var Two_res float64
	Two_res = 1000000 / (461000 * ((378000 / T) + 16.48))
	fmt.Println("Two :", Two_res)
	return Two_res
}

func ZWD(ZTD, ZHD_res float64) float64 {
	var ZWD_res float64
	ZWD_res = ZTD - ZHD_res
	fmt.Println("ZWD :", ZWD_res)
	return ZWD_res
}

func ZHD(P, H float64) float64 {
	var ZHD_res float64
	ZHD_res = 22.7 * P / (0.775 - 0.00028*H)
	fmt.Println("ZHD: ", ZHD_res)
	return ZHD_res
}

func Handler(P, H, T, ZTD float64) {

	PWV(Two(T), ZWD(ZTD, ZHD(P, H)))
}

func main() {
LABEL:
	Help()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("依次输入：P H T ZTD > ")
		b, _, _ := r.ReadLine()
		line := string(b)

		tokens := strings.Split(line, " ")

		// var (
		// 	P, H, T, ZTD float32
		// )
		if len(tokens) == 4 {
			P, err := strconv.ParseFloat(tokens[0], 10)
			H, err := strconv.ParseFloat(tokens[1], 10)
			T, err := strconv.ParseFloat(tokens[2], 10)
			ZTD, err := strconv.ParseFloat(tokens[3], 10)
			if err != nil {
				goto LABEL
			}

			Handler(P, H, T, ZTD)
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
