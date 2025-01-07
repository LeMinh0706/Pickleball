package util

import (
	"math/rand"
)

func RandomAvatar(gender int32) string {
	if gender == 0 {
		women := []string{"upload/avatar/avatar_women_1.jpg", "upload/avatar/avatar_women_2.jpg", "upload/avatar/avatar_women_3.jpg", "upload/avatar/avatar_women_4.jpg"}
		n := len(women)
		return women[rand.Intn(n)]
	} else {
		men := []string{"upload/avatar/avatar_men_1.jpg", "upload/avatar/avatar_men_2.jpg", "upload/avatar/avatar_men_3.jpg", "upload/avatar/avatar_men_4.jpg"}
		n := len(men)
		return men[rand.Intn(n)]
	}
}
