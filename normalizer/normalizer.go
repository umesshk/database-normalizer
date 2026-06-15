package normalizer

import "regexp"

// func Normalize(phone string) string {
//
// 	var new_phone bytes.Buffer
//
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			new_phone.WriteRune(ch)
// 		}
// 	}
//
// 	return new_phone.String()
// }

func Normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}
