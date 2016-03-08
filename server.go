package main 

import (
    "os"
    "fmt"
    "net/http"
    "log"
    "bufio"
    "github.com/rheaa7/Go/trie"
)

func main()  {
     trie := trie.NewTrie()
     
     
    //reads text file and prints it out 
    file, err := os.Open("data/wordsEn.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            fmt.Println(scanner.Text()) //print contents of file
            //trie.AddWord(scanner.Text())
        }

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    
    trie.AddWord("hello")
    trie.AddWord("hel")
    trie.AddWord("hey")
    trie.AddWord("hell")
    trie.AddWord("he")
    
    trie.FindEntries("he" , 10);
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

