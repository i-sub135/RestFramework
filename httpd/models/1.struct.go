package models

type (

	// WilayahProvinsi -- t_wilayah_propinsi on db cw_master
	WilayahProvinsi struct {
		Kode string `json:"Kode"`
		Nama string `json:"Nama"`
	}

	// WilayahKota -- t_wilayah_kota on db cw_master
	WilayahKota struct {
		Kode string `json:"Kode"`
		Nama string `json:"Nama"`
	}

	// WilayahKecamatan -- t_wilayah_kecamatan on db cw_master
	WilayahKecamatan struct {
		Kode string `json:"Kode"`
		Nama string `json:"Nama"`
	}
	// WilayahKelurahan -- t_wilayah_Kelurahan on db cw_master
	WilayahKelurahan struct {
		Kode string `json:"Kode"`
		Nama string `json:"Nama"`
	}
)
