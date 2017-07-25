package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Tobby112/go-practice/06_download_mp3/mp3"
)

func main() {
	fmt.Printf("Start")
	mp3.TestPrint()
	resp, err := http.Get("http://www.icrt.com.tw/info_list01.php?&mlevel1=6&mlevel2=15&mlevel3=1")
	if err == nil {
		fmt.Println(resp)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			bodyStr := string(body)
			fmt.Println(bodyStr)
			ok := strings.Contains(bodyStr, "eznewsaudiodownload_am.php")
			fmt.Printf("found eznewsaudiodownload_am.php : %b\n", ok)
		}
	} else {
		fmt.Println(err)
	}
}
