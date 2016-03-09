package main 

import (
    "os"
    "fmt"
    "net/http"
    "log"
    "bufio"
    "github.com/rheaa7/Go/trie"
    "io"
)

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
    trie := trie.NewTrie()
    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)
    
        for scanner.Scan() {
            fmt.Println(scanner.Text()) //print contents of file
            // trie.AddWord(scanner.Text())
        }
        return trie
}


func main()  {
    
    //accepts command line argument that specifies location of data file to load into trie
    argsWithoutProg := os.Args[1:]
     trie := trie.NewTrie()
     if (len(argsWithoutProg) == 0 ) {
         trie = readFile("data/wordsEn.txt")
     } else {
         trie = readFile(os.Args[1])
    }
  
    //testing add words
    trie.AddWord("hello")
    trie.AddWord("heo")
    trie.AddWord("hel")
    trie.AddWord("hey")
    trie.AddWord("hell")
    trie.AddWord("he")
    

    fmt.Println(trie.FindEntries("he" , 10))
    
    http.Handle("/", http.FileServer(http.Dir("./static")))
    
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