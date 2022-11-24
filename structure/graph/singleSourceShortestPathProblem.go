package graph

// 单源最短路径问题

// 我们之前在图中的层序遍历的方法是在无权图中给出了从A 点到任意其它点的最短路径
// 实际上是 single source forest

// 在有权图中，情况有所不同，A - B 的举例，不一定比 A-C-B 的距离短， relaxation 操作

// 比较典型的是  kijkstra 算法， 局限性，不能有负权边  算法复杂度 Elog(v)
