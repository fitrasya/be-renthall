package routelist

import (
	api "be-renthall/api"
	"github.com/labstack/echo/v4"
)

func ApiRoute(echo *echo.Echo) {
	echo.GET("/pemesan", api.GetPemesan)
	echo.GET("/pemesan/:id", api.GetPemesanByIdorKontak)
	echo.POST("/pemesan", api.PostPemesan)
	echo.PUT("/pemesan", api.PutPemesan)
	echo.DELETE("/pemesan", api.DelPemesan)

	echo.GET("/gedung", api.GetGedung)
	echo.GET("/gedung/:id", api.GetGedungById)
	echo.POST("/gedung", api.PostGedung)
	echo.PUT("/gedung", api.PutGedung)
	echo.DELETE("/gedung", api.DelGedung)

	echo.GET("/pesanan", api.GetPesanan)
	echo.GET("/pesanan/user/:id", api.GetPesananByIdUser)
	echo.GET("/pesanan/:id", api.GetPesananById)
	echo.POST("/pesanan", api.PostPesanan)
	echo.PUT("/pesanan", api.PutPesanan)
	echo.DELETE("/pesanan", api.DelPesanan)
	echo.PUT("/status-pesanan", api.PutPesananStatus)

	echo.GET("/dashboard", api.Dashboard)
	echo.GET("/schedule", api.Schedule)
}
