package main

func main() {
	testSearch()

}

func testSearch() {

}

/*
func testAgentLoad() bool {
	if !Put(SloResource, SloUrl, "") {
		return false
	}
	return Put(TimeseriesResourceV2, Timeseries2Url, "")
}

func testAgentAddSLO(id, controller, threshold string) bool {
	entries := []slo.EntryV1{{
		Id:          id,
		Controller:  controller,
		Threshold:   threshold,
		StatusCodes: "0",
	},
	}
	buf, err := json.Marshal(entries)
	if err != nil {
		fmt.Printf("error: AddSLO() -> %v", err)
		return false
	}
	r := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, SloUrl, io.NopCloser(r))
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return false
	}
	resp, _ := exchange.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
	return true
}

func Put(file, uri, variant string) bool {
	//u, _ := url.Parse(file)
	buf, status := io2.ReadFile(file) //io2.ReadFile(u)
	if !status.OK() {
		fmt.Printf("read file err: %v\n", status.Error())
		return false
	}
	reader := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, uri, reader)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return false
	}
	resp, status := exchange.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
	fmt.Printf("Put() [status:%v]\n", status)
	return true
}

func Delete(uri, variant string) {
	req, err1 := http.NewRequest(http.MethodDelete, uri, nil)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return
	}
	resp, _ := exchange.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
}

func testSearch() {
	uri := "http://localhost:8081/github/advanced-go/search/provider:search?q=golang"
	//req, _ := http.NewRequest(http.MethodGet, "http://localhost:8081/github/advanced-go/search/provider:search?q=golang")

	h := make(http.Header)
	//h.Add(runtime.AcceptEncoding, "gzip, deflate, br")
	resp, status := exchange.Get(nil, uri, h)
	if !status.OK() {
		fmt.Printf("error on Get(): %v\n", status)
		return
	}

	buf, status1 := io2.ReadAll(resp.Body, nil)
	if !status1.OK() {
		fmt.Printf("error on ReadAll(): %v\n", status1)
		return
	}
	err := os.WriteFile("C:\\users\\markb\\github\\example-test\\pkg\\resource\\g-golang-4.html", buf, 066)
	if err != nil {
		fmt.Printf("error on WriteFile(): %v\n", err)
		return
	}

}


*/
