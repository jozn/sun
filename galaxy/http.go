package galaxy

import (
	"fmt"
	"ms/sun/config"
	"net/http"
	"os"
	"strings"
)

type httpHandler struct {
}

func (httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	os.Stdout = nil
	row, err := newRowReq(r.URL)
	fmt.Println(row)
	fmt.Println(err)
	if err == nil {
		http.ServeFile(w, r, row.rowCacheOutFullPath)
	} else {
		http.NotFound(w, r)
	}
}

func (httpHandler) ServeHTTP1(w http.ResponseWriter, r *http.Request) {
	os.Stdout = nil
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[1:])
	fmt.Println(strings.Split(r.URL.Path, "/"))
	fmt.Println(os.Getwd())
	//panic("ADS")
	sep := strings.Split(r.URL.Path, "/")
	fileName := sep[len(sep)-1]
	path := config.GALAXY_CACHE_PARENT_DIR + "/ups/" + fileName
	//http.ServeFile(w,r,"ups/1.png")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("============================")
	}
	fmt.Println(path)
	http.ServeFile(w, r, path)
}
