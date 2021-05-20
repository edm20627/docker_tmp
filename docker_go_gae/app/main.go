package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"
	gaeLog "google.golang.org/appengine/log"
)

func main() {
	http.HandleFunc("/", world)
	http.HandleFunc("/gae", googleAppEngine)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)

	// サービスが google.golang.org/appengine パッケージを使用している場合は、appengine.Main() の呼び出しが必要
	// https://cloud.google.com/appengine/docs/standard/go111/go-differences#writing_a_main_package
	appengine.Main()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
	}
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func googleAppEngine(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	// TODO: "Hello, Google App Engine"とINFOログで出す。
	// TODO: "Hello, Google App Engine"とwに書き込む。
	gaeLog.Infof(ctx, "Hello, Google App Engine")
	fmt.Fprintln(w, "Hello, Google App Engine")
}
