// コンパイル（特に指定がなければ、カレントディレクトリと同名のバイナリファイルが生成されます。）
// go build
// 生成されたバイナリファイルのパスを入力すれば、プログラムが実行されます。
package main // パッケージ名（必須）

import (
	"fmt" // https://golang.org/pkg/fmt/
	"log" // https://golang.org/pkg/log/
)

// （原則として）最初に実行される関数です。
func main() {
	fmt.Println("hello world") // コマンドライン出力
	log.Println("hello world") // 同時に日時も出力してくれます。
}