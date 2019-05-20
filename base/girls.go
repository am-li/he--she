package base

func (this *Girl) NewObject(expectBoy *ExpectBoy) {
	if expectBoy.Country == "" {
		expectBoy.Country = this.Country
	}
	if expectBoy.City == "" {
		expectBoy.City = this.City
	}
	if expectBoy.AgeMin < 19 {
		expectBoy.AgeMin = 19
	}
	this.ExpectBoy = expectBoy
	if GIRLS[this.Country] == nil {
		GIRLS[this.Country] = make(map[int32][]Girl, 520)
		GIRLS[this.Country][this.Age] = make([]Girl, 520)
	}
	GIRLS[this.Country][this.Age] = append(GIRLS[this.Country][this.Age], *this)
}
func findTheBoy(theBoy *ExpectBoy) *[]Boy {
	var boys = make([]Boy, 520)
	var j int
	for age := theBoy.AgeMin; age <= theBoy.AgeMax; age++ {
		boysMap := BOYS[theBoy.Country][age]
		for _, boy := range boysMap {
			if boy.Height < theBoy.HeightMin || boy.Height > theBoy.HeightMax {
				continue
			}
			boys[j] = boy
			j++
		}
	}
	return &boys
}
