package minitask4

type dataDiri struct {
	nama string
	foto string
	email string
	umur int16
	nomor_tlp string
	status_pernikahan string
	riwayat_pendidikan riwayatPendidikan
}

type riwayatPendidikan struct {
	nama string
	jurusan string
}