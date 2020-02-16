package copypasta

import "sort"

// 下面这些都是最小堆
// h.top() 即 h.IntSlice[0] 或 (*h)[0] （注意判断非空）

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}
//func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆

//

type hp64 []int64 // 自定义类型

func (h hp64) Len() int              { return len(h) }
func (h hp64) Less(i, j int) bool    { return h[i] < h[j] } // > 为最大堆
func (h hp64) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp64) Push(v interface{})   { *h = append(*h, v.(int64)) }
func (h *hp64) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

//

// see graph.shortestPathDijkstra
type hPair struct {
	x int64
	y int
}
type pairHeap []hPair

func (h pairHeap) Len() int              { return len(h) }
func (h pairHeap) Less(i, j int) bool    { return h[i].x < h[j].x || h[i].x == h[j].x && h[i].y < h[j].y }
func (h pairHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(v interface{})   { *h = append(*h, v.(hPair)) }
func (h *pairHeap) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }
