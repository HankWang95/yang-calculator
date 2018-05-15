package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Help() {
	fmt.Println("请依次输入变量：Ps Ts ZTD")
}

func Hd(Ts float64) float64 {
	var hd_res float64
	hd_res = 40136 + 148.72*(Ts-273.16)
	return hd_res
}
func P(Ps float64) float64 {
	var p_res float64
	p_res = (10000 + Ps) * 0.1
	return p_res
}

func ZHD(Ts, hd_res, p_res float64) float64 {
	var ZHD_res float64
	ZHD_res = (77.6 * 0.000001 * p_res * (hd_res - 181.418)) / (5 * Ts)
	fmt.Println("ZHD:", ZHD_res)
	return ZHD_res
}

func ZWD(ZTD, ZHD_res float64) float64 {
	var ZWD_res float64
	ZWD_res = ZTD - ZHD_res
	fmt.Println("ZWD:", ZWD_res)
	return ZWD_res
}

func PWV(ZWD_res float64) {
	var PWV_res float64
	PWV_res = 0.15 * ZWD_res
	fmt.Println("PWV:", PWV_res)
}

func Handler(Ps, Ts, ZTD float64) {
	PWV(ZWD(ZTD, ZHD(Ts, Hd(Ts), P(Ps))))
}

func main() {
LABEL:
	Help()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("依次输入：Ps Ts ZTD > ")
		b, _, _ := r.ReadLine()
		line := string(b)

		tokens := strings.Split(line, " ")

		if len(tokens) == 3 {
			Ps, err := strconv.ParseFloat(tokens[0], 10)
			Ts, err := strconv.ParseFloat(tokens[1], 10)
			ZTD, err := strconv.ParseFloat(tokens[2], 10)
			if err != nil {
				goto LABEL
			}

			Handler(Ps, Ts, ZTD)
			goto LABEL
		}

	}
}
