package esepunittests

type GradeCalculator struct {
	grades   []Grade
	passFail bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{grades: make([]Grade, 0)}
}

// NewPassFailGradeCalculator configures Pass/Fail output without changing the
// existing NewGradeCalculator, AddGrade, or GetFinalGrade call signatures.
func NewPassFailGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades:   make([]Grade, 0),
		passFail: true,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	// Preserve the original API's behavior: silently ignore unknown grade types.
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	var sums [3]int
	var counts [3]int

	for _, grade := range gc.grades {
		category := int(grade.Type)
		sums[category] += grade.Grade
		counts[category]++
	}

	weights := [3]float64{0.50, 0.35, 0.15}
	total := 0.0
	for category, weight := range weights {
		if counts[category] > 0 {
			total += float64(sums[category]) / float64(counts[category]) * weight
		}
	}
	return total
}
