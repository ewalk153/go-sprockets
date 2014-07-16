package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// usage: launch rack app with rackup
// then run this application
func main() {
	remote, err := url.Parse("http://localhost:9292")
	if err != nil {
		panic(err)
	}

	manifest, err := http.Get(remote.String() + "/manifest.json")

	// TODO - reload this manifest when changes occur, or add dev mode where we send thru
	// original asset name
	assets, err := NewAssetManager(manifest.Body)
	fmt.Println("Assets", assets)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.Handle("/assets/", proxy)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.New("base").Funcs(template.FuncMap{"assets": assets.FindAsset})
		tmpl, err = tmpl.ParseFiles("./tmpl/index.html")
		if err != nil {
			panic(err)
		}
		err = tmpl.ExecuteTemplate(w, "index.html", struct{}{})
		if err != nil {
			fmt.Println(err)
		}
	})
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 8000")
}
