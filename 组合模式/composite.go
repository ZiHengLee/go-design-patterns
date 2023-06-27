package 组合模式

type Component interface {
	countBytes() int
}

type File struct {
	name string
	text string
}

func (f *File) countBytes() int {
	return len(f.text)
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) countBytes() int {
	ret := 0
	for _, composite := range f.components {
		ret += composite.countBytes()
	}
	return ret
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
