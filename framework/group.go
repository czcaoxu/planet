package framework

type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	Group(string) IGroup
}

type Group struct {
	prefix      string
	c           *Core
	parentGroup *Group

	middlewares []ControllerHandler
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

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parentGroup == nil {
		return g.middlewares
	}
	return append(g.parentGroup.getMiddlewares(), g.middlewares...)
}

func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.c.Get(g.getAbsoluteUri()+uri, allHandlers...)
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.c.Post(g.getAbsoluteUri()+uri, allHandlers...)
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.c.Put(g.getAbsoluteUri()+uri, allHandlers...)
}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.c.Delete(g.getAbsoluteUri()+uri, allHandlers...)
}

func (g *Group) Group(prefix string) IGroup {
	group := NewGroup(prefix, g.c)
	group.parentGroup = g
	return group
}
