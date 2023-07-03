package routelist

import (
	"be-renthall/db"
	"be-renthall/helper"
	"be-renthall/model"
	"fmt"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func GetPesanan(c echo.Context) error {
	dbm := db.Manager().Debug()
	pesanan := []model.PesananDetail{}
	if err := dbm.Raw("select ps.*, gd.nama gedung_nama, pm.nama pemesan_nama, pm.kontak pemesan_kontak, rs.nama status_nama from pesanan ps left join pemesan pm on ps.pemesan_id = pm.id left join gedung gd on ps.gedung_id = gd.id left join ref_status rs on ps.status_id = rs.id order by ps.status_id").Scan(&pesanan).Error; err != nil {
		return helper.ToResponse(c, err, nil)
	}

	return helper.ToResponse(c, nil, pesanan)
}

func GetPesananById(c echo.Context) error {
	id := c.Param("id")
	dbm := db.Manager().Debug()
	pesanan := model.PesananDetail{}

	if err := dbm.Raw("select ps.*, gd.nama gedung_nama, pm.nama pemesan_nama, pm.kontak pemesan_kontak, rs.nama status_nama from pesanan ps left join pemesan pm on ps.pemesan_id = pm.id left join gedung gd on ps.gedung_id = gd.id left join ref_status rs on ps.status_id = rs.id where ps.id = ?", id).Scan(&pesanan).Error; err != nil {
		return helper.ToResponse(c, err, nil)
	}

	return helper.ToResponse(c, nil, pesanan)
}

func PostPesanan(c echo.Context) (err error) {
	pesanan := new(model.Pesanan)
	if err = c.Bind(pesanan); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	if err = c.Validate(pesanan); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	pesanan.Id = uuid.Must(uuid.NewV4(), err).String()

	dbm := db.Manager().Debug()
	pesananCek := &model.Pesanan{}
	cekTanggal := dbm.Table("pesanan").Where("gedung_id = ? and (? <= tanggal_selesai AND ? >= tanggal_mulai)", pesanan.GedungId, pesanan.TanggalMulai, pesanan.TanggalSelesai).Find(&pesananCek)
	if cekTanggal.RowsAffected > 0 {
		return helper.ToResponse(c, fmt.Errorf("Gedung tidak tersedia untuk tanggal dan jam tersebut"), nil)
	}

	qc := dbm.Create(pesanan)
	if qc.RowsAffected == 0 || qc.Error != nil {
		return helper.ToResponse(c, fmt.Errorf("Data pesanan sudah ada"), nil)
	}

	return helper.ToResponse(c, nil, pesanan)
}

func PutPesanan(c echo.Context) (err error) {
	param := new(model.PutPesanan)
	if err = c.Bind(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	if err = c.Validate(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}

	dbm := db.Manager().Debug()
	if qu := dbm.Model(&model.Pesanan{}).Updates(param); qu.Error != nil {
		return helper.ToResponse(c, err, nil)
	} else if qu.RowsAffected == 0 {
		return helper.ToResponse(c, fmt.Errorf("Data pesanan tidak ditemukan"), nil)
	}

	return helper.ToResponse(c, nil, nil)
}

func DelPesanan(c echo.Context) (err error) {
	param := new(model.DelPesanan)
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
		return helper.ToResponse(c, fmt.Errorf("Data pesanan tidak ditemukan"), nil)
	}

	return helper.ToResponse(c, nil, nil)
}

func PutPesananStatus(c echo.Context) (err error) {
	param := new(model.PutPesananStatus)
	if err = c.Bind(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}
	if err = c.Validate(param); err != nil {
		return helper.ToResponse(c, err, nil)
	}

	dbm := db.Manager().Debug()
	if qu := dbm.Where("id = ?", param.Id).Table("pesanan").Update("status_id", param.StatusId); qu.Error != nil {
		return helper.ToResponse(c, err, nil)
	} else if qu.RowsAffected == 0 {
		return helper.ToResponse(c, fmt.Errorf("Data pesanan tidak ditemukan"), nil)
	}

	return helper.ToResponse(c, nil, nil)
}

func Dashboard(c echo.Context) error {
	dbm := db.Manager().Debug()
	dashboard := model.Dashboard{}
	dbm.Raw("select count(id) pemesan from pemesan where deleted_at is null").Find(&dashboard)
	dbm.Raw("select count(id) gedung from gedung where deleted_at is null").Find(&dashboard)
	dbm.Raw("select count(id) pesanan from pesanan where deleted_at is null").Find(&dashboard)
	dbm.Raw("select count(id) diajukan from pesanan where deleted_at is null and status_id = '1'").Find(&dashboard)
	dbm.Raw("select count(id) disetujui from pesanan where deleted_at is null and status_id = '2'").Find(&dashboard)
	dbm.Raw("select count(id) ditolak from pesanan where deleted_at is null and status_id = '3'").Find(&dashboard)

	return helper.ToResponse(c, nil, dashboard)
}

func Schedule(c echo.Context) error {
	dbm := db.Manager().Debug()
	schedule := []model.Schedule{}
	dbm.Raw("select ps.id, gd.nama gedung, pm.nama pemesan, pm.kontak kontak, ps.tanggal_mulai mulai, ps.tanggal_selesai selesai " +
		"from pesanan ps left join gedung gd on ps.gedung_id = gd.id " +
		"left join pemesan pm on pm.id = ps.pemesan_id " +
		"where ps.status_id = '2' " +
		"group by ps.gedung_id, gd.nama, pm.nama, pm.kontak, ps.id " +
		"order by ps.tanggal_mulai").Find(&schedule)

	return helper.ToResponse(c, nil, schedule)
}
