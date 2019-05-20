package base

import (
	"fmt"
	"reflect"
)

//TODO
//TODO
//TODO add email
func Match(country string) {
	go matchByBoy(country)
	go matchByGirl(country)
}
func matchByBoy(country string) {
	go func() {
		for i := 19; i < 29; i++ {
			boys := BOYS[country][int32(i)]
			if len(boys) == 0 {
				continue
			}
			for _, boy := range boys {
				if reflect.ValueOf(boy).IsValid() {
					continue
				}
				for j := boy.ExpectGirl.AgeMin; j < boy.ExpectGirl.AgeMax; j++ {
					girls := GIRLS[boy.ExpectGirl.Country][j]
					for _, girl := range girls {
						if boy.ExpectGirl.HeightMax < girl.Height || boy.ExpectGirl.HeightMin > girl.Height {
							continue
						}
						if girl.ExpectBoy.HeightMax < boy.Height || girl.ExpectBoy.HeightMin > boy.Height ||
							girl.ExpectBoy.AgeMax < boy.Age || girl.ExpectBoy.AgeMin > boy.Age {
							continue
						}
						fmt.Printf("%s.%s--~**~--%s.%s\n", boy.Nickname, boy.Email,
							girl.Email, girl.Nickname)
					}
				}
			}
		}
	}()
	go func() {
		for i := 29; i < 39; i++ {
			boys := BOYS[country][int32(i)]
			for _, boy := range boys {
				if reflect.ValueOf(boy).IsValid() {
					continue
				}
				for j := boy.ExpectGirl.AgeMin; j < boy.ExpectGirl.AgeMax; j++ {
					girls := GIRLS[boy.ExpectGirl.Country][j]
					for _, girl := range girls {
						if boy.ExpectGirl.HeightMax < girl.Height || boy.ExpectGirl.HeightMin > girl.Height {
							continue
						}
						if girl.ExpectBoy.HeightMax < boy.Height || girl.ExpectBoy.HeightMin > boy.Height ||
							girl.ExpectBoy.AgeMax < boy.Age || girl.ExpectBoy.AgeMin > boy.Age {
							continue
						}
						fmt.Printf("%s.%s--~**~--%s.%s\n", boy.Nickname, boy.Email,
							girl.Email, girl.Nickname)
					}
				}
			}
		}
	}()
}
func matchByGirl(country string) {
	//TODO
}
