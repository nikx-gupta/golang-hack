package basicHandler

import (
	"io"
	"net/http"
)

func ReportHandler(w http.ResponseWriter, req * http.Request){
	io.WriteString(w, "This is report handler function handling path" + req.URL.Path)
}

func main() {

}
