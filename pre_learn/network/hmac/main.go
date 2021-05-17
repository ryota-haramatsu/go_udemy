package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

/*
HMAC認証 (Hash-based Message Authentication Code)
HMAC-SHA256
- APIアクセスで認証時にヘッダーに含める時に使用
*/

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))   // 1. ハッシュの生成
	h.Write(data)                                  // 2. ハッシュにメッセージを追加 (メッセージはbyteスライス)
	expectedHMAC := hex.EncodeToString(h.Sum(nil)) // 3. ハッシュのhにnilを足して、hexにエンコード
	fmt.Println(sign == expectedHMAC)              // 送られてきたsignとサーバーサイドで作ったHMACが一致するか
}

func main() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	// sha256という暗号化方式でapiSecretをbyteにしてハッシュを作成する
	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret)) // 1. ハッシュの生成
	h.Write(data)                                // 2. ハッシュにメッセージを追加 (メッセージはbyteスライス)
	sign := hex.EncodeToString(h.Sum(nil))       // 3. ハッシュのhにnilを足して、hexにエンコード

	// 以上のsignをサーバーサイドに投げてユーザーが正しいかチェックすることができる
	Server(apiKey, sign, data) // 3つの情報をサーバーに投げ、情報が正しければtrue
	fmt.Println(sign)
}
