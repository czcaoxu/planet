package framework

type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	Group(string) IGroup
}

type Group struct {
	prefix      string
	c           *Core
	parentGroup *Group
}

func NewGroup(prefix string, core *Core) *Group {
	group := &Group{
		prefix: prefix,
		c:      core,
	}

	return group
}

func (g *Group) getAbsoluteUri() string {
	if g.parentGroup == nil {
		return g.prefix
	}

	return g.parentGroup.getAbsoluteUri() + g.prefix
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	g.c.Get(g.getAbsoluteUri()+uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	g.c.Post(g.getAbsoluteUri()+uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	g.c.Put(g.getAbsoluteUri()+uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	g.c.Delete(g.getAbsoluteUri()+uri, handler)
}

func (g *Group) Group(prefix string) IGroup {
	group := NewGroup(prefix, g.c)
	group.parentGroup = g
	return group
}
