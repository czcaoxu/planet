package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := make(map[string]*Tree, 4)
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

func (c *Core) Get(uri string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(uri string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(uri string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(uri string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(response, request)

	handlers := c.FindRouteByRequest(request)
	if handlers == nil {
		log.Println("cannot find router")
		return
	}
	ctx.SetHandlers(handlers)

	if err := ctx.Next(); err != nil {
		ctx.Json(500, "Internal error")
	}
}

func (c *Core) FindRouteByRequest(requeest *http.Request) []ControllerHandler {
	uri := requeest.URL.Path
	method := requeest.Method

	if methodHandler, ok := c.router[method]; ok {
		return methodHandler.FindHandler(uri)
	}

	return nil
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(prefix, c)
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}
