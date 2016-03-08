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


func readFile(path string) (*trie.Trie) {
     //reads text file and prints it out 
    file, err := os.Open(path)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        return reader(file)
        
}

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
    // trie := trie.NewTrie()
    trie := readFile("data/wordsEn.txt")
  
    trie.AddWord("hello")
    trie.AddWord("heo")
    trie.AddWord("hel")
    trie.AddWord("hey")
    trie.AddWord("hell")
    trie.AddWord("he")
    

    fmt.Println(trie.FindEntries("he" , 10))
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

