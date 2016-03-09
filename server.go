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

var tri = trie.NewTrie()

type SearchResponse struct {
    Word []string `json:"word"`
}

//open and read file 
func readFile(path string) (*trie.Trie) {
     //reads text file and prints it out 
    file, err := os.Open(path)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        return reader(file)
        
}

//initilize reader to read contents of file
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
            // tri.AddWord(scanner.Text())
            // fmt.Println(scanner.Text()) //print contents of file
            
            //    tri.AddWord("accepts")
            //     tri.AddWord("abet")
            //     tri.AddWord("aaron")
            //     tri.AddWord("ahoy")
            //     tri.AddWord("hell")
            //     tri.AddWord("he")
        }
        return tri
}

func search(w http.ResponseWriter, r *http.Request) {
    word := r.URL.Query().Get("search")

    max := r.URL.Query().Get("max")
    newMax, _ := strconv.Atoi(max)
    entries := tri.FindEntries(word, newMax)
    fmt.Println("Word")
    fmt.Println(word)
    fmt.Println("Entires")
    fmt.Println(entries)
    fmt.Println("Max")
    fmt.Println(newMax)
    
    resp := SearchResponse{Word:entries}
      
       //convert struct to JSON
        j, err := json.Marshal(resp)
        if nil != err {
            log.Println(err)
            w.WriteHeader(500)
            w.Write([]byte(err.Error()))
        } else {
            //tell the client we are sending back JSON
            w.Header().Add("Content-Type", "application/json")
            w.Write(j)
        }
}



func main()  {
    
    //accepts command line argument that specifies location of data file to load into trie
    argsWithoutProg := os.Args[1:]
     if (len(argsWithoutProg) == 0 ) {
         tri = readFile("data/wordsEn.txt")
     } else {
         tri = readFile(os.Args[1])
    }
  
    //testing add words
    // trie.AddWord("accepts")
    // trie.AddWord("abet")
    // trie.AddWord("aaron")
    // trie.AddWord("ahoy")
    // trie.AddWord("hell")
    // trie.AddWord("he")
    

    // fmt.Println(trie.FindEntries("a" , 10))
    
    
      http.Handle("/", http.FileServer(http.Dir("./static")))
      http.HandleFunc("/api/v1/suggestions/", search)
    
    //listen and server on port 9000
    fmt.Println("Server listening on port 9000...")
    http.ListenAndServe(":9000", nil)

}






//    dir:="./data/"
//    if len(os.Args) > 1 {
//         if _, err := os.Stat(os.Args[1]); err == nil {
//             dir = os.Args[1]
//             if dir[len(dir)-1:] != "/" {
//                 dir += "/"
//             }
//         }
//     }

//set things to lowercase in add word 
//