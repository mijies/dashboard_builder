package dashboard

type iterator interface {
    hasNext()	bool
    next()		[]string
}

type iterable struct {
	index	int
	length	int
}
type commandsIterable struct {
	iterable
	items	*commands
}
type snippetsIterable struct {
	iterable
	items	*snippets
}

func (i *iterable) hasNext() bool {
    if i.index < i.length {
        return true
    }
    return false
}

func(i *commandsIterable) next() []string {
	// item := (*i.items)[i.index]
	// style := (*i.styles)[i.index]
	i.index++
	return []string{}
}
func(i *snippetsIterable) next() []string {
	// item := (*i.items)[i.index]
	// style := (*i.styles)[i.index]
	i.index++
	return []string{}
}
