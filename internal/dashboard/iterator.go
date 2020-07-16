package dashboard

type iterator interface {
    hasNext()	bool
    next()		([]string, []string)
}

type iter struct {
	index	int
	length	int
	values	*[][]string
	styles	*[][]string
}

func (i *iter) hasNext() bool {
    if i.index < i.length {
        return true
    }
    return false
}

func(i *iter) next() ([]string, []string) {
	value := (*i.values)[i.index]
	style := (*i.styles)[i.index]
	i.index++
	return value, style
}
