package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data/framestruct"
)

const FIVE_HOURS_IN_MILLISECONDS = 18000000

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (d *DataSource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	response := backend.NewQueryDataResponse()

	for _, q := range req.Queries {
		var qm queryModel
		err := json.Unmarshal(q.JSON, &qm)

		if err != nil {
			response.Responses[q.RefID] = backend.DataResponse{Frames: nil, Error: err}
			continue
		}

		res, err := d.client.fetchTimeline(qm.MatchId)

		if err != nil {
			response.Responses[q.RefID] = backend.DataResponse{Frames: nil, Error: err}
			continue
		}

		var startTimestamp int64

		if res.Info.Frames[0].Events[0].Type == "PAUSE_END" && !qm.NormalizeTimerange {
			startTimestamp = res.Info.Frames[0].Events[0].RealTimestamp
		} else {
			startTimestamp = (time.Now().UnixNano() / int64(time.Millisecond)) - FIVE_HOURS_IN_MILLISECONDS
		}

		var participantIndex int
		if len(res.Info.Frames) > 0 {
			for i := range res.Info.Participants {
				if res.Info.Participants[i].PUUID == qm.Player {
					participantIndex = res.Info.Participants[i].ParticipantId
					continue
				}
			}
		}

		var values []TimeLineDataFrameValues
		for i := range res.Info.Frames {
			frame, err := getParticipantFrame(participantIndex, res.Info.Frames[i])

			if err != nil {
				response.Responses[q.RefID] = backend.DataResponse{Frames: nil, Error: err}
				return response, nil
			}

			dfValue := TimeLineDataFrameValues{}
			gameTimestampInMilliseconds := startTimestamp + res.Info.Frames[i].Timestamp
			dfValue.Timestamp = time.Unix(0, gameTimestampInMilliseconds*int64(time.Millisecond))

			dfValue.Value = float64(getDataframeValue(qm.TimelineData, frame))

			values = append(values, dfValue)
		}

		frames, err := framestruct.ToDataFrames(qm.TimelineData, values)

		if err != nil {
			response.Responses[q.RefID] = backend.DataResponse{Frames: nil, Error: err}
			continue
		}

		response.Responses[q.RefID] = backend.DataResponse{Frames: frames, Error: err}
	}

	return response, nil
}

func getDataframeValue(timelineData string, pFrame ParticipantFrame) float64 {
	val := reflect.ValueOf(&pFrame).Elem()
	field := val.FieldByName(strings.Title(timelineData))

	if field != (reflect.Value{}) {
		if timelineData == "goldPerSecond" {
			return reflect.ValueOf(field.Interface()).Float()
		}

		return float64(reflect.ValueOf(field.Interface()).Int())
	}

	val = reflect.ValueOf(&pFrame.ChampionStats).Elem()
	field = val.FieldByName(strings.Title(timelineData))

	if field != (reflect.Value{}) {
		return float64(reflect.ValueOf(field.Interface()).Int())
	}

	val = reflect.ValueOf(&pFrame.DamageStats).Elem()
	field = val.FieldByName(strings.Title(timelineData))

	if field != (reflect.Value{}) {
		return float64(reflect.ValueOf(field.Interface()).Int())
	}

	return 0
}

func getParticipantFrame(participantIndex int, frame Frame) (ParticipantFrame, error) {
	switch participantIndex {
	case 1:
		return frame.ParticipantFrames.ParticipantOneFrame, nil
	case 2:
		return frame.ParticipantFrames.ParticipantTwoFrame, nil
	case 3:
		return frame.ParticipantFrames.ParticipantThreeFrame, nil
	case 4:
		return frame.ParticipantFrames.ParticipantFourFrame, nil
	case 5:
		return frame.ParticipantFrames.ParticipantFiveFrame, nil
	case 6:
		return frame.ParticipantFrames.ParticipantSixFrame, nil
	case 7:
		return frame.ParticipantFrames.ParticipantSevenFrame, nil
	case 8:
		return frame.ParticipantFrames.ParticipantEightFrame, nil
	case 9:
		return frame.ParticipantFrames.ParticipantNineFrame, nil
	case 10:
		return frame.ParticipantFrames.ParticipantTenFrame, nil
	default:
		return ParticipantFrame{}, errors.New("no such participant frame")
	}
}
