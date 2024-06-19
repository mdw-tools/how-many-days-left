package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const verboseDate = "Monday, January 2 of 2006"

func main() {
	sex, birthday := parseCLI()
	ageInDays := calculateAgeInDays(birthday)
	ageInYears := calculateAgeInYears(birthday)
	expectancy := getAdditionalLifeExpectancyInYears(ageInYears, sex)
	projectedDeath := calculateProjectedDeathDate(expectancy, birthday, ageInYears)
	projectedDaysRemaining := countProjectedDaysRemaining(projectedDeath)
	yearsSinceBirth := time.Now().Sub(birthday).Abs().Hours() / 24.0 / 365.0

	fmt.Printf(""+
		"Given that you have lived %s days, a "+
		"lifespan of %.2f years since your birth on %s, ",
		formatDays(ageInDays),
		yearsSinceBirth,
		birthday.Format(verboseDate))

	fmt.Printf(""+
		"and based on medicaid's average life expectancy "+
		"for %ss, you have %s days remaining until "+
		"reaching your projected lifespan of %.2f "+
		"years on %s.\n",
		formatSex[sex],
		formatDays(projectedDaysRemaining),
		float64(ageInYears)+expectancy,
		projectedDeath.Format(verboseDate),
	)
	fmt.Println()
	fmt.Println("Event  Date         Age")
	fmt.Println("------------------------------------------")
	fmt.Printf(""+
		"Birth: %s  0.00 (%s days ago)\n",
		birthday.Format(time.DateOnly),
		formatDays(ageInDays),
	)
	fmt.Printf(""+
		"Today: %s %-5.2f\n",
		time.Now().Format(time.DateOnly),
		yearsSinceBirth,
	)
	fmt.Printf(""+
		"Death: %s %-5.2f (%s days left)\n",
		projectedDeath.Format(time.DateOnly),
		float64(ageInYears)+expectancy,
		formatDays(projectedDaysRemaining),
	)
}

func parseCLI() (sex string, birthday time.Time) {
	var birth string
	flag.StringVar(&birth, "birth", "", "Birth date (YYYY-MM-DD)")
	flag.StringVar(&sex, "sex", "", "Sex (one of 'f' or 'm')")

	flag.Parse()

	birthday, err := time.Parse(time.DateOnly, birth)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal("Failed to parse date of birth.")
	}

	if sex != "m" && sex != "f" {
		flag.PrintDefaults()
		log.Fatal("Failed to parse sex.")
	}

	return sex, birthday
}

func calculateAgeInDays(birth time.Time) (days int) {
	for now := time.Now(); birth.Before(now); days++ {
		birth = birth.AddDate(0, 0, 1)
	}
	return days
}
func calculateAgeInYears(birth time.Time) (years int) {
	for now := time.Now(); birth.Before(now); years++ {
		birth = birth.AddDate(1, 0, 0)
	}
	return years
}
func getAdditionalLifeExpectancyInYears(age int, sex string) float64 {
	if sex == "f" {
		return females[age]
	} else {
		return males[age]
	}
}
func calculateProjectedDeathDate(expectancy float64, birthday time.Time, age int) time.Time {
	yearsToAdd := int(expectancy)
	daysToAdd := int(365.0 * (expectancy - float64(yearsToAdd)))
	return birthday.AddDate(age, 0, 0).AddDate(yearsToAdd, 0, daysToAdd)
}
func countProjectedDaysRemaining(projectedDeath time.Time) (projectedDaysRemaining int) {
	for d := time.Now(); d.Before(projectedDeath); projectedDaysRemaining++ {
		d = d.AddDate(0, 0, 1)
	}
	return projectedDaysRemaining
}

var formatSex = map[string]string{
	"f": "female",
	"m": "male",
}

func formatDays(days int) string {
	DAYS := fmt.Sprint(days)
	if days >= 10000 {
		return DAYS[:2] + "," + DAYS[2:]
	}
	if days >= 1000 {
		return DAYS[:1] + "," + DAYS[1:]
	}
	return DAYS
}

