package tvdb

import (
	"fmt"
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetSeries(t *testing.T) {
	t.Parallel()

	Convey("It should find valid single result series", t, func() {
		series, err := GetSeries("A-Team")
		So(series[0].Name, ShouldEqual, "The A-Team")
		So(series[0].Id, ShouldEqual, "77904")
		So(err, ShouldBeNil)
	})

	Convey("It should find valid multiple result series", t, func() {
		series, err := GetSeries("Lost")
		So(len(series), ShouldEqual, 7)
		So(err, ShouldBeNil)
	})

	Convey("It should handle invalid series names", t, func() {
		series, err := GetSeries("Lasdfasdfsd")
		So(len(series), ShouldEqual, 0)
		So(err, ShouldBeNil)
	})
}

func Test_GetEpisodes(t *testing.T) {
	t.Parallel()

	Convey("It returns a full list of episodes", t, func() {
		series := Series{Name: "The A-Team", Id: "77904"}

		err := series.GetEpisodes()
		So(len(series.Episodes), ShouldEqual, 100)
		So(err, ShouldBeNil)
	})

	Convey("It requires a Series ID", t, func() {
		series := Series{Name: "The A-Team"}

		err := series.GetEpisodes()
		So(len(series.Episodes), ShouldEqual, 0)
		So(err, ShouldNotBeNil)
		So(err.Error(), shouldMatch, "Id is required")
	})

	Convey("It requires a valid Series ID", t, func() {
		series := Series{Name: "The A-Team", Id: "INVALID_77904"}

		err := series.GetEpisodes()
		So(len(series.Episodes), ShouldEqual, 0)
		So(err, ShouldNotBeNil)
		So(err.Error(), shouldMatch, "could not be found")
	})
}

func shouldMatch(actual interface{}, expected ...interface{}) string {
	found, _ := regexp.MatchString(expected[0].(string), actual.(string))
	if !found {
		return fmt.Sprintf("Query string: %v not as expected: %v", actual, expected)
	} else {
		return ""
	}
}
