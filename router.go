package main

import "plat/framework"

func registerRouter(core *framework.Core) {
	core.Get("/user/login", UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
		subjectApi.Put("/:id", SubjectPutController)
		subjectApi.Delete("/:id", SubjectDelController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
