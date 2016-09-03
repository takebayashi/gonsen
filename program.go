package gonsen

import (
	"bytes"
	"github.com/mitchellh/packer/common/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Program struct {
	MediaType    string
	ThumbnailUrl string
	MediaUrl     string
	Title        string
	Slug         string
	Personality  string
	Guest        string
	Updated      time.Time
	Index        int
}

type rawProgram struct {
	Type           string
	ThumbnailPath  string
	MoviePath      map[string]string
	Title          string
	Personality    string
	Guest          string
	Update         string
	Count          string
	Schedule       string
	OptionText     string
	Mail           string
	Copyright      string
	Url            string
	Link           []map[string]string
	RecommendGoods []map[string]string
	RecommendMovie []map[string]string
	Cm             []map[string]string
	AllowExpand    string
}

func newProgram(raw rawProgram) Program {
	p := Program{}
	p.MediaType = raw.Type
	p.ThumbnailUrl = "http://www.onsen.ag" + raw.ThumbnailPath
	p.MediaUrl = raw.MoviePath["pc"]
	p.Title = raw.Title
	p.Slug = raw.Url
	p.Personality = raw.Personality
	p.Guest = raw.Guest
	p.Updated, _ = time.Parse("2006.1.2", raw.Update)
	p.Index, _ = strconv.Atoi(raw.Count)
	return p
}

func GetProgram(name string) (Program, error) {
	res, err := http.Get("http://www.onsen.ag/data/api/getMovieInfo/" + name)
	if err != nil {
		return Program{}, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Program{}, err
	}
	b = bytes.TrimPrefix(b, []byte("callback("))
	b = bytes.TrimSpace(b)
	b = bytes.TrimSuffix(b, []byte(");"))
	raw := rawProgram{}
	json.Unmarshal(b, &raw)
	return newProgram(raw), nil
}
