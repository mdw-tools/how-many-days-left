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
	yearsRemaining := time.Now().Sub(projectedDeath).Abs().Hours() / 24.0 / 365.0

	fmt.Printf(""+
		"Congratulations, you have lived %s days, a "+
		"lifespan of %.2f years since your birth on %s.\n",
		formatDays(ageInDays),
		yearsSinceBirth,
		birthday.Format(verboseDate))

	fmt.Printf(""+
		"Based on medicaid's average life expectancy "+
		"for %ss you have %s days remaining until "+
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
		yearsRemaining,
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

// See https://www.health.ny.gov/health_care/medicaid/publications/docs/adm/06adm-5att8.pdf
var males = map[int]float64{
	0:   74.14,
	1:   73.70,
	2:   72.74,
	3:   71.77,
	4:   70.79,
	5:   69.81,
	6:   68.82,
	7:   67.83,
	8:   66.84,
	9:   65.85,
	10:  64.86,
	11:  63.87,
	12:  62.88,
	13:  61.89,
	14:  60.91,
	15:  59.93,
	16:  58.97,
	17:  58.02,
	18:  57.07,
	19:  56.14,
	20:  55.20,
	21:  54.27,
	22:  53.35,
	23:  52.42,
	24:  51.50,
	25:  50.57,
	26:  49.64,
	27:  48.71,
	28:  47.77,
	29:  46.84,
	30:  45.90,
	31:  44.96,
	32:  44.03,
	33:  43.09,
	34:  42.16,
	35:  41.23,
	36:  40.30,
	37:  39.38,
	38:  38.46,
	39:  37.55,
	40:  36.64,
	41:  35.73,
	42:  34.83,
	43:  33.94,
	44:  33.05,
	45:  32.16,
	46:  31.29,
	47:  30.42,
	48:  29.56,
	49:  28.70,
	50:  27.85,
	51:  27.00,
	52:  26.16,
	53:  25.32,
	54:  24.50,
	55:  23.68,
	56:  22.86,
	57:  22.06,
	58:  21.27,
	59:  20.49,
	60:  19.72,
	61:  18.96,
	62:  18.21,
	63:  17.48,
	64:  16.76,
	65:  16.05,
	66:  15.36,
	67:  14.68,
	68:  14.02,
	69:  13.38,
	70:  12.75,
	71:  12.13,
	72:  11.53,
	73:  10.95,
	74:  10.38,
	75:  9.83,
	76:  9.29,
	77:  8.77,
	78:  8.27,
	79:  7.78,
	80:  7.31,
	81:  6.85,
	82:  6.42,
	83:  6.00,
	84:  5.61,
	85:  5.24,
	86:  4.89,
	87:  4.56,
	88:  4.25,
	89:  3.97,
	90:  3.70,
	91:  3.45,
	92:  3.22,
	93:  3.01,
	94:  2.82,
	95:  2.64,
	96:  2.49,
	97:  2.35,
	98:  2.22,
	99:  2.11,
	100: 2.00,
	101: 1.89,
	102: 1.79,
	103: 1.69,
	104: 1.59,
	105: 1.50,
	106: 1.41,
	107: 1.33,
	108: 1.25,
	109: 1.17,
	110: 1.10,
	111: 1.03,
	112: 0.96,
	113: 0.89,
	114: 0.83,
	115: 0.77,
	116: 0.71,
	117: 0.66,
	118: 0.61,
	119: 0.56,
}
var females = map[int]float64{
	0:   74.14,
	1:   73.70,
	2:   72.74,
	3:   71.77,
	4:   70.79,
	5:   69.81,
	6:   68.82,
	7:   67.83,
	8:   66.84,
	9:   65.85,
	10:  64.86,
	11:  63.87,
	12:  62.88,
	13:  61.89,
	14:  60.91,
	15:  59.93,
	16:  58.97,
	17:  58.02,
	18:  57.07,
	19:  56.14,
	20:  55.20,
	21:  54.27,
	22:  53.35,
	23:  52.42,
	24:  51.50,
	25:  50.57,
	26:  49.64,
	27:  48.71,
	28:  47.77,
	29:  46.84,
	30:  45.90,
	31:  44.96,
	32:  44.03,
	33:  43.09,
	34:  42.16,
	35:  41.23,
	36:  40.30,
	37:  39.38,
	38:  38.46,
	39:  37.55,
	40:  36.64,
	41:  35.73,
	42:  34.83,
	43:  33.94,
	44:  33.05,
	45:  32.16,
	46:  31.29,
	47:  30.42,
	48:  29.56,
	49:  28.70,
	50:  27.85,
	51:  27.00,
	52:  26.16,
	53:  25.32,
	54:  24.50,
	55:  23.68,
	56:  22.86,
	57:  22.06,
	58:  21.27,
	59:  20.49,
	60:  19.72,
	61:  18.96,
	62:  18.21,
	63:  17.48,
	64:  16.76,
	65:  16.05,
	66:  15.36,
	67:  14.68,
	68:  14.02,
	69:  13.38,
	70:  12.75,
	71:  12.13,
	72:  11.53,
	73:  10.95,
	74:  10.38,
	75:  9.83,
	76:  9.29,
	77:  8.77,
	78:  8.27,
	79:  7.78,
	80:  7.31,
	81:  6.85,
	82:  6.42,
	83:  6.00,
	84:  5.61,
	85:  5.24,
	86:  4.89,
	87:  4.56,
	88:  4.25,
	89:  3.97,
	90:  3.70,
	91:  3.45,
	92:  3.22,
	93:  3.01,
	94:  2.82,
	95:  2.64,
	96:  2.49,
	97:  2.35,
	98:  2.22,
	99:  2.11,
	100: 2.00,
	101: 1.89,
	102: 1.79,
	103: 1.69,
	104: 1.59,
	105: 1.50,
	106: 1.41,
	107: 1.33,
	108: 1.25,
	109: 1.17,
	110: 1.10,
	111: 1.03,
	112: 0.96,
	113: 0.89,
	114: 0.83,
	115: 0.77,
	116: 0.71,
	117: 0.66,
	118: 0.61,
	119: 0.56,
}
