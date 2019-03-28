
package main

type TrieTree struct {
    word        []string
    next        [26]*TrieTree
}

func NewTrie() *TrieTree {
    return &TrieTree{word:make([]string, 0)}
}

func (t *TrieTree) Insert(s string) {
    for _, c := range s {
        cr := t.next[c-'a']
        if cr == nil {
            cr = NewTrie()
            t.next[c-'a'] = cr
        }
        cr.word = append(cr.word, s)
        t = cr
    }
}

func (t *TrieTree) Search(s string) []string {
    for _, c := range s {
        cr := t.next[c-'a']
        if cr == nil {
            return nil
        }
        t = cr
    }
    return t.word
}





