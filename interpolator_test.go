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

	actual, err := interpolate(template, testStruct)
	suite.Assert().NoError(err)
	suite.Assert().Equal(expected, actual)
}

func (suite *InterpolatorSuite) TestInterpolateMap() {

	testMap := map[string]string{
		"FieldOne": "value-one",
	}

	template := "test {{ .FieldOne }} test"
	expected := "test value-one test"

	actual, err := interpolate(template, testMap)
	suite.Assert().NoError(err)
	suite.Assert().Equal(expected, actual)
}

func TestInterpolatorSuite(t *testing.T) {
	suite.Run(t, new(InterpolatorSuite))
}
