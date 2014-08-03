package handlers 

import (
	"net/http"
	"html/template"
	"log"
	"conf"	
	"path/filepath"
	mgr "book/manager"
)

type IndexHandler struct {
	
}

func (self *IndexHandler)IndexAction(w http.ResponseWriter, r *http.Request) {
	indexDir := filepath.Join(conf.TEMPLATE_DIR, "template/html/index.html")
	t, _ := template.ParseFiles(indexDir)

	bookInfos, err := mgr.Manager.GetFileInfos()
	log.Printf("error:", bookInfos)
	if err != nil {
		log.Printf("error: ", err)
	}
	t.Execute(w, nil)
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

var Index IndexHandler = *NewIndexHandler()
