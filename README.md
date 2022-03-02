# Date Diff

The purpose of this program is to calculate the number of the days between two dates. The boundaries of the interval are excluded from the calculation.

For example, between 1st and 3rd of January of this year, there is only 1 day.

# Usage

Start the program with the following command:

`go run .`

The program asks for the two dates. Both must be given in `<day>/<month>/<year>` format and they must be valid dates of the 1900-2999 years.

* In case of entering valid dates, the progam writes a message to the standard output like this:

```
The difference between 3/1/1989 and 21/12/1989 days is 351 day(s).
```

* In case of providing invalid data, the program stops with meaningful error messages.
  