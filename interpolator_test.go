package interpolator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type InterpolatorSuite struct {
	suite.Suite
}

func (suite *InterpolatorSuite) TestInterpolateStruct() {
	type TestStruct struct {
		FieldOne string
	}

	testStruct := TestStruct{
		FieldOne: "value-one",
	}

	template := "test {{ .FieldOne }} test"
	expected := "test value-one test"

	actual, err := Interpolate(template, testStruct)
	suite.Assert().NoError(err)
	suite.Assert().Equal(expected, actual)
}

func (suite *InterpolatorSuite) TestInterpolateMap() {

	testMap := map[string]string{
		"FieldOne": "value-one",
	}

	template := "test {{ .FieldOne }} test"
	expected := "test value-one test"

	actual, err := Interpolate(template, testMap)
	suite.Assert().NoError(err)
	suite.Assert().Equal(expected, actual)
}

func (suite *InterpolatorSuite) TestInterpolate() {
	data := map[string]string{
		"Field1": "value1",
		"field2": "value2",
	}
	tmpl := "abc {{ .field2 }} def {{.Field1}}"
	expected := "abc value2 def value1"
	actual, err := Interpolate(tmpl, data)
	suite.NoError(err)
	suite.Equal(expected, actual)
}

func (suite *InterpolatorSuite) TestInterpolateWithNonExistantKey() {
	data := map[string]string{
		"Field1": "value1",
		"field2": "value2",
	}
	tmpl := "abc {{ .field2 }} def {{.Field3}}"
	expected := "abc value2 def <no value>"
	actual, err := Interpolate(tmpl, data)
	suite.NoError(err)
	suite.Equal(expected, actual)
}

func (suite *InterpolatorSuite) TestInterpolateWithExtraMapKey() {
	data := map[string]string{
		"Field1": "value1",
		"field2": "value2",
		"Field3": "value3",
	}
	tmpl := "abc {{ .field2 }} def {{.Field1}}"
	expected := "abc value2 def value1"
	actual, err := Interpolate(tmpl, data)
	suite.NoError(err)
	suite.Equal(expected, actual)
}


func TestInterpolatorSuite(t *testing.T) {
	suite.Run(t, new(InterpolatorSuite))
}
