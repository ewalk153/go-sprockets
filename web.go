package main

import (
	"fmt"
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
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>Sprockets and Go</title>
  <meta name="description" content="The HTML5 Herald">
  <meta name="author" content="SitePoint">

  <link rel="stylesheet" href="/assets/`+assets.FindAsset("application.css")+`">

  <!--[if lt IE 9]>
  <script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
  <![endif]-->
</head>

<body>
  <h1>Hello world</h1>
  <script src="/assets/`+assets.FindAsset("application.js")+`"></script>
</body>
</html>`)
	})
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 8000")
}
