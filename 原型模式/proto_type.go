package 原型模式

import (
	"encoding/json"
)

/*
用一个比较好的例子去理解原型模式在go语言中的使用

需求: 假设现在数据库中有大量数据，包含了关键词，关键词被搜索的次数等信息，模块 A 为了业务需要
会在启动时加载这部分数据到内存中
并且需要定时更新里面的数据
同时展示给用户的数据每次必须要是相同版本的数据，不能一部分数据来自版本 1 一部分来自版本 2
*/

type Keyword struct {
	word  string
	visit int
}

func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

// Keywords 关键字 map
type Keywords map[string]*Keyword

// Clone 复制一个新的 keywords
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式
func (words Keywords) Clone(updatedWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		// 这里是浅拷贝，直接拷贝了地址
		newKeywords[k] = v
	}

	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updatedWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}
