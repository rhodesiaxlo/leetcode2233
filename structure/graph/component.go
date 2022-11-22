package graph

type Component struct {
	graph   interface{}
	iter    iIterator
	v       int
	visited []bool
	count   int
	id      []int
}

func NewComponent(g interface{}, v int) Component {
	com := Component{
		graph:   g,
		v:       v,
		visited: nil,
		count:   0,
	}
	com.Init()
	return com
}

func (c *Component) Init() {
	c.visited = make([]bool, c.v)
	c.id = make([]int, c.v)
	for i := 0; i < len(c.id); i++ {
		c.id[i] = -1
	}
}

func (c *Component) GetComponent() {
	for i := 0; i < c.v; i++ {
		if !c.visited[i] {
			c.dfs(i)
			c.count++
		}
	}
}

func (c *Component) IsConnected(i, j int) bool {
	return c.id[i] != -1 && c.id[i] == c.id[j]
}
func (c *Component) dfs(in int) {
	c.visited[in] = true
	c.id[in] = c.count
	iter := NewIterator(c.graph, in)
	for v := iter.begin(); !iter.end(); v = iter.next() {
		if !c.visited[v] {
			c.dfs(v)
		}
	}
}
