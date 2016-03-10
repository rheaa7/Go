package trie_test

import (
	"testing"
	"fmt"
    "github.com/rheaa7/Go/trie"
)

//testing add entry method
func TestAddEntry(t *testing.T) {
	fmt.Println("Testing add entry to trie")
    var tri = trie.NewTrie()
	tri.AddWord("accepts")
    tri.AddWord("abet")
    tri.AddWord("aaron")
    tri.AddWord("ahoy")
    tri.AddWord("hell")
    tri.AddWord("hey")
    
	fmt.Println(tri.Root)
}

//testing find entries method
func TestFindEntries(t * testing.T) {
    fmt.Println("Testing find entries")
    var tri = trie.NewTrie()
	tri.AddWord("accepts")
    tri.AddWord("abet")
    tri.AddWord("aaron")
    tri.AddWord("ahoy")
    tri.AddWord("hell")
    tri.AddWord("hey")
    
    list := (tri.FindEntries("a" , 10))
    correctList := make([]string, 4)
    correctList[0] = "accepts"
    correctList[1] = "abet"
    correctList[2] = "aaron"
    correctList[3] = "ahoy"
    
    if (list [0] == correctList [0] || list [0] == correctList [1] || list [0] == correctList [2] || list [0] == correctList [3]) {
        fmt.Println("all entries have been found")
    } else if (list [1] == correctList [0] || list [1] == correctList [1] || list [1] == correctList [2] || list [1] == correctList [3]) {
        fmt.Println("all entries have been found")
    } else if (list [2] == correctList [0] || list [2] == correctList [1] || list [2] == correctList [2] || list [2] == correctList [3]) {
        fmt.Println("all entries have been found")
    } else if (list [3] == correctList [0] || list [3] == correctList [1] || list [3] == correctList [2] || list [3] == correctList [3]) {
       fmt.Println("all entries have been found")
    } else {
        t.Errorf("not all entries found, compare your strings below")
    }
   
    fmt.Println(correctList)
    fmt.Println(list)

}

//testing count of tries
func TestCount(t * testing.T) {
    var tri = trie.NewTrie()
    tri.AddWord("accepts")
    tri.AddWord("abet")
    tri.AddWord("aaron")
    tri.AddWord("ahoy")
    tri.AddWord("hell")
    tri.AddWord("hey")
    count := (tri.FindEntries("a" , 10))
    
    if (len(count) == 4) {
        fmt.Println("passed count test")
    } else {
        t.Errorf("did not pass count test")
    }
}