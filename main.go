package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	age := flag.Int("age", -1, "Age in years.")
	sex := flag.String("sex", "", "Sex (one of 'f' or 'm')")
	flag.Parse()

	if *age < 0 {
		flag.PrintDefaults()
		log.Fatal("Failed to parse date of birth.")
	}
	if *sex != "m" && *sex != "f" {
		flag.PrintDefaults()
		log.Fatal("Failed to parse sex.")
	}

	expected := getLifeExpectancy(*age, *sex)
	yearsToAdd := int(expected)
	daysToAdd := int(365.0 * (expected - float64(yearsToAdd)))
	death := time.Now().AddDate(yearsToAdd, 0, daysToAdd)

	daysRemaining := 0

	for d := time.Now(); d.Before(death); d = d.AddDate(0, 0, 1) {
		daysRemaining++
	}

	fmt.Println(daysRemaining)
}

func getLifeExpectancy(age int, sex string) float64 {
	reader := csv.NewReader(strings.NewReader(CSVLifeExpectancy))
	reader.Comma = '\t'
	reader.FieldsPerRecord = 3

	all, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range all {
		if record[0] == fmt.Sprint(age) {
			if sex == "f" {
				return parseFloat(record[2])
			} else {
				return parseFloat(record[1])
			}
		}
	}

	log.Fatal("Provided age not found:", age)
	return 0
}

func parseFloat(s string) float64 {
	parsed, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("Invalid float64: '%s' Parse Error: [%v]", s, err)
	}
	return parsed
}

// See https://www.health.ny.gov/health_care/medicaid/publications/docs/adm/06adm-5att8.pdf
const CSVLifeExpectancy = `Age	Male	Female
0	74.14	79.45
1	73.7	78.94
2	72.74	77.97
3	71.77	77
4	70.79	76.01
5	69.81	75.03
6	68.82	74.04
7	67.83	73.05
8	66.84	72.06
9	65.85	71.07
10	64.86	70.08
11	63.87	69.09
12	62.88	68.09
13	61.89	67.1
14	60.91	66.11
15	59.93	65.13
16	58.97	64.15
17	58.02	63.17
18	57.07	62.2
19	56.14	61.22
20	55.2	60.25
21	54.27	59.28
22	53.35	58.3
23	52.42	57.33
24	51.5	56.36
25	50.57	55.39
26	49.64	54.41
27	48.71	53.44
28	47.77	52.47
29	46.84	51.5
30	45.9	50.53
31	44.96	49.56
32	44.03	48.6
33	43.09	47.63
34	42.16	46.67
35	41.23	45.71
36	40.3	44.76
37	39.38	43.8
38	38.46	42.86
39	37.55	41.91
40	36.64	40.97
41	35.73	40.03
42	34.83	39.09
43	33.94	38.16
44	33.05	37.23
45	32.16	36.31
46	31.29	35.39
47	30.42	34.47
48	29.56	33.56
49	28.7	32.65
50	27.85	31.75
51	27	30.85
52	26.16	29.95
53	25.32	29.07
54	24.5	28.18
55	23.68	27.31
56	22.86	26.44
57	22.06	25.58
58	21.27	24.73
59	20.49	23.89
60	19.72	23
61	18.96	22.24
62	18.21	21.43
63	17.48	20.63
64	16.76	19.84
65	16.05	19.06
66	15.36	18.3
67	14.68	17.54
68	14.02	16.8
69	13.38	16.07
70	12.75	15.35
71	12.13	14.65
72	11.53	13.96
73	10.95	13.28
74	10.38	12.62
75	9.83	11.97
76	9.29	11.33
77	8.77	10.71
78	8.27	10.11
79	7.78	9.52
80	7.31	8.95
81	6.85	8.4
82	6.42	7.87
83	6	7.36
84	5.61	6.88
85	5.24	6.42
86	4.89	5.98
87	4.56	5.56
88	4.25	5.17
89	3.97	4.81
90	3.7	4.47
91	3.45	4.15
92	3.22	3.86
93	3.01	3.59
94	2.82	3.35
95	2.64	3.13
96	2.49	2.93
97	2.35	2.75
98	2.22	2.58
99	2.11	2.43
100	2	2.29
101	1.89	2.15
102	1.79	2.02
103	1.69	1.89
104	1.59	1.77
105	1.5	1.66
106	1.41	1.55
107	1.33	1.44
108	1.25	1.34
109	1.17	1.25
110	1.1	1.16
111	1.03	1.07
112	0.96	0.99
113	0.89	0.91
114	0.83	0.84
115	0.77	0.77
116	0.71	0.71
117	0.66	0.66
118	0.61	0.61
119	0.56	0.56
`
