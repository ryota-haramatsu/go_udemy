package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 単純なGETリクエスト
	// resp, _ := http.Get("https://bondlingo.tv")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// パース処理
	// base, _ := url.Parse("https://example.com")
	// fmt.Println(base)

	// クエリデータあり
	base, _ := url.Parse("https://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	// fmt.Println(endpoint)

	// リクエスト
	req, _ := http.NewRequest("GET", endpoint, nil) // この状態ではリクエストは投げていない
	req.Header.Add("If-None-Match", `W/wzyyy`)      // ヘッダー情報の追加
	q := req.URL.Query()                            // map[a:[1] b:[2]]のようにマップでクエリの値
	q.Add("c", "3")                                 // クエリを追加
	// fmt.Println(q)
	req.URL.RawQuery = q.Encode() // 元に戻すときはエンコード

	// アクセスするときはクライアントを作成
	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req) // client.Do()のなかにリクエスト情報を入れてアクセスする
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
