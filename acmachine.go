package acmachine

type node struct {
	children map[interface{}]*node
	fail     *node
	stop     bool
}

func newNode() *node {
	return &node{
		children: map[interface{}]*node{},
	}
}

type MatchResult struct {
	Pattern    interface{}
	StartIndex int
	EndIndex   int
}

type Machine interface {
	AddPattern(pattern interface{})
	Build()
	Match(target interface{}) []MatchResult
}

type machine struct {
	root    *node
	split   func(interface{}) []interface{}
	combine func([]interface{}) interface{}
}

func NewMachine(split func(interface{}) []interface{}, combine func([]interface{}) interface{}) Machine {
	return &machine{
		root:    newNode(),
		split:   split,
		combine: combine,
	}
}

func (m *machine) AddPattern(pattern interface{}) {
	pl := m.split(pattern)
	root := m.root
	for _, p := range pl {
		if _, ok := root.children[p]; !ok {
			root.children[p] = newNode()
		}
		root = root.children[p]
	}
	root.stop = true
}

func (m *machine) Build() {
	q := NewQueue(1024)
	q.Push(m.root)
	for !q.Empty() {
		v, _ := q.Pop()
		curr := v.(*node)
		for c, child := range curr.children {
			fail := curr.fail
			for fail != nil {
				if _, ok := fail.children[c]; ok {
					child.fail = fail.children[c]
					break
				}
				fail = fail.fail
			}
			if fail == nil {
				child.fail = m.root
			}
			q.Push(child)
		}
	}
}
func (m *machine) Match(target interface{}) []MatchResult {
	pl := m.split(target)
	curr := m.root
	rets := make([]MatchResult, 0, 8)
	var start int
	for i := 0; i < len(pl); i++ {
		for curr != nil {
			if curr == m.root {
				start = i
			}
			if _, ok := curr.children[pl[i]]; ok {
				curr = curr.children[pl[i]]
				break
			}
			curr = curr.fail
		}
		if curr != nil && curr.stop {
			rets = append(rets, MatchResult{EndIndex: i, StartIndex: start, Pattern: m.combine(pl[start : i+1])})
		}
		if curr == nil {
			curr = m.root
		}
	}
	return rets
}

func SplitString(v interface{}) []interface{} {
	str, ok := v.(string)
	if !ok {
		return nil
	}
	rets := make([]interface{}, 0, len(str))
	for _, r := range str {
		rets = append(rets, r)
	}
	return rets
}

func CombineString(vs []interface{}) interface{} {
	rs := make([]rune, 0, len(vs))
	for _, v := range vs {
		rs = append(rs, v.(rune))
	}
	return string(rs)
}
