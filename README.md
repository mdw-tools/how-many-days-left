# how-many-days-left

Inspired by: https://kk.org/ct2/my-life-countdown-1/

```
$ how-many-days-left -help
  -birth string
        Birth date (YYYY-MM-DD)
  -data-year int
        Data year (default 2024)
  -sex string
        Sex (one of 'f' or 'm')

$ how-many-days-left -birth 1970-01-01 -sex f
Given that you have lived 20,288 days, a lifespan of 55.58 years since your birth on Thursday, January 1 of 1970, and based on the 2024 release of medicaid's average life expectancy for females, you have 10,002 days remaining until reaching your projected lifespan of 82.93 years on Thursday, December 5 of 2052.

Event  Date         Age
------------------------------------------
Birth: 1970-01-01  0.00 (20,288 days ago)
Today: 2025-07-18 55.58
Death: 2052-12-05 82.93 (10,002 days left)
```

- 2023 data: https://www.health.ny.gov/health_care/medicaid/publications/docs/gis/23ma09_att1.pdf
- 2024 data: https://www.health.ny.gov/health_care/medicaid/publications/docs/gis/24ma05_att1.pdf