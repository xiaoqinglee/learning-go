package supplement

import (
	"bufio"
	"fmt"
	"net/http"
)

func HttpClient() {
	//resp, err := http.Get("https://music.163.com")
	//resp, err := http.Get("http://localhost:9000/hello")
	resp, err := http.Get("http://localhost:9000/cancelable-hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("response status: %s\n", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 15; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err() != nil; err {
		panic(err)
	}
}
