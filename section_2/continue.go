package main

import "fmt"

func main() {
	fmt.Println("vim.go")

	flag := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		}
		fmt.Println(i)
	}
	fmt.Println(flag)

	// switch

	day := "FRI"
	switch day {
	case "MON":
		fmt.Println("Monday")
	case "SAT", "SUN":
		fmt.Println("Weekend!")
	case "FRI":
		fmt.Println("TGIF!")
	default:
		fmt.Println("default")
	}

	switch {
	case day == "FRI":
		fmt.Println("TGIF")
		break
	}

	err := error.New("sadusahpdu saudha")
	switch err {
	case err:
		fmt.Println("It's error!")
	}
}
