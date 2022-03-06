# Date Diff

The purpose of this program is to calculate the number of the days between two dates. The boundaries of the interval are excluded from the calculation.

For example, between 1st and 3rd of January of this year, there is only 1 day.

# Install

This program is distributed in source code only. 
In order to run it, you need to install Go compiler. 
Follow the instructions on [this](https://go.dev/doc/install) page.

After installing Go, you can install the program with the following command:

```
go install github.com/halmai/datediff@latest
```

# Usage

Start the program with the following command:

`datediff`

Note: If it doesn't start, check whether your `$GOPATH` directory, where the program has been installed into, is in your `$PATH`. 
The `go env GOPATH` command tells the installation directory. Add this value to your `$PATH` if needed.

The program asks for the two dates. Both must be given in `<day>/<month>/<year>` format and they must be valid dates of the 1900-2999 years.

* In case of entering valid dates, the program writes a message to the standard output like this:

```
The difference between the given dates is 351 days.
```

* In case of providing invalid data, the program stops with a meaningful error message.

# Contribution

If you want to contribute to the project, please raise pull requests on GitHub.
