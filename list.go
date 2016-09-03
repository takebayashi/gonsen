package gonsen

import (
	"github.com/mitchellh/packer/common/json"
	"io/ioutil"
	"net/http"
)

type response struct {
	Result []string `json:"result"`
}

func GetProgramNames() ([]string, error) {
	res, err := http.Get("http://www.onsen.ag/api/shownMovie/shownMovie.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	ps := response{}
	json.Unmarshal(b, &ps)
	return ps.Result, nil
}
