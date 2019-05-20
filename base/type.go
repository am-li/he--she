package base

type (
	Boy struct {
		Nickname   string
		Height     int32 //cm
		Weight     int32 //kg
		Age        int32
		Education  string //
		Career     string
		Country    string
		City       string //work city
		Car        bool
		House      bool
		Contact    string //contact way
		Email      string
		ExpectGirl *ExpectGirl
	}
	Girl struct {
		Nickname  string
		Height    int32 //cm
		Weight    int32 //kg
		Age       int32
		Education string //
		Career    string
		Country   string
		City      string //work city
		Car       bool
		House     bool
		Contact   string //contact way
		Email     string
		ExpectBoy *ExpectBoy
	}
	ExpectGirl struct {
		AgeMin    int32
		AgeMax    int32
		HeightMin int32  //cm
		HeightMax int32  //cm
		Weight    int32  //kg
		Education string //
		Career    string
		Country   string
		City      string //work city
		Car       bool
		House     bool
	}
	ExpectBoy struct {
		AgeMin    int32
		AgeMax    int32
		HeightMin int32  //cm
		HeightMax int32  //cm
		Weight    int32  //kg
		Education string //
		Career    string
		Country   string
		City      string //work city
		Car       bool
		House     bool
	}
)

//map[country]map[age]
var BOYS = make(map[string]map[int32][]Boy, 1314*520)
var GIRLS = make(map[string]map[int32][]Girl, 1314*520)
