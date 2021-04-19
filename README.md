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