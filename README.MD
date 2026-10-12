# ESEP Grade Calculator

A small Go library for calculating a course's weighted final letter grade from assignments, exams, and essays.

## Requirements and project layout

- Go 1.24.4 or a compatible later Go version (see `grade-calculator/go.mod`).
- `grade-calculator/go.mod`: the Go module declaration, `esep/grade-calculator`.
- `grade-calculator/grade_calculator.go`: grade types, in-memory storage, grade calculation, and letter-grade conversion.
- `grade-calculator/grade_calculator_test.go`: the original unit tests.

The Go module is in the **`grade-calculator/`** directory (with a hyphen), not the repository root.

## Running the tests

```sh
cd grade-calculator
go test -v ./...
```

To collect statement coverage and view the coverage report:

```sh
cd grade-calculator
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

`coverage.out` is a generated diagnostic file and should not be committed.

## Intended API and grading policy

```go
calculator := NewGradeCalculator()
calculator.AddGrade("Assignment 1", 92, Assignment)
calculator.AddGrade("Midterm", 85, Exam)
calculator.AddGrade("Essay 1", 88, Essay)
finalLetter := calculator.GetFinalGrade()
```

All source and test files declare the `esepunittests` Go package.

Within each category, grades are equally weighted. The intended course weighting is 50% assignments, 35% exams, and 15% essays. The weighted score is converted to A (90+), B (80+), C (70+), D (60+), or F (below 60); C or above passes. Scores are expected to be integers from 0 to 100.

## Known issues in the inherited version

At the time of the Part 2 documentation-only commit, the supplied implementation incorrectly averages slice indices rather than scores, uses the exam slice twice in the weighted calculation, and divides by zero on an empty category. The supplied `TestGetGradeF` also asserts an F for inputs that should produce an A. These issues were deliberately left unchanged in the Part 2 **documentation-only** commit and corrected separately in Part 3.

## Pass/Fail mode (Part 7 extra credit)

The original public methods remain compatible. Use `NewGradeCalculator()` to return A/B/C/D/F, or `NewPassFailGradeCalculator()` to return `"Pass"` for a weighted score of 70 or above and `"Fail"` otherwise. This uses the same category weights and does not alter the grade-collection interface.

The specification does not explicitly choose Pass/Fail output strings; this implementation adopts the full strings `"Pass"` and `"Fail"` as a documented assumption.
