package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/scnewma/ynab-exporter/collector"
	"github.com/scnewma/ynab-exporter/version"
	log "github.com/sirupsen/logrus"
	"go.bmvs.io/ynab"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		app = kingpin.New(
			"ynab-exporter",
			"A prometheus exporter for YNAB.",
		)
		listenAddress = app.Flag(
			"web.listen-address",
			"Address on which to expose metrics and web interface.",
		).Default(":9721").Envar("LISTEN_ADDRESS").String()
		metricsPath = app.Flag(
			"web.telemetry-path",
			"Path under which to expose metrics.",
		).Default("/metrics").Envar("METRICS_PATH").String()
		logLevel = app.Flag(
			"log.level",
			"Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal]",
		).Default(log.InfoLevel.String()).Envar("LOG_LEVEL").String()
		token = app.Flag(
			"ynab.access-token",
			"YNAB personal access token to use when communicating with the YNAB API.",
		).Required().Envar("YNAB_TOKEN").String()
	)
	kingpin.Version(version.Print())
	kingpin.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	setLogLevel(*logLevel)

	log.Infof("Starting %s %s", app.Name, version.Info())

	reg := prometheus.NewPedanticRegistry()

	reg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
		collector.NewYNABCollector(ynab.NewClient(*token)),
	)

	http.Handle(*metricsPath, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>YNAB Exporter</title></head>
			<body>
			<h1>YNAB Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}

func setLogLevel(logLevel string) {
	lvl, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal("Unable to parse log level. Valid levels: [debug, info, warn, error, fatal]")
	}

	log.SetLevel(lvl)
}
