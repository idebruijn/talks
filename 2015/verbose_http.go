client := new(http.Client)
payload := strings.NewReader("key=value")
req , err := http.NewRequest("POST",
	"http://localhost:12346/v1/properties",
	payload)
if err != nil {
	log.Fatal("invalid url:%v",err)
}
req.Header.Add("Content-Type","text/plain")
req.Header.Add("Accept","application/json")
resp, err := client.Do(req)
if err != nil {
	log.Fatal("failed to send request:%v",err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	log.Fatal("failed to read response:%v",err)
}
if resp.StatusCode != http.StatusOK {
	var serviceError ServiceError
	err = json.Unmarshal(body, &serviceError )
}
...
