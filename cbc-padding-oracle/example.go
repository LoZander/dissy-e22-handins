package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func printBytes(bytes []byte) {
	for _, b := range bytes {
		fmt.Printf("%d ", b)
	}
	fmt.Println()
}

func testSystemsSecurity(baseURL string) {
	res, _ := http.Get(baseURL + "/")
	ciphertextHex := res.Cookies()[0].Value
	fmt.Printf("[+] received ciphertext: %s\n", ciphertextHex)

	ciphertextBytes, _ := hex.DecodeString(ciphertextHex)
	ciphertextBytes[len(ciphertextBytes)-1] = 0x01
	printBytes(ciphertextBytes)

	/*send a request with a modified ciphertext*/

	req, _ := http.NewRequest(http.MethodGet, baseURL+"/check/", nil)
	/*If we modify the ciphertext we get a response like: invaliud type if.  normal is plain ciphertextHex */

	req.AddCookie(&http.Cookie{Name: "authtoken", Value: (hex.EncodeToString(ciphertextBytes))})
	res, _ = http.DefaultClient.Do(req)
	resBody, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("[+] done:\n%s\n", resBody)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run example.go <base url>")
		return
	}
	testSystemsSecurity(os.Args[1])
}
