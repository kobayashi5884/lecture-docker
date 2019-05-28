package greet
// パッケージとディレクトリは、特に理由がなければ、同名にします。

import (
	"fmt"
)

// インポート先で利用する変数や関数は頭文字を大文字にします。
// パッケージ内部でのみ利用する変数や関数は、頭文字を小文字にします。
func HelloWithName(name string) {
	fmt.Println("hello", name)
}