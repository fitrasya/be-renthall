package routelist

import (
	"be-renthall/db"
	"be-renthall/helper"
	"be-renthall/model"
	"fmt"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func GetGedung(c echo.Context) error {
	dbm := db.Manager().Debug()
	gedung := []model.Gedung{}
	if err := dbm.Find(&gedung).Error; err != nil {
		return helper.ToResponse(c, err, nil)
	}

	return helper.ToResponse(c, nil, gedung)
}

func GetGedungById(c echo.Context) error {
	id := c.Param("id")
	dbm := db.Manager().Debug()
	gedung := model.Gedung{}
	if err := dbm.Where("id = ?", id).Take(&gedung).Error; err != nil {
		return helper.ToResponse(c, err, nil)
	}

	return helper.ToResponse(c, nil, gedung)
}

func PostGedung(c echo.Context) (err error) {
	gedung := new(model.Gedung)
	if err = c.Bind(gedung); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	if err = c.Validate(gedung); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	gedung.Id = uuid.Must(uuid.NewV4(), err).String()

	dbm := db.Manager().Debug()
	qc := dbm.Create(gedung)
	if qc.RowsAffected == 0 || qc.Error != nil {
		return helper.ToResponse(c, fmt.Errorf("Data gedung sudah ada"), nil)
	}

	return helper.ToResponse(c, nil, gedung)
}

func PutGedung(c echo.Context) (err error) {
	param := new(model.PutGedung)
	if err = c.Bind(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	if err = c.Validate(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}

	dbm := db.Manager().Debug()
	if qu := dbm.Model(&model.Gedung{}).Updates(param); qu.Error != nil {
		return helper.ToResponse(c, err, nil)
	} else if qu.RowsAffected == 0 {
		return helper.ToResponse(c, fmt.Errorf("Data gedung tidak ditemukan"), nil)
	}

	return helper.ToResponse(c, nil, nil)
}

func DelGedung(c echo.Context) (err error) {
	param := new(model.DelGedung)
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
		return helper.ToResponse(c, fmt.Errorf("Data gedung tidak ditemukan"), nil)
	}

	return helper.ToResponse(c, nil, nil)
}
