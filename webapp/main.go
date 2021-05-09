package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

// タイトル名でファイルを作成する
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 読み込むための関数
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//キャッシング用
// template.ParseFiles()は可変長引数をとり、その引数としてキャッシュさせたいファイルの名前を指定
// template.Must()は内部でエラーチェックするので返り値にerrorなし
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, _ := template.ParseFiles(tmpl + ".html")
	// t.Execute(w, p)
	// ここでExecuteとExecuteTemplateが出てきた
	// テンプレート一つ Execute
	// テンプレート一つor複数 ExecuteTemplate
	// 以下の書き方だと毎回呼び出す必要がなくなる
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound) // StatusFound 302
}

// お決まりのvalidPath 正規表現で/path/foo のようなパス
var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

/*
http.HandlerFuncで必要なハンドラを作成する
viewHandler関数などを引数にとってhttp.HandlerFuncを返す関数
*/
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ここでvalidPathを使用 ([a-zA-Z0-9]+)$ のpathを取得し
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println(m)
		// /view/foo のfoo以降がなければ404NotFoundを返す
		if m == nil {
			http.NotFound(w, r)
			return
		}
		// m[2]について m = [/view/test view test] という値が入っているので、3番目を指定　
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
