package main
// １つのディレクトリには１つのパッケージしか作れませんが、
// 同じパッケージであれば、ファイルをいくつでも作れます。
// 要するに、パッケージとディレクトリは、１対１対応になります。
import (
	// 標準パッケージ以外のパッケージをインポートする場合には、
	// /go/src/を起点とする相対パスを指定します。
	"app/greet"
)

var name string = "Docker"
// グローバル変数の宣言方法です。
// name := "Docker"
// という宣言方法（リテラル）もあり、このように記載すると、型（string）の指定を省略できます。

func usePackageGreet() {
	greet.HelloWithName(name)
}