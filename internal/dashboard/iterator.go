package dashboard

type iterator interface {
    hasNext()	bool
    next()		[]string
}

type iter struct {
	index	int
	length	int
	items	*[][]string
}

func (i *iter) hasNext() bool {
    if i.index < i.length {
        return true
    }
    return false
}

func(i *iter) next() []string {
	item := (*i.items)[i.index]
	i.index++
	return item
}
