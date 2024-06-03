package route

//import (
//	"bytes"
//	"io"
//	"net/http"
//	"testing"
//)
//
//func TestSetUpRouter(t *testing.T) {
//	client := http.Client{}
//
//	type test struct {
//		method   string
//		body     []byte
//		endpoint string
//		response string
//	}
//
//	tests := []test{
//		{
//			method:   "GET",
//			endpoint: "localhost:8080/positions",
//			response: "[]",
//		},
//	}
//
//	for _, test := range tests {
//		req, err := http.NewRequest(test.method, test.endpoint, bytes.NewBuffer(test.body))
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		res, err := client.Do(req)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		defer res.Body.Close()
//		body, err := io.ReadAll(res.Body)
//		if err != nil
//	}
//}
