package main

var tsAPI *APITesting

func main(){
	
// START OMIT
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "404") {
			w.WriteHeader(404)
			return
		}
		if strings.HasSuffix(r.URL.Path, "jsonarray") {
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintln(w, "[42]")
			return
		}
	}
	
	tsAPI = NewClient(ts.URL, new(http.Client)) // HL
	
// END OMIT
}	
