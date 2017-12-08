// Test of groupcache useage.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/groupcache"
)

var (
	cache *groupcache.Group
)

type falseDB struct {
	data map[string]string
}

type request struct {
	Key   string `json:"key"`
	Value string `json:"value, omitempty"`
}

func NewDB() *falseDB {
	falseDB := new(falseDB)
	falseDB.data = make(map[string]string)
	return falseDB
}

func (f *falseDB) GetValue(key string) string {
	fmt.Printf("Get %s's value from data source\n", key)
	return f.data[key]
}

func (f *falseDB) SetValue(key string, value string) {
	fmt.Printf("Set %s's value\n", key)
	f.data[key] = value
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func newResponse(res, msg string) map[string]string {
	return map[string]string{"result": res, "message": msg}
}

func Get(w http.ResponseWriter, r *http.Request) {
	// if err1 := cache.Get(nil, "key1", groupcache.AllocatingByteSliceSink(&retrievedData)); err1 == nil {
	// 	fmt.Printf("Get value of key1: %s\n", retrievedData)
	// }
	// if err2 := cache.Get(nil, "key1", groupcache.AllocatingByteSliceSink(&retrievedData)); err2 == nil {
	// 	fmt.Printf("Get value of key1: %s\n", retrievedData)
	// }
	// if err3 := cache.Get(nil, "key2", groupcache.AllocatingByteSliceSink(&retrievedData)); err3 == nil {
	// 	fmt.Printf("Get value of key2: %s\n", retrievedData)
	// }
	// if err4 := cache.Get(nil, "key2", groupcache.AllocatingByteSliceSink(&retrievedData)); err4 == nil {
	// 	fmt.Printf("Get value of key2: %s\n", retrievedData)
	// }
	if r.Method != http.MethodPost { // First check request method
		WriteJSON(w, http.StatusMethodNotAllowed, newResponse("failed", "request method not allowed"))
		return
	}

	body, _ := ioutil.ReadAll(r.Body) // Second get bytes array
	var req request
	if err := json.Unmarshal(body, &req); err != nil { // Third check it can convert to local type successfully.
		WriteJSON(w, http.StatusBadRequest, newResponse("failed", err.Error()))
		return
	}

	if req.Key == "" { // Last check if the key is empty.
		WriteJSON(w, http.StatusBadRequest, newResponse("failed", "key should not be empty."))
		return
	}

	var resultValue []byte
	var err1 error
	if err1 := cache.Get(nil, req.Key, groupcache.AllocatingByteSliceSink(&resultValue)); err1 == nil {
		fmt.Printf("Get value of key %s as %s\n", req.Key, resultValue)
		WriteJSON(w, http.StatusOK, newResponse("success", string(resultValue)))
		return
	}
	log.Fatal("Cannot get value for key %s", req.Key)
	WriteJSON(w, http.StatusBadRequest, newResponse("failed", err1.Error()))
	return
}

func main() {
	testDB := NewDB()
	testDB.SetValue("key1", "value1")
	testDB.SetValue("key2", "value2")

	cache = groupcache.NewGroup("DBCache", 64<<22, groupcache.GetterFunc( // Setup group cache.
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error { // Call this func if call all peers' cache getter func and no matched.
			value := testDB.GetValue(key)
			dest.SetBytes([]byte(value)) // Put new item in cache.
			return nil
		}))
	peerListenPort := ":8081"
	serviceHost := "http://localhost"
	peers := groupcache.NewHTTPPool(serviceHost + peerListenPort) // The listener port
	peers.Set("http://localhost:8081", "http://localhost:8082")   // Setup peers.
	// The 8081 cache would be the first to handle the get request. If no matched key, pass the request to its peer caches to check.
	// go log.Panic(http.ListenAndServe(":8081", http.HandlerFunc(peers.ServeHTTP))) // can't have log in go routine.
	go http.ListenAndServe(":8081", http.HandlerFunc(peers.ServeHTTP))
	peerPort, _ := strconv.Atoi(strings.Split(peerListenPort, ":")[1])
	servicePort := peerPort + 1000
	serviceListenPort := ":" + strconv.Itoa(servicePort)
	log.Infof("Service is listening at port %s\n", serviceListenPort)
	http.HandleFunc("/get", Get)
	log.Panic(http.ListenAndServe(serviceListenPort, nil))
}
