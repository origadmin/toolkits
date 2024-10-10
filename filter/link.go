package filter

type Link struct {
	Name string
	Next []*Link
}

func (l *Link) Equal(path string) bool {
	return l.Name == path
}

func (l *Link) Has(name string) bool {
	for i := range l.Next {
		if l.Next[i].Name == name {
			return true
		}
	}
	return false
}

func (l *Link) Contains(paths []string) bool {
	if len(paths) == 0 {
		return false
	}
	if !l.Equal(paths[0]) {
		return false
	}
	next := paths[1:]
	if len(next) == 0 {
		return true
	}
	if len(l.Next) == 0 {
		return false
	}
	var link *Link
	for i := range l.Next {
		link = l.Next[i]
		if link.Contains(paths[1:]) {
			return true
		}
	}
	return false
}

func (l *Link) Add(next *Link) bool {
	for i := range l.Next {
		if l.Next[i].Name == next.Name {
			return false
		}
	}
	l.Next = append(l.Next, next)
	return true
}

func (l *Link) String() string {
	return l.Name
}
