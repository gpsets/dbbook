package fetch

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicln("Get url: ", url, "error: ", err)
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")

	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		log.Panicln("Get url: ", url, "error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	e := DetermineEncoding(resp.Body)
	utf8Body := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Body)
}

func DetermineEncoding(body io.Reader) encoding.Encoding {
	rBody, err := bufio.NewReader(body).Peek(1024)
	if err != nil {
		log.Println("DetermineEncoding err: ", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(rBody, "")
	return e
}
