package trie

import "strings"

// Size is the number of possible characters in the trie ranging from (a-z)
const Size = 26

type TreeMap interface {
	Insert(string, interface{})
	Get(string) interface{}
	Exists(string) bool
	Prefix(string) bool
}

func NewTreeMap() TreeMap {
	return &trie{root: new(node)}
}

type node struct {
	// single letter
	char  rune
	value interface{}
	// array of nodes with each child having 26 children
	children [Size]*node
	// indicates if the current node is at the end of a word
	terminal bool
}

type trie struct {
	root *node
}

// Insert a new value in the TreeMap at the given key word
func (t *trie) Insert(key string, value interface{}) {
	// track the current node for traversing the tree starting with the root
	current := t.root
	lower := strings.ToLower(key)
	for _, character := range lower {
		// mapping of the character to index within the english alphabet
		// by taking the decimal representation of the character from the ascii table.
		// for example 'a' - 'a' = index 0, 'b' - 'a' = index 1, etc.
		charIdx := character - 'a'
		if current.children[charIdx] == nil {
			current.children[charIdx] = &node{char: character, value: value}
		}
		current = current.children[charIdx]
	}
	current.terminal = true
}

// Get returns the value stored at the given key.
func (t *trie) Get(key string) interface{} {
	current := t.root
	lower := strings.ToLower(key)
	for _, character := range lower {
		charIdx := character - 'a'
		if current == nil || current.children[charIdx] == nil {
			return nil
		}
		current = current.children[charIdx]
	}
	return current.value
}

// Exists returns true if the entire word exists in the TreeMap
func (t *trie) Exists(word string) bool {
	current := t.root
	lower := strings.ToLower(word)
	for _, character := range lower {
		charIdx := character - 'a'
		if current == nil || current.children[charIdx] == nil {
			return false
		}
		current = current.children[charIdx]
	}
	return current.terminal
}

// Prefix returns true if there is a word in the TreeMap that starts with a given prefix
func (t *trie) Prefix(prefix string) bool {
	current := t.root
	lower := strings.ToLower(prefix)
	for _, character := range lower {
		charIdx := character - 'a'
		if current == nil || current.children[charIdx] == nil {
			return false
		}
		current = current.children[charIdx]
	}
	return true
}

type byteNode struct {
	children [16]*byteNode
	data     []byte
}

func (n *byteNode) copy() *byteNode {
	out := byteNode{
		children: n.children,
		data:     make([]byte, len(n.data)),
	}
	copy(out.data, n.data)
	return &out
}

func lookup(n *byteNode, k []byte) []byte {
	if n == nil {
		return nil
	} else if len(k) == 0 {
		return n.data
	} else {
		return lookup(n.children[k[0]], k[1:])
	}
}

func insert(n *byteNode, k []byte, v []byte) *byteNode {
	if n == nil {
		out := byteNode{}
		return insert(&out, k, v)
	} else if len(k) == 0 {
		out := n.copy()
		out.data = v
		return out
	} else {
		out := n.copy()
		out.children[k[0]] = insert(n.children[k[0]], k[1:], v)
		return out
	}
}

//type KVStore interface {
//	// This inserts the new (key, value) pair into
//	// the trie. k must contain values in [0, 15].
//	Insert(k, b []byte) Trie
//
//	// Search the trie for the value associated with
//	// the given key. k must contain values in [0, 15].
//	// nil is returned if the key is not found.
//	Lookup(k []byte) []byte
//}

func (n *byteNode) Lookup(k []byte) []byte {
	return lookup(n, k)
}

func (n *byteNode) Insert(k, v []byte) *byteNode {
	return insert(n, k, v)
}

func New() *byteNode {
	return nil
}
