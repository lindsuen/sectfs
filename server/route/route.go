// sectfs - route.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package route

import (
	"github.com/labstack/echo/v4"
	h "github.com/lindsuen/sectfs/server/handler"
)

const routePrefix = "/file"

// LoadEchoRoutes can load routes of Echo.
func LoadEchoRoutes(r *echo.Echo) {
	r.GET("/", h.GetRoot)

	r.GET(routePrefix+"/", h.GetRoot2)

	r.POST(routePrefix+"/upload", h.UploadFile)

	r.GET(routePrefix+"/download", h.DownloadFile)
}
