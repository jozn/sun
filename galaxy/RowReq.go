package galaxy

import (
	"fmt"
	"ms/sun/config"
	"ms/sun/helper"
	"net/url"
	"os"
	"strings"
)

type rowReq struct {
	fileName                  string
	fileDataStoreId           int //64bit
	fileExtensionWithoutDot   string
	fileFormat                fileFormat
	requestedImageSize        int
	rowCacheOutFullPathDir    string
	rowCacheOutFullPath       string
	rowCacheOutFullPathThumb  string
	hasError                  bool
	isLocalDiskCacheAvailable bool
	cacheFullModuleDirectory  string // "/web/galxy/chat/"
}

func newRowReq(url *url.URL) (row *rowReq, err error) {
	urlPath := strings.Trim(url.Path, "/")
	sep := strings.Split(urlPath, "/")
	fileName := sep[len(sep)-1]
	id, size, ext, err := nameToParts(fileName)
	if err != nil {
		return
	}
	row = &rowReq{
		fileName:                 fileName,
		fileDataStoreId:          id,
		fileExtensionWithoutDot:  ext,
		requestedImageSize:       size,
		cacheFullModuleDirectory: config.GALAXY_CACHE_PARENT_DIR + "cache/",
	}
	err = row.setOutputChaseFullPath()
	if err != nil {
		return
	}
	return
}

//1254_24.jpg -> 1254 24 "jpg" nil  ------ 1254.jpg -> 1254 0 "jpg"
//erors : 125 125_   156
//todo add size support
func nameToParts(name string) (id int, size int, ext string, err error) {
	sep := strings.Split(name, "/")
	fileName := sep[len(sep)-1]
	sep2 := strings.Split(fileName, ".")
	fmt.Println("6666666", name, "9999999")
	if len(sep2) == 2 && len(sep2[0]) > 0 && len(sep2[1]) > 0 {
		id = helper.StrToInt(sep2[0], 0)
		ext = sep2[1] //it dosent have any dot
		if id > 0 {
			return
		}
	}
	err = errWrongFileName
	return
}

func (r *rowReq) setOutputChaseFullPath() (err error) {
	ids := r.fileName
	r.rowCacheOutFullPathDir = fmt.Sprintf("%s%c/%c/%c/%c/%c/%c/", r.cacheFullModuleDirectory,
		rune(ids[0]), rune(ids[1]), rune(ids[2]), rune(ids[3]), rune(ids[4]), rune(ids[5]))
	r.rowCacheOutFullPath = fmt.Sprintf("%s%d.%s", r.rowCacheOutFullPathDir, r.fileDataStoreId, r.fileExtensionWithoutDot)
	r.rowCacheOutFullPathThumb = fmt.Sprintf("%s%d_thumb.%s", r.rowCacheOutFullPathDir, r.fileDataStoreId, r.fileExtensionWithoutDot)

	return
}

func (r *rowReq) createRowOutCacheDir() {
	os.MkdirAll(r.rowCacheOutFullPathDir, os.ModeDir)
}

///dep
func fileIdToChachePath(id int, parentChachDir string) (filePath string, err error) {
	ids := fmt.Sprint(id)
	if len(ids) < 6 {
		err = errFileNameTooShort
		return

	}
	filePath = fmt.Sprintf("%s%d%d%d%d%d%d.%s", parentChachDir, ids[0], ids[1], ids[2], ids[3], ids[4], ids[5])
	return
}
