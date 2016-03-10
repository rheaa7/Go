package main 

import (
    "os"
    "fmt"
    "net/http"
    "log"
    "bufio"
    "github.com/rheaa7/Go/trie"
    "io"
    "encoding/json"
    "strconv"
    "runtime"
    "strings"
)

//create new trie
var tri = trie.NewTrie()
var req bool

//construct search function
type SearchResponse struct {
    Word []string `json:"word"`
}

//open and read file 
func readFile(path string) {
     //reads text file and prints it out 
    file, err := os.Open(path)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        tri = reader(file)
        req = true
        
}

//initilize reader to read contents of file and monitors memory 
//usage as the trie loads, and stop loading when your memory usage 
//approaches 1GB.
func reader(file io.Reader) (*trie.Trie) {
    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)
    var memstats = new(runtime.MemStats)
        for scanner.Scan() {
            runtime.ReadMemStats(memstats)
                line := strings.TrimSpace(scanner.Text())
            var limit = 850000000
            if memstats.Alloc < uint64(limit) {
               tri.AddWord(line)
            } else {
                break
            }
        }
    return tri
}


//takes in searched query and suggests words
func search(w http.ResponseWriter, r *http.Request) {
    if (req) {
        word := r.URL.Query().Get("search")
        max := r.URL.Query().Get("max")
        newMax, _ := strconv.Atoi(max)
        
        entries := tri.FindEntries(word, newMax)
        resp := SearchResponse{Word:entries}
      
            //convert struct to JSON
            j, err := json.Marshal(resp)
            if nil != err {
                log.Println(err)
                w.WriteHeader(500)
                w.Write([]byte(err.Error()))
            } else {
                //tell the client we are sending back JSON
                if (len(entries) == 0) {
                    w.WriteHeader(400)
                }
                w.Header().Add("Content-Type", "application/json")
                w.Write(j)
            }
     } else {
         //returns error if file still loading
         w.WriteHeader(400)
         w.Write([]byte("error not loading"))
     }
}



func main()  {
    req = false
    
    //calls go routine
    go readFile(os.Args[1]) 
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/api/v1/suggestions/", search)
    
    //listen and server on port 9000
    fmt.Println("Server listening on port 9000...")
    http.ListenAndServe(":9000", nil)
    
}