// 2023 Medicaid Life Expectancy Tables:
// https://www.health.ny.gov/health_care/medicaid/publications/docs/gis/23ma09_att1.pdf
var males = map[int]float64{
	0:   74.12,
	1:   73.55,
	2:   72.58,
	3:   71.60,
	4:   70.62,
	5:   69.63,
	6:   68.64,
	7:   67.65,
	8:   66.65,
	9:   65.66,
	10:  64.67,
	11:  63.68,
	12:  62.69,
	13:  61.70,
	14:  60.71,
	15:  59.73,
	16:  58.76,
	17:  57.79,
	18:  56.84,
	19:  55.90,
	20:  54.97,
	21:  54.04,
	22:  53.12,
	23:  52.21,
	24:  51.30,
	25:  50.39,
	26:  49.48,
	27:  48.57,
	28:  47.66,
	29:  46.76,
	30:  45.86,
	31:  44.97,
	32:  44.07,
	33:  43.18,
	34:  42.29,
	35:  41.39,
	36:  40.50,
	37:  39.62,
	38:  38.73,
	39:  37.85,
	40:  36.97,
	41:  36.09,
	42:  35.21,
	43:  34.34,
	44:  33.46,
	45:  32.59,
	46:  31.73,
	47:  30.87,
	48:  30.01,
	49:  29.17,
	50:  28.33,
	51:  27.50,
	52:  26.67,
	53:  25.86,
	54:  25.06,
	55:  24.27,
	56:  23.48,
	57:  22.71,
	58:  21.95,
	59:  21.21,
	60:  20.47,
	61:  19.74,
	62:  19.03,
	63:  18.32,
	64:  17.63,
	65:  16.94,
	66:  16.26,
	67:  15.58,
	68:  14.91,
	69:  14.24,
	70:  13.59,
	71:  12.94,
	72:  12.30,
	73:  11.67,
	74:  11.05,
	75:  10.46,
	76:  9.88,
	77:  9.32,
	78:  8.77,
	79:  8.25,
	80:  7.74,
	81:  7.25,
	82:  6.77,
	83:  6.31,
	84:  5.88,
	85:  5.47,
	86:  5.07,
	87:  4.70,
	88:  4.35,
	89:  4.02,
	90:  3.72,
	91:  3.44,
	92:  3.18,
	93:  2.96,
	94:  2.75,
	95:  2.57,
	96:  2.42,
	97:  2.28,
	98:  2.15,
	99:  2.04,
	100: 1.93,
	101: 1.83,
	102: 1.73,
	103: 1.63,
	104: 1.54,
	105: 1.45,
	106: 1.36,
	107: 1.28,
	108: 1.20,
	109: 1.13,
	110: 1.05,
	111: 0.98,
	112: 0.92,
	113: 0.85,
	114: 0.79,
	115: 0.74,
	116: 0.68,
	117: 0.63,
	118: 0.58,
	119: 0.53,
}
var females = map[int]float64{
	0:   79.78,
	1:   79.17,
	2:   78.19,
	3:   77.21,
	4:   76.22,
	5:   75.23,
	6:   74.24,
	7:   73.25,
	8:   72.25,
	9:   71.26,
	10:  70.27,
	11:  69.27,
	12:  68.28,
	13:  67.29,
	14:  66.30,
	15:  65.31,
	16:  64.32,
	17:  63.34,
	18:  62.36,
	19:  61.38,
	20:  60.41,
	21:  59.44,
	22:  58.47,
	23:  57.50,
	24:  56.54,
	25:  55.58,
	26:  54.61,
	27:  53.66,
	28:  52.70,
	29:  51.74,
	30:  50.79,
	31:  49.84,
	32:  48.89,
	33:  47.94,
	34:  47.00,
	35:  46.06,
	36:  45.12,
	37:  44.18,
	38:  43.24,
	39:  42.31,
	40:  41.38,
	41:  40.45,
	42:  39.52,
	43:  38.60,
	44:  37.68,
	45:  36.76,
	46:  35.85,
	47:  34.94,
	48:  34.04,
	49:  33.14,
	50:  32.24,
	51:  31.35,
	52:  30.47,
	53:  29.59,
	54:  28.72,
	55:  27.86,
	56:  27.01,
	57:  26.16,
	58:  25.32,
	59:  24.49,
	60:  23.67,
	61:  22.85,
	62:  22.04,
	63:  21.24,
	64:  20.45,
	65:  19.66,
	66:  18.88,
	67:  18.10,
	68:  17.34,
	69:  16.58,
	70:  15.82,
	71:  15.08,
	72:  14.36,
	73:  13.64,
	74:  12.94,
	75:  12.26,
	76:  11.60,
	77:  10.95,
	78:  10.31,
	79:  9.70,
	80:  9.10,
	81:  8.53,
	82:  7.98,
	83:  7.44,
	84:  6.93,
	85:  6.44,
	86:  5.99,
	87:  5.55,
	88:  5.15,
	89:  4.76,
	90:  4.41,
	91:  4.08,
	92:  3.78,
	93:  3.51,
	94:  3.27,
	95:  3.05,
	96:  2.85,
	97:  2.68,
	98:  2.52,
	99:  2.37,
	100: 2.23,
	101: 2.09,
	102: 1.96,
	103: 1.84,
	104: 1.72,
	105: 1.61,
	106: 1.50,
	107: 1.40,
	108: 1.30,
	109: 1.21,
	110: 1.12,
	111: 1.03,
	112: 0.95,
	113: 0.88,
	114: 0.80,
	115: 0.74,
	116: 0.68,
	117: 0.63,
	118: 0.58,
	119: 0.53,
}
