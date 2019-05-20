package countries

import "github.com/am-li/he-she/base"

func China(c chan string) {
	boys()
	girls()
	c <- "china"
}
func boys() {
	my := base.Boy{
		Height:    178,
		Weight:    65,
		Age:       23,
		Education: "",
		Career:    "IT",
		Country:   "china",
		City:      "somewhere",
		Contact:   "github.com/am-li",
		Email:     "yzxkhd@gmail.com",
	}
	my.NewObject(&base.ExpectGirl{
		HeightMin: 100,
		HeightMax: 200,
		Weight:    65,
		AgeMin:    19,
		AgeMax:    26,
		Education: "",
		Career:    "",
		Country:   "china",
		City:      "somewhere",
	})
}
func girls() {

}
