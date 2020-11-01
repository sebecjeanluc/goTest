// パッケージ名を指定
package main

// formatをimport
import "fmt"

// interfaceを定義
type Car interface {
	Drive()
	Stop()
}

// 名前をコンストラクト、stringで指定
type Tesla struct {
	TeslaModel string
}

type Toyota struct {
	ToyotaModel string
}

func (te *Tesla) Drive() {
	fmt.Println("Tesla S is ignated")
	fmt.Println(te.TeslaModel)
}

func (to *Toyota) Drive() {
	fmt.Println("Toyota is on move")
	fmt.Println(to.ToyotaModel)
}

func main() {
	te := Tesla{"S3"}
	to := Toyota{"Alphard"}
	to.Drive()
	te.Drive()
}
