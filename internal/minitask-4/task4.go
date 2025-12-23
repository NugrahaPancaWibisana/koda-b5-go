package minitask4

import "fmt"

func DataDiri()  {
	dataDiri := dataDiri{
		nama: "Nugraha Panca Wibisana",
		foto: "https://media.licdn.com/dms/image/v2/D5603AQHfrRHBDvIwFQ/profile-displayphoto-scale_200_200/B56ZsmeEcWIEAc-/0/1765876978721?e=1767830400&v=beta&t=-JQxyAGXvR3vRepSBm5XqGi7TSuF9-KUCGR3YY4i8UA",
		email: "nugrahapancawibisana@gmail.com",
		umur: 18,
		nomor_tlp: "085712619452",
		status_pernikahan: "Belum Menikah",
		riwayat_pendidikan: riwayatPendidikan{
			nama: "SMK Negeri 1 Subang",
			jurusan: "Rekayasa Perangkat Lunak",
		},
	}

	fmt.Println(dataDiri)
}