package main

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/7/2017.
 */

import (
	"net/http"
	"time"
	"os"
	"github.com/Sirupsen/logrus"

)

//Create a new instances of the logger
var log2StdOut = logrus.New()
var log2file = logrus.New()

func init()  {
	//Set formatters
	log2StdOutFormatter := new(logrus.TextFormatter)
	log2StdOutFormatter.TimestampFormat = "02-01-2006 15:04:05"
	log2StdOutFormatter.FullTimestamp = true
	log2fileFormatter := new(logrus.TextFormatter)
	log2fileFormatter.TimestampFormat = "02-01-2006 15:04:05"
	log2fileFormatter.FullTimestamp = true
	log2StdOut.Formatter = log2StdOutFormatter
	log2file.Formatter = log2fileFormatter

	//Set logging level
	log2StdOut.SetLevel(logrus.DebugLevel)
	log2file.SetLevel(logrus.DebugLevel)

	//Set Output for each logger instance
	file, err := os.OpenFile("logrus.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err == nil {
		log2file.Out = file

	} else {
		log2file.Error("Failed to log to file, using default stderr")

	}

	log2StdOut.Out = os.Stdout

}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log2StdOut.Info(r.Method, " ", r.RequestURI, " ", name, " ", time.Since(start))
		log2file.Info(r.Method, " ", r.RequestURI, " ", name, " ", time.Since(start))

	})

}

