package main

import (
    "log"
    "net/http"
    "os"
    "github.com/op/go-logging"
    "github.com/urfave/cli"
)

const LOG_FORMAT = "%{color}%{time:2006/01/02 15:04:05 -07:00 MST} [%{level:.6s}] %{shortfile} : %{color:reset}%{message}"

func runServer(c *cli.Context) {

    ///////////////// LOGGER instance
    backend := logging.NewLogBackend(os.Stderr, "", 0)
    format := logging.MustStringFormatter(LOG_FORMAT)
    backendFormatter := logging.NewBackendFormatter(backend, format)

    backendLeveled := logging.AddModuleLevel(backendFormatter)
    logLevel, err := logging.LogLevel(c.GlobalString("log-level"))
    FatalOnError(err, "Failed to set logger log level")

    backendLeveled.SetLevel(logLevel, "")

    logging.SetBackend(backendLeveled)
    logger := logging.MustGetLogger("server")
    FatalOnError(err, "Cannot create logger for level %s (%v)", c.GlobalString("log-level"))

    /////////////// HTTP HANDLERS
    handler := NewHandler(logger)
    http.Handle("/", handler)

    logger.Infof("Listening on %s...", c.GlobalString("bind-address"))
    err = http.ListenAndServe(c.GlobalString("bind-address"), nil)
    FatalOnError(err, "Failed to bind on %s: ", c.GlobalString("bind-address"))
}

func FatalOnError(err error, msg string, args ...interface{}) {
    if err != nil {
        log.Fatalf(msg, args...)
        os.Exit(1)
    }
}

func main() {
    app := cli.NewApp()

    app.Name = "HTTP Analyzer"
    app.Version = "1.0"
    app.Authors = []cli.Author{
        {
            Name:  "Michal Nezerka",
            Email: "michal.nezerka@gmail.com",
        },
    }
    app.Usage = "PIOT Adapter for VD Gps sensors"
    app.Action = runServer
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:   "bind-address,b",
            Usage:  "Listen address for API HTTP endpoint",
            Value:  "0.0.0.0:8080",
            EnvVar: "BIND_ADDRESS",
        },
        cli.StringFlag{
            Name:   "log-level,l",
            Usage:  "Logging level",
            Value:  "INFO",
            EnvVar: "LOG_LEVEL",
        },
    }

    app.Run(os.Args)
}
