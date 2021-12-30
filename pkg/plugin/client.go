package plugin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
)

type Client interface {
	setPuuid(puuid string)
	getSummonerData() (SummonerData, error)
	fetchMatchIds(queueType string, count string) ([]string, error)
	fetchMatchSelfInfo(matchId string) (MatchParticipantInfo, error)
	fetchTimeline(matchId string) (MatchTimeline, error)
	fetchMatchParticipants(matchId string) ([]MatchParticipantInfo, error)
}

const (
	RiotURLStub string = "https://%s.api.riotgames.com"
)

const (
	BR1  string = "BR1"
	EUN1 string = "EUN1"
	EUW1 string = "EUW1"
	JP1  string = "JP1"
	KR   string = "KR"
	LA1  string = "LA1"
	LA2  string = "LA2"
	NA1  string = "NA1"
	OC1  string = "OC1"
	TR1  string = "TR1"
	RU   string = "RU"
)

const (
	EUROPE   = "europe"
	AMERICAS = "americas"
	ASIA     = "asia"
)

func NewClient(cs ConnectionSettings) (Client, error) {
	return &LolClient{cs}, nil
}

func GetSettings(s backend.DataSourceInstanceSettings) (*ConnectionSettings, error) {
	var jsonData SettingsJsonData
	if err := json.Unmarshal(s.JSONData, &jsonData); err != nil {
		return nil, err
	}

	region := AMERICAS
	if jsonData.Platform == KR || jsonData.Platform == JP1 {
		region = ASIA
	} else if jsonData.Platform == EUN1 || jsonData.Platform == EUW1 || jsonData.Platform == TR1 || jsonData.Platform == RU {
		region = EUROPE
	}

	settings := &ConnectionSettings{
		URL:          RiotURLStub,
		SummonerName: jsonData.SummonerName,
		Platform:     jsonData.Platform,
		Region:       region,
		UID:          s.UID,
	}

	if val, ok := s.DecryptedSecureJSONData["apiToken"]; ok {
		settings.ApiToken = val
	}
	return settings, nil
}

func (c *LolClient) getSummonerData() (SummonerData, error) {
	res := SummonerData{}
	url := fmt.Sprintf("/lol/summoner/v4/summoners/by-name/%s", c.ConnectionSettings.SummonerName)
	err := c.fetchData(url, strings.ToLower(c.ConnectionSettings.Platform), "", &res)

	return res, err
}

func (c *LolClient) setPuuid(puuid string) {
	c.ConnectionSettings.PUUID = puuid
}

func (c *LolClient) fetchData(path string, urlBit string, params string, v interface{}) error {
	rsp, err := c.get(path, urlBit, params)
	if err != nil {
		return err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return errors.New(strconv.Itoa(rsp.StatusCode))
	}

	if err := json.NewDecoder(rsp.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

func (c *LolClient) fetchMatchParticipants(matchId string) ([]MatchParticipantInfo, error) {
	res := &MatchJson{}

	url := fmt.Sprintf("/lol/match/v5/matches/%s", matchId)
	err := c.fetchData(url, strings.ToLower(c.Region), "", &res)

	if err != nil {
		return []MatchParticipantInfo{}, errors.New("error returning match participant data")
	}

	var matchParticipantsInfo []MatchParticipantInfo
	for i := 0; i < len(res.Info.Participants); i++ {
		matchParticipantsInfo = append(matchParticipantsInfo, res.Info.Participants[i])
	}

	return matchParticipantsInfo, nil
}

func (c *LolClient) fetchTimeline(matchId string) (MatchTimeline, error) {
	res := MatchTimeline{}
	url := fmt.Sprintf("/lol/match/v5/matches/%s/timeline", matchId)

	err := c.fetchData(url, strings.ToLower(c.Region), "", &res)

	if err != nil {
		return MatchTimeline{}, errors.New("no timeline found")
	}

	return res, nil
}

func (c *LolClient) fetchMatchIds(queueType string, count string) ([]string, error) {
	res := []string{}
	url := fmt.Sprintf("/lol/match/v5/matches/by-puuid/%s/ids", c.PUUID)

	paramTemplate := "type=%s&start=0&count=%s"
	params := fmt.Sprintf(paramTemplate, queueType, count)

	err := c.fetchData(url, strings.ToLower(c.Region), params, &res)

	if err != nil {
		return nil, errors.New("no matches found")
	}

	return res, nil
}

func (c *LolClient) fetchMatchSelfInfo(matchId string) (MatchParticipantInfo, error) {
	res := &MatchJson{}

	url := fmt.Sprintf("/lol/match/v5/matches/%s", matchId)
	err := c.fetchData(url, strings.ToLower(c.Region), "", &res)

	if err != nil {
		return MatchParticipantInfo{}, errors.New("error returning match data")
	}

	index := -1

	for i := 0; i < len(res.Metadata.Participants); i++ {
		if res.Metadata.Participants[i] == c.PUUID {
			index = i
		}
	}

	if index == -1 {
		return MatchParticipantInfo{}, errors.New("puuid did not match with any participants")
	}

	res.Info.Participants[index].MatchId = matchId

	return res.Info.Participants[index], nil
}

func (c *LolClient) get(path string, urlBit string, params string) (*http.Response, error) {
	http, endpoint, err := c.getHttpClient(path, urlBit, params)
	if err != nil {
		return nil, err
	}
	return http.Get(endpoint)
}

func (c *LolClient) getHttpClient(path string, urlBit string, params string) (*http.Client, string, error) {
	http, err := httpclient.New(httpclient.Options{
		Timeouts: &httpclient.TimeoutOptions{
			Timeout: 30 * time.Second,
		},
		Headers: map[string]string{"X-Riot-Token": c.ConnectionSettings.ApiToken},
	})
	if err != nil {
		return nil, "", err
	}

	formattedUrl := fmt.Sprintf(c.URL, urlBit)
	base, err := url.Parse(formattedUrl)

	if err != nil {
		return nil, "", err
	}

	base.Path += path
	base.RawQuery = params
	return http, base.String(), nil
}
