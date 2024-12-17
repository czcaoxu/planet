package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]*Tree
}

func NewCore() *Core {
	router := make(map[string]*Tree, 4)
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

func (c *Core) Get(uri string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(uri, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(uri string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(uri, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(uri string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(uri, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(uri string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(uri, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(response, request)

	handler := c.FindRouteByRequest(request)
	if handler == nil {
		log.Println("cannot find router")
		return
	}

	handler(ctx)
}

func (c *Core) FindRouteByRequest(requeest *http.Request) ControllerHandler {
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
