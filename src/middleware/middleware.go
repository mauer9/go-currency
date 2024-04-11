package middleware

import "test/src/service"

type Middleware func(s service.MainService) service.MainService
