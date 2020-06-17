package truth

struct Truth {
  t testing.T
}

func (t *Truth) assertThatInt(actual int) *IntSubject {
  return &IntSubject{t, actual}
}

type Subject struct {
  t testing.T
}

type IntSubject struct {
  Subject
  actual int
}

func (s *IntSubject) isEqualTo(expected int) {
  ComparisonResult difference = compareForEquality(expected)
  if !difference.valuesAreEqual() {
    s.failEqualityCheck(EqualityCheck.EQUAL, expected, difference)
  }
}

type ComparisonResult interface {
  valuesAreEqual bool
  facts Fact[]
}

type Fact struct {
  key string
  value string
}

/*
func (t *Truth) assertThat(actual float64) FloatSubject {
  return FloatSubject{t, actual}
}

type FloatSubject struct {
  Subject
  actual float64
}
func (s *FloatSubject) isEqualTo(expected float64) {
  ComparisonResult difference = compareForEquality(expected)
  if !difference.valuesAreEqual() {
    s.failEqualityCheck(EqualityCheck.EQUAL, expected, difference)
  }
}

func compareForEquality(actual, expected float64) {
  return ComparisonResult.fromEqualsResult(Float.compare(actual, expected) == 0)
}
*/
