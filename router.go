package main

import "plat/framework"

func registerRouter(core *framework.Core) {
	core.Get("fool", FoolControllerHandler)
}
