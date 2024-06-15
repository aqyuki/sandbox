package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()

	tmp := strings.Split(scn.Text(), " ")
	n, _ := strconv.Atoi(tmp[0]) // 家の数
	d, _ := strconv.Atoi(tmp[1]) // 家との距離がこの値以下の場合　尋ねる

	scn.Scan()
	tmp = strings.Split(scn.Text(), " ")
	x, _ := strconv.Atoi(tmp[0]) // 家の座標
	y, _ := strconv.Atoi(tmp[1]) // 家の座標

	x_addr := make([]int, n)
	y_addr := make([]int, n)

	// 家の座標を読み込む
	for i := 0; i < n; i++ {
		scn.Scan()
		tmp = strings.Split(scn.Text(), " ")
		t_x, _ := strconv.Atoi(tmp[0])
		t_y, _ := strconv.Atoi(tmp[1])

		x_addr[i] = t_x
		y_addr[i] = t_y
	}

	// 距離を計算しカウントする
	var count int

	for i := 0; i < n; i++ {
		dis_x := x - x_addr[i]
		dis_y := y - y_addr[i]

		if dis_x < 0 {
			dis_x = -1 * dis_x
		}
		if dis_y < 0 {
			dis_y = -1 * dis_y
		}

		if dis_x+dis_y <= d {
			count++
		}
	}
	fmt.Println(count)
}
