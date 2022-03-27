package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type crClient struct {
	handRing ClientInterface
}

func (c crClient) post() {

	cr := c.handRing.ReadPostInformation()
	data, _ := json.Marshal(cr)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8090/post", "application/json", r)
	if err != nil {
		log.Println("WARNING: register fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("Posting Return：", string(data))
	}
}

func (c crClient) list() {
	resp, err := http.Get("http://localhost:8090/list")
	if err != nil {
		log.Println("WARNING: Listing fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("Listing Return：", string(data))
	}
}

func (c *crClient) delete(personId uint32) {
	url := fmt.Sprintf("http://localhost:8090/delete/%d", personId)
	log.Println("url", url)
	rep, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println("WARNING: Delete request creation fails:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(rep)
	if err != nil {
		log.Println("WARNING: Delete execution fails:", err)
		return
	}

	defer resp.Body.Close()

	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("Delete Return：", string(data))
	}
}

func main() {
	crCli := &crClient{handRing: &fakeCircleInterface{
		pId:          111,
		pName:        "tom",
		content:      "hello world",
		byTimeTall:   1.77,
		byTimeWeight: 72.0,
		byTimeAge:    23,
	},
	}
	crCli.post()
	crCli.list()
	crCli.delete(crCli.handRing.GetPId())
	crCli.list()
}
