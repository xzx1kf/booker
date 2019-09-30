package util

type slot struct {
	Court, Hour, Min string
}

type TimeslotMap struct {
	ts map[slot]string
}

func NewTimeslotMap() *TimeslotMap {
	tsm := &TimeslotMap{ts: make(map[slot]string)}
	tsm.Init()
	return tsm
}

func (t *TimeslotMap) Lookup(court, hour, min string) string {
	return t.ts[slot{court, hour, min}]
}

func (t *TimeslotMap) Init() {
	// Court 1
	t.ts[slot{"1", "09", "10"}] = "1"
	t.ts[slot{"1", "09", "50"}] = "2"
	t.ts[slot{"1", "10", "30"}] = "3"
	t.ts[slot{"1", "11", "10"}] = "4"
	t.ts[slot{"1", "11", "50"}] = "5"
	t.ts[slot{"1", "12", "30"}] = "6"
	t.ts[slot{"1", "13", "10"}] = "7"
	t.ts[slot{"1", "13", "50"}] = "8"
	t.ts[slot{"1", "14", "30"}] = "9"
	t.ts[slot{"1", "15", "10"}] = "10"
	t.ts[slot{"1", "15", "50"}] = "11"
	t.ts[slot{"1", "16", "30"}] = "12"
	t.ts[slot{"1", "17", "10"}] = "13"
	t.ts[slot{"1", "17", "50"}] = "14"
	t.ts[slot{"1", "18", "30"}] = "15"
	t.ts[slot{"1", "19", "10"}] = "16"
	t.ts[slot{"1", "19", "50"}] = "17"
	t.ts[slot{"1", "20", "30"}] = "18"
	t.ts[slot{"1", "21", "10"}] = "19"
	t.ts[slot{"1", "21", "50"}] = "20"

	// Court 2
	t.ts[slot{"2", "09", "10"}] = "21"
	t.ts[slot{"2", "09", "50"}] = "22"
	t.ts[slot{"2", "10", "30"}] = "23"
	t.ts[slot{"2", "11", "10"}] = "24"
	t.ts[slot{"2", "11", "50"}] = "25"
	t.ts[slot{"2", "12", "30"}] = "26"
	t.ts[slot{"2", "13", "10"}] = "27"
	t.ts[slot{"2", "13", "50"}] = "28"
	t.ts[slot{"2", "14", "30"}] = "29"
	t.ts[slot{"2", "15", "10"}] = "30"
	t.ts[slot{"2", "15", "50"}] = "31"
	t.ts[slot{"2", "16", "30"}] = "32"
	t.ts[slot{"2", "17", "10"}] = "33"
	t.ts[slot{"2", "17", "50"}] = "34"
	t.ts[slot{"2", "18", "30"}] = "35"
	t.ts[slot{"2", "19", "10"}] = "36"
	t.ts[slot{"2", "19", "50"}] = "37"
	t.ts[slot{"2", "20", "30"}] = "38"
	t.ts[slot{"2", "21", "10"}] = "39"
	t.ts[slot{"2", "21", "50"}] = "40"

	// Court 3
	t.ts[slot{"3", "09", "0"}] = "41"
	t.ts[slot{"3", "09", "40"}] = "42"
	t.ts[slot{"3", "10", "20"}] = "43"
	t.ts[slot{"3", "11", "0"}] = "44"
	t.ts[slot{"3", "11", "40"}] = "45"
	t.ts[slot{"3", "12", "20"}] = "46"
	t.ts[slot{"3", "13", "0"}] = "47"
	t.ts[slot{"3", "13", "40"}] = "48"
	t.ts[slot{"3", "14", "20"}] = "49"
	t.ts[slot{"3", "15", "0"}] = "50"
	t.ts[slot{"3", "15", "40"}] = "51"
	t.ts[slot{"3", "16", "20"}] = "52"
	t.ts[slot{"3", "17", "0"}] = "53"
	t.ts[slot{"3", "17", "40"}] = "54"
	t.ts[slot{"3", "18", "20"}] = "55"
	t.ts[slot{"3", "19", "0"}] = "56"
	t.ts[slot{"3", "19", "40"}] = "57"
	t.ts[slot{"3", "20", "20"}] = "58"
	t.ts[slot{"3", "21", "0"}] = "59"
	t.ts[slot{"3", "21", "40"}] = "60"

	// Court 4
	t.ts[slot{"4", "09", "0"}] = "61"
	t.ts[slot{"4", "09", "40"}] = "62"
	t.ts[slot{"4", "10", "20"}] = "63"
	t.ts[slot{"4", "11", "0"}] = "64"
	t.ts[slot{"4", "11", "40"}] = "65"
	t.ts[slot{"4", "12", "20"}] = "66"
	t.ts[slot{"4", "13", "0"}] = "67"
	t.ts[slot{"4", "13", "40"}] = "68"
	t.ts[slot{"4", "14", "20"}] = "69"
	t.ts[slot{"4", "15", "0"}] = "70"
	t.ts[slot{"4", "15", "40"}] = "71"
	t.ts[slot{"4", "16", "20"}] = "72"
	t.ts[slot{"4", "17", "0"}] = "73"
	t.ts[slot{"4", "17", "40"}] = "74"
	t.ts[slot{"4", "18", "20"}] = "75"
	t.ts[slot{"4", "19", "0"}] = "76"
	t.ts[slot{"4", "19", "40"}] = "77"
	t.ts[slot{"4", "20", "20"}] = "78"
	t.ts[slot{"4", "21", "0"}] = "79"
	t.ts[slot{"4", "21", "40"}] = "80"

	// Court 5
	t.ts[slot{"5", "09", "5"}] = "81"
	t.ts[slot{"5", "09", "45"}] = "82"
	t.ts[slot{"5", "10", "25"}] = "83"
	t.ts[slot{"5", "11", "5"}] = "84"
	t.ts[slot{"5", "11", "45"}] = "85"
	t.ts[slot{"5", "12", "25"}] = "86"
	t.ts[slot{"5", "13", "5"}] = "87"
	t.ts[slot{"5", "13", "45"}] = "88"
	t.ts[slot{"5", "14", "25"}] = "89"
	t.ts[slot{"5", "15", "5"}] = "90"
	t.ts[slot{"5", "15", "45"}] = "91"
	t.ts[slot{"5", "16", "25"}] = "92"
	t.ts[slot{"5", "17", "5"}] = "93"
	t.ts[slot{"5", "17", "45"}] = "94"
	t.ts[slot{"5", "18", "25"}] = "95"
	t.ts[slot{"5", "19", "5"}] = "96"
	t.ts[slot{"5", "19", "45"}] = "97"
	t.ts[slot{"5", "20", "25"}] = "98"
	t.ts[slot{"5", "21", "5"}] = "99"
	t.ts[slot{"5", "21", "45"}] = "100"
}
