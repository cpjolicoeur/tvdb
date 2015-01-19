package tvdb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	BASE_URL     = `http://thetvdb.com/api`
	TVDB_API_KEY = `5127BBF385F9FA53`
	LANGUAGE     = `en`
)

// Represents a TVDB Series
type Series struct {
	Name     string    `xml:"SeriesName"`
	Id       string    `xml:"seriesid"`
	Episodes []Episode `xml:"Episode"`
}

// Represents a TVDB Episode
type Episode struct {
	Id       string  `xml:"id"`
	Name     string  `xml:"EpisodeName"`
	Number   float32 `xml:"EpisodeNumber"`
	Season   float32 `xml:"SeasonNumber"`
	AirDate  string  `xml:"FirstAired"`
	Overview string  `xml:"Overview"`
	SeasonId int32   `xml:"seasonid"`
	SeriesId int32   `xml:"seriesid"`
}

// GetSeries returns an array of Series entries that match
// the given name parameter
func GetSeries(name string) ([]Series, error) {
	endpoint := fmt.Sprintf("%s/GetSeries.php?seriesname=%s&language=%s", BASE_URL, url.QueryEscape(name), LANGUAGE)
	resp, err := http.Get(endpoint)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	xmlBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type Data struct {
		SeriesList []Series `xml:"Series"`
	}
	v := Data{}

	err = xml.Unmarshal(xmlBody, &v)
	if err != nil {
		return nil, err
	}

	return v.SeriesList, nil
}

// GetEpisodes populates the Series with episode information
func (s *Series) GetEpisodes() error {
	if len(strings.TrimSpace(s.Id)) == 0 {
		return errors.New("A Series.Id is required to GetEpisodes")
	}

	endpoint := fmt.Sprintf("%s/%s/series/%s/all/%s.xml", BASE_URL, TVDB_API_KEY, s.Id, LANGUAGE)
	resp, err := http.Get(endpoint)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		msg := fmt.Sprintf("A series with the following ID could not be found: %s", s.Id)
		return errors.New(msg)
	}

	xmlBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type Data struct {
		EpisodeList []Episode `xml:"Episode"`
	}
	v := Data{}

	err = xml.Unmarshal(xmlBody, &v)
	if err != nil {
		return err
	}

	s.Episodes = v.EpisodeList
	return nil
}
