package routelist

import (
	"be-renthall/db"
	"be-renthall/helper"
	"be-renthall/model"
	"fmt"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func GetPemesan(c echo.Context) error {
	dbm := db.Manager().Debug()
	pemesan := []model.Pemesan{}
	if err := dbm.Find(&pemesan).Error; err != nil {
		return helper.ToResponse(c, err, nil)
	}

	return helper.ToResponse(c, nil, pemesan)
}

func GetPemesanByIdorKontak(c echo.Context) error {
	id := c.Param("id")
	dbm := db.Manager().Debug()
	pemesan := model.Pemesan{}
	if err := dbm.Where("id = ? or kontak = ?", id, id).Take(&pemesan).Error; err != nil {
		return helper.ToResponse(c, err, nil)
	}

	return helper.ToResponse(c, nil, pemesan)
}

func PostPemesan(c echo.Context) (err error) {
	pemesan := new(model.Pemesan)
	if err = c.Bind(pemesan); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	if err = c.Validate(pemesan); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	pemesan.Id = uuid.Must(uuid.NewV4(), err).String()

	dbm := db.Manager().Debug()
	qc := dbm.Create(pemesan)
	if qc.RowsAffected == 0 || qc.Error != nil {
		dataPemesan := model.Pemesan{}
		if err := dbm.Where("kontak = ?", pemesan.Kontak).Take(&dataPemesan).Error; err != nil {
			return helper.ToResponse(c, err, nil)
		}
		response := &map[string]interface{}{
			"status":  "Success",
			"error":   false,
			"message": "Data pemesan sudah ada",
			"data":    dataPemesan,
		}

		return c.JSON(http.StatusOK, response)
	}

	return helper.ToResponse(c, nil, pemesan)
}

func PutPemesan(c echo.Context) (err error) {
	param := new(model.PutPemesan)
	if err = c.Bind(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	if err = c.Validate(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}

	dbm := db.Manager().Debug()
	if qu := dbm.Model(&model.Pemesan{}).Updates(param); qu.Error != nil {
		return helper.ToResponse(c, err, nil)
	} else if qu.RowsAffected == 0 {
		return helper.ToResponse(c, fmt.Errorf("Data pemesan tidak ditemukan"), nil)
	}

	return helper.ToResponse(c, nil, nil)
}

func DelPemesan(c echo.Context) (err error) {
	param := new(model.DelPemesan)
	if err = c.Bind(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	if err = c.Validate(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}

	dbm := db.Manager().Debug()
	if qd := dbm.Delete(param); qd.Error != nil {
		return helper.ToResponse(c, qd.Error, nil)
	} else if qd.RowsAffected == 0 {
		return helper.ToResponse(c, fmt.Errorf("Data pemesan tidak ditemukan"), nil)
	}

	return helper.ToResponse(c, nil, nil)
}
