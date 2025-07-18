package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//go:embed medicaid-data
var medicaid embed.FS

var Version = "dev"

const verboseDate = "Monday, January 2 of 2006"
const daysPerYear = 365.0 // ignoring leap years
const hoursPerDay = 24.0  // ignoring leap seconds

func main() {
	var (
		sex, birthday, dataYear = parseCLI()
		data                    = parseLifeExpectancyTable(dataYear)
		ageInDays               = calculateAgeInDays(birthday)
		ageInYears              = calculateAgeInYears(birthday)
		expectancy              = data[sex][ageInYears]
		projectedDeath          = calculateProjectedDeathDate(expectancy, birthday, ageInYears)
		projectedDaysRemaining  = countProjectedDaysRemaining(projectedDeath)
		yearsSinceBirth         = time.Now().Sub(birthday).Abs().Hours() / hoursPerDay / daysPerYear
	)
	fmt.Printf(""+
		"Given that you have lived %s days, a "+
		"lifespan of %.2f years since your birth on %s, ",
		formatDays(ageInDays),
		yearsSinceBirth,
		birthday.Format(verboseDate),
	)

	fmt.Printf(""+
		"and based on the %d release of medicaid's average life expectancy "+
		"for %ss, you have %s days remaining until "+
		"reaching your projected lifespan of %.2f "+
		"years on %s.\n",
		dataYear,
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

func parseLifeExpectancyTable(year int) map[string]map[int]float64 {
	result := make(map[string]map[int]float64)
	result["m"] = make(map[int]float64)
	result["f"] = make(map[int]float64)
	file, err := medicaid.Open(fmt.Sprintf("medicaid-data/%d.txt", year))
	if err != nil {
		log.Fatalln("Failed to read data file:", err)
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		age, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalln("Failed to parse age:", err)
		}

		spanMale, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Fatalln("Failed to parse male life expectancy:", err)
		}
		result["m"][age] = spanMale

		spanFemale, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			log.Fatalln("Failed to parse female life expectancy:", err)
		}
		result["f"][age] = spanFemale
	}
	return result
}

func parseCLI() (sex string, birthday time.Time, dataYear int) {
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	var birth string
	flags.StringVar(&birth, "birth", "", "Birth date (YYYY-MM-DD)")
	flags.StringVar(&sex, "sex", "", "Sex (one of 'f' or 'm')")
	flags.IntVar(&dataYear, "data-year", 2024, "Data year")
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage of %s:\n", flags.Name())
		flags.PrintDefaults()
	}
	_ = flags.Parse(os.Args[1:])

	birthday, err := time.Parse(time.DateOnly, birth)
	if err != nil {
		flags.PrintDefaults()
		log.Fatal("Failed to parse date of birth.")
	}

	if sex != "m" && sex != "f" {
		flags.PrintDefaults()
		log.Fatal("Failed to parse sex.")
	}

	return sex, birthday, dataYear
}

func calculateAgeInDays(birth time.Time) (result int) {
	for now := time.Now(); birth.Before(now); result++ {
		birth = birth.AddDate(0, 0, 1)
	}
	return result
}
func calculateAgeInYears(birth time.Time) (result int) {
	for now := time.Now(); birth.Before(now); result++ {
		birth = birth.AddDate(1, 0, 0)
	}
	return result
}
func calculateProjectedDeathDate(expectancy float64, birthday time.Time, age int) time.Time {
	yearsToAdd := int(expectancy)
	daysToAdd := int(daysPerYear * (expectancy - float64(yearsToAdd)))
	return birthday.AddDate(age, 0, 0).AddDate(yearsToAdd, 0, daysToAdd)
}
func countProjectedDaysRemaining(projectedDeath time.Time) (result int) {
	for d := time.Now(); d.Before(projectedDeath); result++ {
		d = d.AddDate(0, 0, 1)
	}
	return result
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
