package hackersguild

import (
    "html/template"
    "net/http"
)

func init() {
    http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  var indexTempl = template.Must(template.ParseFiles("HTML/index.html"))
  w.Header().Set("Content-Type", "text/html")
  indexTempl.Execute(w, nil)
}
