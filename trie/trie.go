package trie

import  (
    "fmt"
)

type Trie struct {
    Root *Node
    Count int
}

type Node struct {
    Letter rune
    Children map[rune]*Node
    WordEnd bool
    
}

func NewTrie() *Trie {
    node := &Node{Children: make(map[rune]*Node)}
    root := Trie{Root: node}
    return &root
}

func (trie *Trie) AddWord(entry string) {
    curr := trie.Root
    for i, letter := range entry {
        // fmt.Println(curr)
        if curr.Children[letter] == nil {
            var nodeMap = make(map[rune] *Node)
            var node = Node{Letter: letter, Children: nodeMap}
            
            curr.Children[letter] = &node
            //fmt.Println(curr.Children[letter])
            
            if i == len(entry) - 1 {
                fmt.Println("what")
                trie.Count++
                curr.Children[letter].WordEnd = true
                fmt.Println(curr);
            } else {
                curr.Children[letter].WordEnd = false
            }
            curr = curr.Children[letter]
        }
    }
    // fmt.Println(trie.Root)

}

func (trie *Trie) FindEntries(prefix string, max uint8) string {
    fmt.Println(trie.Count)
    currentNode := trie.Root
    currentNode = make(map[rune] *Node)
    for i := range prefix {
        letter := string(i)
        currentNode = currentNode.Children(letter)
    }
    // return FindEntriesHelper(max, uint8, word string, currentNode node) 
    //return "hello"
}

func FindEntriesHelper(max uint8, word string, currentNode *Node) {
    var listOfWords []string
    if (max == 0) {
        return []listOfWords
    } 
    if (currentNode.WordEnd) {
        listOfWords = append.([]listOfWords, word)
        fmt.Println([]listOfWords)
    }
    if (currentNode.Children == nil) {
        return []listOfWords
    } else {
        for i := range currentNode.Children {
            currentNode = currentNode.Children[i]
            return append([]listOfWords, FindEntriesHelper(max, (word + currentNode.child), currentNode.Children))
        }
    }
}
