// パッケージとディレクトリは、特に理由がなければ、同名にします。
package human

import (
	"fmt"
)

// 自分で型を定義することができます。
type Human struct {
	Name string
	Age uint
}
// 自分で型を定義したら、その型の変数にのみ適用する関数（メソッド）を定義することができます。
func (h Human) Greet() {
	fmt.Println("Hello, My name is", h.Name)
	fmt.Println("I am", h.Age, "years old")
}
// インポート先で利用する変数や関数は頭文字を大文字にします。
// パッケージ内部でのみ利用する変数や関数は、頭文字を小文字にします。