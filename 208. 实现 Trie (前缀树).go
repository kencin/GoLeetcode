package main

import "fmt"

/*
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type Trie struct{
	Root *Node
}

type Node struct{
	Child map[rune]*Node
	Char rune
	isDone bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		newNode(' '),
	}
}

func newNode(char rune) *Node{
	return &Node{
		Char:char,
		Child:make(map[rune]*Node),
		isDone:false,
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	Root := this.Root
	Chars := []rune(word)
	for _,c:=range Chars{
		_,hasNode:=Root.Child[c]
		if !hasNode{   //如果这个字母不存在
			//构建一个新Child
			node := newNode(c)
			Root.Child[c] = node
		}
		// 父节点替换
		Root = Root.Child[c]
	}
	// 最后设置标志位以示这是一个完整数据
	Root.isDone = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	Chars := []rune(word)
	Root := this.Root
	for _,c:=range Chars{
		node,hasNode := Root.Child[c]
		if !hasNode{
			return false
		}
		Root = node
	}
	if Root.isDone{
		return true
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	Chars := []rune(prefix)
	Root := this.Root
	for _,c:=range Chars{
		node,hasNode := Root.Child[c]
		if !hasNode{
			return false
		}
		Root = node
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

 func main(){
 	trie := Constructor()
 	trie.Insert("apple")
 	fmt.Println(trie.Search("app"))
	 fmt.Print(trie.StartsWith("app"))
 }