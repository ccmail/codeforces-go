package copypasta

/* 套题
https://blog.csdn.net/weixin_43914593/article/details/104108049 算法竞赛专题解析（3）：并查集
*/

// 普通并查集
// https://oi-wiki.org/ds/dsu/
// https://cp-algorithms.com/data_structures/disjoint_set_union.html
// 并查集时间复杂度证明 https://oi-wiki.org/ds/dsu-complexity/
// 模板题 https://www.luogu.com.cn/problem/P3367
// 思维转换题! https://nanti.jisuanke.com/t/43488
// https://codeforces.com/problemset/problem/292/D
func unionFind(n int) {
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	initFa(n + 1) //
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }
	same := func(x, y int) bool { return find(x) == find(y) }

	// 总是合并到代表元更大的树上
	mergeBig := func(from, to int) int {
		ff, ft := find(from), find(to)
		if ff > ft {
			ff, ft = ft, ff
		}
		fa[ff] = ft
		return ft
	}

	// 离散化版本
	faMap := map[int]int{}
	find = func(x int) int {
		if fx, ok := faMap[x]; ok && fx != x {
			faMap[x] = find(fx)
			return faMap[x]
		}
		return x
	}

	// merge，并返回是否有新的 merge
	sameMerge := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		fa[x] = y
		return true
	}

	mergeRangeTo := func(l, r, to int) { // 常用：to=r+1，这时建议用左闭右开表示区间
		//if l < 0 {
		//	l = 0
		//}
		//if r > n {
		//	r = n
		//}
		for i := find(l); i <= r; i = find(i + 1) { // initFa 需要开 n+1 空间
			fa[i] = to
		}
	}

	//rangeFullMerged := func() bool { return find(0) == n }

	// 连通分量个数
	// countRoots > 1 表示整个图不连通
	countRoots := func(st int) (cnt int) {
		for i := st; i < len(fa); i++ {
			if find(i) == i {
				cnt++
			}
		}
		return
	}

	// 所有代表元
	getRoots := func() (roots []int) {
		for i := range fa {
			if find(i) == i {
				roots = append(roots, i)
			}
		}
		return
	}

	// 连通分量
	getComps := func() (comps map[int][]int) {
		comps = map[int][]int{}
		for i := range fa {
			f := find(i)
			comps[f] = append(comps[f], i)
		}
		return
	}

	{
		rank := make([]int, n)
		merge := func(x, y int) {
			x = find(x)
			y = find(y)
			if x == y {
				return
			}
			if rank[x] < rank[y] {
				fa[x] = y
			} else {
				fa[y] = x
				if rank[x] == rank[y] {
					rank[x]++
				}
			}
		}
		_ = merge
	}

	_ = []interface{}{
		initFa, merge, same,
		mergeBig, sameMerge, mergeRangeTo, getRoots, countRoots, getComps,
	}
}

// 并查集 - 维护点权
// 维护的可以是集合的大小、最值、XOR、GCD 等
// https://codeforces.com/edu/course/2/lesson/7/1/practice/contest/289390/problem/B
func unionFindVertexWeight(n int) {
	var fa, size []int
	initFa := func(n int) {
		fa = make([]int, n)
		size = make([]int, n)
		for i := range fa {
			fa[i] = i
			size[i] = 1
		}
	}
	initFa(n + 1) //
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from != to {
			size[to] += size[from]
			fa[from] = to
		}
	}
	same := func(x, y int) bool { return find(x) == find(y) }

	_ = []interface{}{initFa, merge, same}
}

// 并查集 - 维护边权（种类）
// 简单易懂的讲解：https://www.bilibili.com/video/av68342657?p=2
// https://cp-algorithms.com/data_structures/disjoint_set_union.html#toc-tgt-11
// https://cp-algorithms.com/data_structures/disjoint_set_union.html#toc-tgt-12
// https://oi-wiki.org/ds/dsu/#_9
// 模板题 https://codeforces.com/problemset/problem/1074/D https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/D
// 种类并查集：同义词反义词 https://codeforces.com/problemset/problem/766/D
// 种类并查集：食物链 https://www.luogu.com.cn/problem/P2024
func unionFindEdgeWeight(n int) {
	const kinds = 2
	var fa, dis []int // dis 表示点到其所在集合根节点（代表元）的距离
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
		dis = make([]int, n)
	}
	initFa(n + 1) //
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			ffx := find(fa[x])
			dis[x] += dis[fa[x]]
			fa[x] = ffx
		}
		return fa[x]
	}
	delta := func(x, y int) int { return ((dis[x]-dis[y])%kinds + kinds) % kinds } // 调用前需要保证 same(x, y) 为 true
	merge := func(from, to int, d int) bool { // 返回是否与已知条件矛盾
		if fFrom, fTo := find(from), find(to); fFrom != fTo {
			dis[fFrom] = d + dis[to] - dis[from]
			fa[fFrom] = fTo
			return true
		}
		return delta(from, to) == d
	}
	same := func(x, y int) bool { return find(x) == find(y) }

	// 离散化版本
	faMap, disMap := map[int]int{}, map[int]int{}
	find = func(x int) int {
		if fx, ok := faMap[x]; ok && fx != x {
			ffx := find(fx)
			disMap[x] += disMap[fx]
			faMap[x] = ffx
			return ffx
		}
		return x
	}

	_ = []interface{}{initFa, merge, same, delta}
}

// 并查集组（一般用于涉及到位运算的题目）
// NOTE: 也可以写成 struct 的形式
func multiUnionFind(n, m int) {
	fas := make([][]int, m)
	for i := range fas {
		fas[i] = make([]int, n) // n+1
		for j := range fas[i] {
			fas[i][j] = j
		}
	}
	var find func([]int, int) int
	find = func(fa []int, x int) int {
		if fa[x] != x {
			fa[x] = find(fa, fa[x])
		}
		return fa[x]
	}
	merge := func(fa []int, from, to int) { fa[find(fa, from)] = find(fa, to) }
	same := func(fa []int, x, y int) bool { return find(fa, x) == find(fa, y) }
	mergeRange := func(fa []int, l, r int) { // 左闭右开
		for i := find(fa, l); i < r; i = find(fa, i+1) {
			fa[i] = r // merge 到 r 上
		}
	}

	_ = []interface{}{merge, same, mergeRange}
}

// 可持久化并查集
// todo
