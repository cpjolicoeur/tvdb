package tvdb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASE_URL     = `http://thetvdb.com/api`
	TVDB_API_KEY = `5127BBF385F9FA53`
	LANGUAGE     = `en`
)

type SeriesPacket struct {
	SeriesList []Series `xml:"Series"`
}

type EpisodesPacket struct {
	EpisodeList []Episode `xml:"Episode"`
}

type Series struct {
	Name     string    `xml:"SeriesName"`
	Id       string    `xml:"seriesid"`
	Episodes []Episode `xml:"Episode"`
}

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
	endpoint := fmt.Sprintf("%s/GetSeries.php?seriesname=%s&language=%s", BASE_URL, name, LANGUAGE)
	resp, err := http.Get(endpoint)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var sPacket SeriesPacket
	xmlBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(xmlBody, &sPacket)
	if err != nil {
		return nil, err
	}

	return sPacket.SeriesList, nil
}

// GetEpisodes populates the Series with episode information
func (s *Series) GetEpisodes() error {
	endpoint := fmt.Sprintf("%s/%s/series/%s/all/%s.xml", BASE_URL, TVDB_API_KEY, s.Id, LANGUAGE)
	resp, err := http.Get(endpoint)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	var ePacket EpisodesPacket
	xmlBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(xmlBody, &ePacket)
	if err != nil {
		return err
	}

	s.Episodes = ePacket.EpisodeList
	return nil
}
