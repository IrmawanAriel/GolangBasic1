package basictask1

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi int
}

func (bangun kubus) luas() int {
	totalLuas := bangun.sisi * bangun.sisi
	return totalLuas
}

func (bangun kubus) keliling() int {
	totalKeliling := bangun.sisi * 4
	return totalKeliling
}

func (bangun kubus) volume() int {
	totalVolume := bangun.sisi * bangun.sisi * bangun.sisi
	return totalVolume
}
