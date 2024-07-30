package basictask1

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Kubus struct {
	Sisi float64
}

func (bangun Kubus) Luas() float64 {
	totalLuas := bangun.Sisi * bangun.Sisi
	return totalLuas
}

func (bangun Kubus) Keliling() float64 {
	totalKeliling := bangun.Sisi * 4
	return totalKeliling
}

func (bangun Kubus) Volume() float64 {
	totalVolume := bangun.Sisi * bangun.Sisi * bangun.Sisi
	return totalVolume
}
