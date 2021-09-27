package simplemux

import (
	"fmt"
	"math/rand"
	"net/http"
)

type ManualServerMux struct {
}

func (p *ManualServerMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	giveRandom(w, req)
	return
}

func CreateMultiMux() *http.ServeMux {
	var multiMux = http.NewServeMux()
	multiMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	multiMux.HandleFunc("/randomInt", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, rand.Int31())
	})

	return multiMux
}

func giveRandom(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main() {
}
