# go_udemy

# delve Goのデバッグツール
go get -u github.com/derekparker/delve/cmd/dlv

# go doc 

# package document
fmt
%v %T
https://golang.org/pkg/fmt/ 

strings

# gofmt 整形してくれる
- gofmt -w ファイル名

# 配列とスライス　
- 配列var a [2]int
- スライス n := []int{1,2,3,4,5}
- ２次元配列 スライスの中にスライス
    var board = [][]int{
        []int{0,1,2},
        []int{3,4,5},
        []int{6,7,8}
    }
- 配列はサイズの変更ができない ×append()
- スライスのmake cap
    - どちらも同じ初期化
    - b := make([]int, 0) → 0をスライスに確保
    - var c []int → メモリに確保しない
    以下3つの挙動が違う
    - https://play.golang.org/p/3Tp_959mZCj 

# map
- pythonでいう辞書型 phpでいう php の連想配列に似ている
- 初期化 
    - m := make(map[string]int)
    - var m2 map[string]int  nilになるのでappendしないと追加できない
- v, ok := ["apple"] 
- スライスもmapもvarで宣言するとnilが初期値になる

# バイト
- b := []byte{72, 73}
- string(b) でキャスト

# 関数
- 引数の型が全て同じであれば最後に１つだけ書けば良い
    - func add(x , y int) {}
- 返り値が複数の時は 「,」 で引数とreturn
    - func add(x, y int) int, int {return x + y, x - y}
- 返り値に名前をつけることができる
    -  func calc(price, item int) (result int) {
        result = price * item
        return result
    }
    - 返り値が明らかにわかるものに関しては変数名を付けなくても良い
- 関数の中に関数を定義できる
    - func(x int) {}() のように変数に入れずに()を記入することで即時実行

# クロージャー
https://play.golang.org/p/fS3MG-QoBIr

# 可変長引数
https://play.golang.org/p/7Hc1Tli2cCT

# 