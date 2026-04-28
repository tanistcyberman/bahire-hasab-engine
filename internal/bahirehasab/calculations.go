package bahirehasab
import "strings"


var EthiopianMonths = []string{
	"meskerem",
	"tikimt",
	"hidar",
	"tahsas",
	"ter",	
	"yekatit",
	"megabit",
	"miazia",
	"ginbot",
	"sene",
	"hamle",
	"nehasse",
	"pagume",
}

func  AmeteAlem (year int) int{
	return 5500 + year
 
}
func Wengelawi(year int) string {

	ameteAlem := AmeteAlem(year)
	remainder := ameteAlem % 4

	switch remainder {
	case 1:
		return "Matthew"
	case 2:
		return "Mark"
	case 3:
		return "Luke"
	default:
		return "John"
	}
}
func MeteneRabiet(year int) int {	
	return AmeteAlem(year) / 4
}
func TinteQemer(year int) int {
	aa := AmeteAlem(year)
	mr := MeteneRabiet(year)

	return (aa + mr) % 7
}


func Medeb(year int) int {
	return AmeteAlem(year) % 19 
}

func Wenber(year int) int {
	medeb := Medeb(year)
	if (medeb <= 0){
		return 18
	}
	return medeb - 1
}

func Abekte(year int) int {
	w := Wenber(year)
	return (w * 11) % 30
	
}

func IsValidMetqi(m int) bool {
    switch m {
    case 1, 3, 6, 9, 11, 14, 17, 20, 22, 25, 28:
        return false
    default:
        return true
    }
}

func Metqi(year int) int {
	w := Wenber(year)
	return (w * 19) % 30
	
}
func BealeMetqi(year int ) EthiopianDate {
	metqi := Metqi(year)
	if metqi > 14 {
		return EthiopianDate{
			Day: metqi,
			Month: EthiopianMonths[0],
				}
	}
	return EthiopianDate{
		Day: metqi,
		Month: EthiopianMonths[1],
	}
}
func WeekdayOfBealeMetqi(year int) int{
	tq := TinteQemer(year)
	bm := BealeMetqi(year)

	offset := bm.Day - 1 
	weekday := (tq + offset) % 7
	return weekday
}

func Tewsak(weekday int) int {
	switch weekday {
	case 0:
		return 6
	case 1:
		return 5
	case 2:
		return 4
	case 3:
		return 3
	case 4:
		return 2
	case 5:
		return 8
	case 6:
		return 7
	default:
		return 0
	}
}

func MebajaHamer(year int) int {
	metqi := Metqi(year)
	weekday := WeekdayOfBealeMetqi(year)
	tewsak := Tewsak(weekday)

	total := metqi + tewsak

	if total > 30 {
		return total - 30
	}
	return total
}	

func Nenewe(year int) EthiopianDate {
	mh := MebajaHamer(year)
	metqi := Metqi(year)

	if metqi > 14 {
		return EthiopianDate{Day: mh, Month: "ter"}
	}

	return EthiopianDate{Day: mh, Month: "yekatit"}
}

type EthiopianDate struct{
	Day  int `json:"day"`
	Month string `json:"month"`
}

func AddDays(day int, monthIndex int, daysToAdd int) EthiopianDate {
	total := day + daysToAdd

	for total > 30 {
		total -= 30
		monthIndex++
	}

	return EthiopianDate{
		Day: total,
		Month: EthiopianMonths[monthIndex],
	}
}

func Fasika(year int) EthiopianDate {
	n := Nenewe(year)

	monthIndex := 0
	for i, m := range EthiopianMonths {
		if m == n.Month {
			monthIndex = i
			break
		}
	}

	return AddDays(n.Day, monthIndex, 69)
}

func AbiyTsom(year int) EthiopianDate {
	n := Nenewe(year)
	return AddDays(n.Day, getMonthIndex(n.Month), 14)
}

func DebreZeit(year int) EthiopianDate{
	n := Nenewe(year)
	monthIndex := 0
	for i, m := range EthiopianMonths {
		if m == n.Month {
			monthIndex = i
			break
		}
	}
	return AddDays(n.Day, monthIndex, 41)
}

func Hosanna(year int) EthiopianDate{
	n := Nenewe(year)
	monthIndex := 0 
	for i, m := range EthiopianMonths {
		if m == n.Month {
			monthIndex = i
			break
		}
}
	return AddDays(n.Day, monthIndex, 62)
}
 
func Siklet(year int) EthiopianDate {
	n := Nenewe(year)

	monthIndex := 0
	for i, m := range EthiopianMonths {
		if m == n.Month {
			monthIndex = i
			break
		}
	}

	return AddDays(n.Day, monthIndex, 67)
}

 func getMonthIndex(month string) int {
	for i, m := range EthiopianMonths {
		if strings.ToLower(m) == strings.ToLower(month) {
			return i
		}
	}
	return 0
}

 func RikbeKahinat(year int) EthiopianDate{
	n := Nenewe(year)
	monthIndex := getMonthIndex(n.Month)
	return AddDays(n.Day, monthIndex, 93)
 }
 func Ascension(year int) EthiopianDate{
	n := Nenewe(year)
	monthIndex := getMonthIndex(n.Month)
	return AddDays(n.Day, monthIndex, 108)
 }
 func Paraclete(year int) EthiopianDate {
	n := Nenewe(year)
	monthIndex := getMonthIndex(n.Month)
	return AddDays(n.Day, monthIndex, 118)
 } 

 func Hawaryat(year int) EthiopianDate {
	n := Nenewe(year)
	monthIndex := getMonthIndex(n.Month)
	return AddDays(n.Day, monthIndex, 119)
 }

 type BahireHasabResult struct {
	Year       int           `json:"year"`
	Nenewe     EthiopianDate `json:"nenewe"`
	AbiyTsom   EthiopianDate `json:"abiy_tsom"`
	DebreZeit  EthiopianDate `json:"debre_zeit"`
	Hosanna    EthiopianDate `json:"hosanna"`
	Siklet     EthiopianDate `json:"siklet"`
	Fasika     EthiopianDate `json:"fasika"`
}

func CalculateBahireHasab(year int) BahireHasabResult {
	return BahireHasabResult{
		Year:      year,
		Nenewe:    Nenewe(year),
		AbiyTsom:  AbiyTsom(year),
		DebreZeit: DebreZeit(year),
		Hosanna:   Hosanna(year),
		Siklet:    Siklet(year),
		Fasika:    Fasika(year),
	}
}


