package nfa

import "container/list"

type state struct {
	c     string
	edge1 *state
	edge2 *state
}

type Frag struct {
	start *state
	out   *state
}

func Post2nfa(postfix list.List) *Frag {

	fragstack := []*Frag{}

	for e := postfix.Front(); e != nil; e = e.Next() {
		switch e.Value {
		case ".":
			frag2 := fragstack[len(fragstack)-1]
			fragstack = fragstack[:len(fragstack)-1]
			frag1 := fragstack[len(fragstack)-1]
			fragstack = fragstack[:len(fragstack)-1]
			frag1.out.edge1 = frag2.start
			fragstack = append(fragstack, &Frag{start: frag1.start, out: frag2.out})
		case '|':
			frag2 := fragstack[len(fragstack)-1]
			fragstack = fragstack[:len(fragstack)-1]
			frag1 := fragstack[len(fragstack)-1]
			fragstack = fragstack[:len(fragstack)-1]
			accept := state{}
			initial := state{edge1: frag1.start, edge2: frag2.start}
			frag1.out.edge1 = &accept
			frag2.out.edge1 = &accept
			fragstack = append(fragstack, &Frag{start: &initial, out: &accept})
		case '*':
			fragment := fragstack[len(fragstack)-1]
			fragstack = fragstack[:len(fragstack)-1]
			accept := state{}
			initial := state{edge1: fragment.start, edge2: &accept}
			fragment.out.edge1 = fragment.start
			fragment.out.edge2 = &accept
			fragstack = append(fragstack, &Frag{start: &initial, out: &accept})
		case '+':
			fragment := fragstack[len(fragstack)-1]
			fragstack = fragstack[:len(fragstack)-1]
			accept := state{}
			initial := state{edge1: fragment.start, edge2: &accept}
			fragment.start.edge1 = &initial
			fragstack = append(fragstack, &Frag{start: fragment.start, out: &accept})
		default:
			accept := state{}
			initial := state{c: e.Value.(string), edge1: &accept}
			fragstack = append(fragstack, &Frag{start: &initial, out: &accept})
		}
	}
	return fragstack[0]
}
