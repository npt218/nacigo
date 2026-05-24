package collect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"errors"
	// "os"

	"github.com/npt218/nacigo/src/dbg"
)

var pre_url string
var prefix string

type MyJSON struct {
	ID      	int    	`json:"id"`
	Link    	string 	`json:"link"`
	Rendered	string
}

func init() {
	pre_url = ""
	prefix = "out/out_"

}

func Hello() {
	fmt.Printf("Collect start %s\n", pre_url)
}

func UnmarshalJSON(data []byte) ([]MyJSON, error) {
	var aux []struct {
		ID      int    `json:"id"`
		Link    string `json:"link"`
		Content struct {
			Rendered string `json:"rendered"`
		} `json:"content"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return nil, err
	}

	var result []MyJSON
	for _, item := range aux {
		result = append(result, MyJSON{
			ID:       item.ID,
			Link:     item.Link,
			Rendered: item.Content.Rendered,
		})
	}

	return result, nil
}

func Do(number int) ([]MyJSON, error, error) {
	var page []MyJSON

	url := fmt.Sprintf(pre_url+"%d", number)

	resp, err := http.Get(url)

	if err != nil {
		return nil, nil, err
	}
	
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, err
	}

	page, err = UnmarshalJSON(body)

	if err != nil {
		var invalid struct {
			Code string `json:"code"`
		}
		if err := json.Unmarshal(body, &invalid); err != nil {
			return nil, nil, err
		}
		return nil, errors.New("Finish!"), err
	}

	return page, nil, nil

}

func All() {
	var total []MyJSON
	page := 0
	i := 0
	id := 0

	for {
		page += 1
		Res, invalid, err := Do(page)

		fmt.Printf("\rCrawl page: %d", page)
		if err != nil {
			if invalid != nil {
				fmt.Println()
				break
			}
			fmt.Println()
			msg := fmt.Sprintf("Failed to collect page: %d\n", page)
			dbg.E(msg, err)
			break
		}
		total = append(total, Res...)
		
		if i == 10 {
			//build data of 10 pages to a json
			data, err := json.Marshal(total)
			if err != nil {
				dbg.E("Build Json", err)
			}
			
			//write json of 10 pages to a file
			output := fmt.Sprintf(prefix + "%d.json", id)
			fmt.Println("\nWrite file ", output)

			if err := os.WriteFile(output, data, 0755); err != nil {
				dbg.E("Write file failed", err)
			}

			//reset count and json data for new file; increase file id
			i = 0
			total = nil
			id += 1
		} else {
			i += 1
		}
		
	}

}
