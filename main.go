package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.olx.com.br/autos-e-pecas/carros-vans-e-utilitarios/estado-df?q=pulse"

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36'}"

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Host", "olx.com.br")

	// url := "https://df.olx.com.br/distrito-federal-e-regiao/autos-e-pecas/carros-vans-e-utilitarios/pulse-impetus-2022-1-0-turbo-zero-1265691610?lis=listing_2020"
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Status different of 200: StatusCode#%d", res.StatusCode))
	}

	html.Parse(res.Request.Body)

}
