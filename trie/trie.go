package trie

import  (
    "fmt"
)

type Trie struct {
    Root *Node
    Count int
}

type Node struct {
    Letter string
    Children map[string]*Node
    WordEnd bool
    
}

func NewTrie() *Trie {
    node := &Node{Children: make(map[string]*Node)}
    root := Trie{Root: node}
    return &root
}

func (trie *Trie) AddWord(entry string) {
    curr := trie.Root

    index := 0
    // for i := 0; i < len(entry); i++ {
    for _, i := range entry {
        index++
        letter := string(i)
        if (curr.Children[letter] == nil) {
            curr.Children[letter] = &Node{Letter: letter, Children: make(map[string]*Node), WordEnd: false}
        }
        if (index == len(entry)) {
            trie.Count++
            curr.Children[letter].WordEnd = true
        }
        curr = curr.Children[letter]
    }
    // fmt.Println(curr)
    // index := 0
    // for _, i := range entry {
    //     letter := string(i)
    //     index++
    //     length := len(entry)
    //     // fmt.Println(curr)
    //     if curr.Children[letter] == nil {
    //         if (index == length) {
    //             var nodeMap = make(map[string] *Node)
    //             curr.Children[letter] = &Node{Letter: letter, Children: nodeMap, WordEnd: true}
    //         } else {
    //             var nodeMap = make(map[string] *Node)
    //             curr.Children[letter] = &Node{Letter: letter, Children: nodeMap, WordEnd: false}
    //         }
             
    //         // curr.Children[letter] = &node
    //         // //fmt.Println(curr.Children[letter])
    //     }
    //     if (index == length) {
    //         curr.Children[letter].WordEnd = true
    //     }
        
       
    //     // curr = curr.Children[letter]
    //     //     if i == len(entry) - 1 {
    //     //         fmt.Println("what")
    //     //         trie.Count++
    //     //         curr.Children[letter].WordEnd = true
    //     //         fmt.Println(curr);
    //     //     } else {
    //     //         curr.Children[letter].WordEnd = false
    //     //     }
    //     //     curr = curr.Children[letter]
    //     // }
    // }
    // fmt.Println(trie.Root)

}

func (trie *Trie) FindEntries(prefix string, max int) []string {
    if (prefix == "") {
        return make([]string, 0)
    }
    
    // fmt.Println(trie.Count)
    currentNode := trie.Root
    fmt.Println(currentNode)
    for _, i := range prefix {
        fmt.Println(currentNode)
        if (currentNode != nil) {
            letter := string(i)
            currentNode = currentNode.Children[letter]
        } else {
            return make([]string, 0)
        }
    }
    
    fmt.Println(currentNode)
    
    listOfWords := trie.FindEntriesHelper(max, prefix, currentNode)
    return listOfWords 
    //return "hello"
}

func (trie *Trie) FindEntriesHelper(max int, word string, currentNode *Node) []string {
    // fmt.Println(currentNode)
    listOfWords := make([]string, 0)
    if (max == 0) {
        return listOfWords
    } 
    if (currentNode.WordEnd) {
        listOfWords = append(listOfWords, word)
        max--
        // fmt.Println(listOfWords)
    }
    if (currentNode.Children == nil) {
        return listOfWords
    } else {
        for letter, Children := range currentNode.Children {
            currentNode = Children
            list := trie.FindEntriesHelper(max, word + letter, currentNode)
            if (len(list) > 0) {
                max = max - len(list)
            }
            // fmt.Println(listOfWords)
            // listOfWords = append(listOfWords, trie.FindEntriesHelper(max, (word + currentNode.Children), currentNode.Children)...)
            listOfWords = append(listOfWords, list...)
        }
        return listOfWords
    }
}
