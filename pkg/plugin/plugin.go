package plugin

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type DataSource struct {
	client     Client
	ctx        context.Context
	cancelFunc context.CancelFunc
	router     *mux.Router
}

var (
	_ backend.QueryDataHandler      = (*DataSource)(nil)
	_ backend.CheckHealthHandler    = (*DataSource)(nil)
	_ backend.CallResourceHandler   = (*DataSource)(nil)
	_ instancemgmt.InstanceDisposer = (*DataSource)(nil)
)

func NewDataSource(s backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	settings, err := GetSettings(s)
	if err != nil {
		return nil, err
	}

	client, err := NewClient(*settings)
	if err != nil {
		return nil, err
	}

	summonerData, err := client.getSummonerData()

	if err != nil {
		return nil, err
	}

	client.setPuuid(summonerData.Puuid)

	r := mux.NewRouter()
	ctx, cancelFunc := context.WithCancel(context.Background())
	ds := &DataSource{
		client:     client,
		ctx:        ctx,
		cancelFunc: cancelFunc,
		router:     r,
	}

	addRoutes(r, ds)
	return ds, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewSampleDatasource factory function.
func (d *DataSource) Dispose() {
	// Clean up datasource instance resources.
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (d *DataSource) CheckHealth(_ context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	log.DefaultLogger.Info("CheckHealth called", "request", req)

	_, err := d.client.getSummonerData()

	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: "Error connecting to Riot API",
		}, nil
	}

	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Succesfuly connected to Riot API",
	}, nil
}

// SubscribeStream is called when a client wants to connect to a stream. This callback
// allows sending the first message.
func (d *DataSource) SubscribeStream(_ context.Context, req *backend.SubscribeStreamRequest) (*backend.SubscribeStreamResponse, error) {
	log.DefaultLogger.Info("SubscribeStream called", "request", req)

	status := backend.SubscribeStreamStatusPermissionDenied
	if req.Path == "stream" {
		// Allow subscribing only on expected path.
		status = backend.SubscribeStreamStatusOK
	}
	return &backend.SubscribeStreamResponse{
		Status: status,
	}, nil
}

func (d *DataSource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.router.ServeHTTP(w, r)
}
