package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
)

func addRoutes(r *mux.Router, ds *DataSource) {
	r.HandleFunc("/match/ids", ds.HandleGetMatchIds)
	r.HandleFunc("/match/info", ds.HandleGetMatchInfo)
	r.HandleFunc("/match/participants", ds.HandleGetMatchParticipants)
}

func (d *DataSource) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	return httpadapter.New(d).CallResource(ctx, req, sender)
}

func (d *DataSource) HandleGetMatchParticipants(w http.ResponseWriter, r *http.Request) {
	matchId := r.URL.Query().Get("matchId")

	matchInfo, err := d.client.fetchMatchParticipants(matchId)

	if err != nil {
		writeResponse(w, nil, err)
		return
	}

	writeResponse(w, matchInfo, nil)
}

func (d *DataSource) HandleGetMatchIds(w http.ResponseWriter, r *http.Request) {
	matchType := r.URL.Query().Get("type")
	no := r.URL.Query().Get("no")

	matches, err := d.client.fetchMatchIds(matchType, no)

	if err != nil {
		writeResponse(w, nil, err)
		return
	}

	writeResponse(w, matches, nil)
}

func (d *DataSource) HandleGetMatchInfo(w http.ResponseWriter, r *http.Request) {
	matchId := r.URL.Query().Get("matchId")

	matchInfo, err := d.client.fetchMatchInfo(matchId)

	if err != nil {
		writeResponse(w, nil, err)
		return
	}

	writeResponse(w, matchInfo, nil)
}

func writeResponse(w http.ResponseWriter, value interface{}, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		if err != nil {
			log.DefaultLogger.Error(err.Error())
		}
		return
	}
	err = json.NewEncoder(w).Encode(value)
	if err != nil {
		log.DefaultLogger.Error(err.Error())
	}
}
