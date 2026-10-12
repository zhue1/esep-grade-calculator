package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 59, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// Part 4: The table checks the exact letter-grade thresholds, including both
// sides of the C/D boundary that determines passing status.
func TestLetterGradeBoundaries(t *testing.T) {
	testCases := []struct {
		score    int
		expected string
	}{
		{0, "F"}, {59, "F"}, {60, "D"}, {69, "D"},
		{70, "C"}, {79, "C"}, {80, "B"}, {89, "B"},
		{90, "A"}, {100, "A"},
	}

	for _, tc := range testCases {
		calculator := NewGradeCalculator()
		calculator.AddGrade("assignment", tc.score, Assignment)
		calculator.AddGrade("exam", tc.score, Exam)
		calculator.AddGrade("essay", tc.score, Essay)
		if got := calculator.GetFinalGrade(); got != tc.expected {
			t.Errorf("uniform score %d: expected %s, got %s", tc.score, tc.expected, got)
		}
	}
}

func TestCategoryWeightsAndEssayIndependence(t *testing.T) {
	calculator := NewGradeCalculator()
	calculator.AddGrade("assignment", 100, Assignment)
	calculator.AddGrade("exam", 100, Exam)
	calculator.AddGrade("essay", 0, Essay)

	// 100(0.50) + 100(0.35) + 0(0.15) = 85 -> B.
	if got := calculator.GetFinalGrade(); got != "B" {
		t.Errorf("expected B when essay score is zero; got %s", got)
	}
}

func TestAverageMultipleGradesWithinEachCategory(t *testing.T) {
	calculator := NewGradeCalculator()
	calculator.AddGrade("assignment 1", 100, Assignment)
	calculator.AddGrade("assignment 2", 0, Assignment)
	calculator.AddGrade("exam 1", 100, Exam)
	calculator.AddGrade("exam 2", 100, Exam)
	calculator.AddGrade("essay 1", 100, Essay)
	calculator.AddGrade("essay 2", 100, Essay)

	// Assignments average 50; exams and essays average 100 -> 75 -> C.
	if got := calculator.GetFinalGrade(); got != "C" {
		t.Errorf("expected C when assignment scores average 50; got %s", got)
	}
}

func TestNoGradesAndMissingCategories(t *testing.T) {
	t.Run("no grades", func(t *testing.T) {
		calculator := NewGradeCalculator()
		if got := calculator.GetFinalGrade(); got != "F" {
			t.Errorf("expected F for no recorded grades; got %s", got)
		}
	})

	t.Run("only assignments", func(t *testing.T) {
		calculator := NewGradeCalculator()
		calculator.AddGrade("assignment", 100, Assignment)
		// Missing exams and essays contribute zero under the assumed policy.
		if got := calculator.GetFinalGrade(); got != "F" {
			t.Errorf("expected F when only 50%% category is present; got %s", got)
		}
	})
}

func TestGradeTypeString(t *testing.T) {
	for _, tc := range []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
		{GradeType(-1), ""},
	} {
		if got := tc.gradeType.String(); got != tc.expected {
			t.Errorf("type %d: expected %q, got %q", tc.gradeType, tc.expected, got)
		}
	}
}

// Part 7: Specify the new behavior before adding its implementation.
// Existing calls to NewGradeCalculator and AddGrade remain unchanged.
func TestPassFailBoundary(t *testing.T) {
	for _, tc := range []struct {
		score    int
		expected string
	}{
		{0, "Fail"}, {69, "Fail"}, {70, "Pass"}, {100, "Pass"},
	} {
		calculator := NewPassFailGradeCalculator()
		calculator.AddGrade("assignment", tc.score, Assignment)
		calculator.AddGrade("exam", tc.score, Exam)
		calculator.AddGrade("essay", tc.score, Essay)
		if got := calculator.GetFinalGrade(); got != tc.expected {
			t.Errorf("Pass/Fail score %d: expected %s, got %s", tc.score, tc.expected, got)
		}
	}
}

func TestPassFailUsesWeightedGrade(t *testing.T) {
	calculator := NewPassFailGradeCalculator()
	calculator.AddGrade("assignment", 100, Assignment)
	calculator.AddGrade("exam", 50, Exam)
	calculator.AddGrade("essay", 0, Essay)

	// 100(0.50) + 50(0.35) + 0(0.15) = 67.5 -> Fail.
	if got := calculator.GetFinalGrade(); got != "Fail" {
		t.Errorf("expected Fail for weighted score 67.5; got %s", got)
	}
}

func TestPassFailDoesNotChangeLetterGradeMode(t *testing.T) {
	calculator := NewGradeCalculator()
	calculator.AddGrade("assignment", 70, Assignment)
	calculator.AddGrade("exam", 70, Exam)
	calculator.AddGrade("essay", 70, Essay)
	if got := calculator.GetFinalGrade(); got != "C" {
		t.Errorf("existing constructor should preserve letter grades; got %s", got)
	}
}
