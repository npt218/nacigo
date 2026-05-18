package collect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/npt218/nacigo/src/dbg"
)

var url string
var posts []Posts_Json

type Posts_Json struct {
	ID int;
	Link string
}


func init() {
	url = ""
}

func Hello() {
	fmt.Printf("Collect start %s\n", url)
}

func Do() {
	resp, err := http.Get(url)

	if err != nil {
		dbg.E("Get error!", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		dbg.E("Read error!", err)
		return
	}

	json_err := json.Unmarshal(body, &posts)

	if json_err != nil {
		dbg.E("Decode Json error!", json_err)
		return
	}

	fmt.Printf("Res: %+v\n", posts)

}
