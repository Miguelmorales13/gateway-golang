package main

import (
	db "./db/conexion"
	"./db/entities"
	"./db/models"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
)

func gategayFunc(w http.ResponseWriter, r *http.Request) {
	url := r.Header.Get("Url")
	fmt.Println(r.Method, url, nil)
	sendType(w, r)
}

func main() {
	err := godotenv.Load()
	if err != nil { log.Fatal("Error loading .env file") }
	my := db.DbConn()
	defer my.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", gategayFunc)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

func sendType(w http.ResponseWriter, r *http.Request) {
	urlComplete := r.Header.Get("Url")
	url,errorsote:=url2.Parse(urlComplete)
	if errorsote!=nil || url.Host == "" {
		http.Error(w, "Url in existent", 500)
		return
	}
	scheme:=url.Scheme
	host:=url.Host
	println(url.Scheme,url.Host,url.Path,url.RequestURI())
	proxySecret := r.Header.Get("Proxy-Authorization")
	id, auth := ValidateHttps(scheme+"://"+host, proxySecret)
	log.Println(id, auth)
	if !auth {
		http.Error(w, "Unauthorized", 500)
		return
	}
	// creation request
	req, err := http.NewRequest(r.Method, urlComplete, r.Body)
	if err != nil {
		// if error	send error
		fmt.Fprintf(w, "Error to request"+err.Error())
		return
	}
	lenHeadersRequest := cloneHeaders(req, r)
	// creation agent sending request
	client := &http.Client{}
	// send request
	response, errorResponse := client.Do(req)
	if errorResponse != nil {
		http.Error(w, errorResponse.Error(), 500)
		return
	}

	lenHeadersResponse := copyHeadersResponse(w.Header(), response.Header)
	w.WriteHeader(response.StatusCode)
	// copy response to service client and return data sizing in binary
	body, errors := io.Copy(w, response.Body)
	if errors != nil {
		http.Error(w,"Error to copy body request", 500)
		return
	}
	// save data request
	// creation object to save Report
	var log = &entities.Report{
		IpClient:       r.RemoteAddr,
		HttpMethod:     r.Method,
		HttpMimeType:   r.Header.Get("Content-Type"),
		HttpStatusCode: strconv.Itoa(response.StatusCode),
		HttpUrl:        urlComplete,
		HttpReplySize:  int(body) + lenHeadersRequest + lenHeadersResponse,
		AddressID:      strconv.FormatUint(uint64(id), 10),
		TimeResponse:   10,
		ProxyStatus:    "DONE",
		CarrierID:      os.Getenv("CARRIER_ID"),
	}
	// save report
	go models.Create(log)




}

func ValidateHttps(host string, proxySecret string) (uint, bool) {
	if proxySecret == "" {
		return 0, false
	}
	address := models.FindByHost(host, "", proxySecret)
	println(address.ID)
	if address.ID == 0 {
		return 0, false
	}
	if address.Active == false {
		return 0, false
	}
	return address.ID, true
}
func cloneHeaders(from *http.Request, to *http.Request) (counter int) {
	counter = 0
	for k, vs := range to.Header {
		for _, v := range vs {
			// log.Println("Request", k, v)
			if k != "Accept-Encoding" {
				counter = counter + len(k) + len(v)
				from.Header.Add(string(k), string(v))
			}
		}
	}
	return
}
func copyHeadersResponse(dst, src http.Header) (counter int) {
	counter = 0
	for k, vs := range src {
		for _, v := range vs {
			// log.Println("Response", k, v)
				counter = counter + len(k) + len(v)
				dst.Add(k, v)

		}
	}
	return
}
