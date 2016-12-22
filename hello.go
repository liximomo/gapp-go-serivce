package hello

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/wangbin/jiebago"
)

var seg jiebago.Segmenter

func wordSegment(sentence string) []string {
	result := make([]string, 0)
	useHmm := true
	ch := seg.Cut(sentence, useHmm)
	for word := range ch {
		result = append(result, word)
	}
	return result
}

func init() {
	// 加载默认词典
	err := seg.LoadDictionary("./dict/default/dict.txt")
	if nil != err {
		log.Printf("load dict %s", err)
	}

	// 加载额外词典
	err = seg.LoadUserDictionary("./dict/extra/stop_words.txt")
	if nil != err {
		log.Printf("load dict %s", err)
	}

	http.HandleFunc("/ws", wordSegmentHandle)
}

func wordSegmentHandle(w http.ResponseWriter, req *http.Request) {
	/* words := []int{1, 2, 3} */
	query := req.URL.Query()
	srcText, exist := query["text"]
	if !exist {
		srcText = []string{""}
	}

	words := wordSegment(srcText[0])

	jsonString, err := json.Marshal(words)
	log.Printf("words %v", jsonString)
	if err != nil {
		// log.Printf("transform to json fail %s", err)
		jsonString = []byte("[]")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}
