package base

func (this *Boy) NewObject(expectGirl *ExpectGirl) {
	if expectGirl.Country == "" {
		expectGirl.Country = this.Country
	}
	if expectGirl.City == "" {
		expectGirl.City = this.City
	}
	if expectGirl.AgeMin < 19 {
		expectGirl.AgeMin = 19
	}
	this.ExpectGirl = expectGirl
	if BOYS[this.Country] == nil {
		BOYS[this.Country] = make(map[int32][]Boy, 520)
		BOYS[this.Country][this.Age] = make([]Boy, 520)
	}
	BOYS[this.Country][this.Age] = append(BOYS[this.Country][this.Age], *this)
}
func findTheGirl(theGirl *ExpectGirl) *[]Girl {
	var girls = make([]Girl, 520)
	var j int
	for age := theGirl.AgeMin; age <= theGirl.AgeMax; age++ {
		girlsMap := GIRLS[theGirl.Country][age]
		for _, girl := range girlsMap {
			if girl.Height < theGirl.HeightMin || girl.Height > theGirl.HeightMax {
				continue
			}
			girls[j] = girl
			j++
		}
	}
	return &girls
}
