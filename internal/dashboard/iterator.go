package dashboard

type iterator interface {
    hasNext()	bool
    next()		[][]string
}

type iter struct {
	index	int
	length	int
	comp	dashboard_component
}

func (i *iter) hasNext() bool {
    if i.index < i.length {
        return true
    }
    return false
}

func(i *iter) next() [][]string {
	item := i.comp.intoRow(i.index)
	i.index += 1
	return item
}
