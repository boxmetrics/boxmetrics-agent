package agent

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// EventCode type
type EventCode int

const (
	_ EventCode = iota
	// Info event
	Info
	// Script event
	Script
	// Command event
	Command
)

// UnmarshalJSON parse JSON
func (ec *EventCode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		return errors.New("EventCode not support")
	case "info":
		*ec = Info
	case "script":
		*ec = Script
	case "command":
		*ec = Command
	}

	return nil
}

// MarshalJSON return JSON
func (ec EventCode) MarshalJSON() ([]byte, error) {
	var s string
	switch ec {
	default:
		s = ""
	case Info:
		s = "info"
	case Script:
		s = "script"
	case Command:
		s = "command"
	}

	return json.Marshal(s)
}

type options struct {
	Args []string `json:"args"`
	Env  []string `json:"env"`
	Dir  string   `json:"pwd"`
	Pid  int32    `json:"pid"`
}

type event struct {
	Type    EventCode `json:"type"`
	Value   string    `json:"value"`
	Options options   `json:"options"`
	Format  bool      `json:"format"`
}

func (e *event) validate() error {

	Log.WithField("event", e).Debug("event receive")
	if cmp.Equal(e, new(event)) {
		return errors.New("Empty event")
	}

	if !Config.GetBool("jwt_auth") {
		return nil
	}
	// TODO: Ajouter authentifaction par token
	return nil
}

func (e event) String() string {
	s, _ := json.Marshal(e)
	return string(s)
}

// Status type
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SuccessStatus return on success
var SuccessStatus = Status{200, "Request succeed"}

// ErrorStatus return on success
var ErrorStatus = Status{400, "Request failed"}

type response struct {
	Event     event       `json:"event"`
	Data      interface{} `json:"data"`
	StartDate time.Time   `json:"startDate"`
	EndDate   time.Time   `json:"endDate"`
	Duration  string      `json:"duration"`
	Status    Status      `json:"status"`
	Error     error       `json:"error"`
}

func (r response) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

func (r response) HasError() bool {
	return r.Error != nil
}

func (r *response) SetStatus(s Status) {
	r.Status = s
}

func (r *response) SetError(err error) {
	r.Error = err
	r.SetStatus(ErrorStatus)
}

func (r *response) SetData(d interface{}) {
	r.Data = d
	r.SetStatus(SuccessStatus)
}

func newResponse(e event) *response {
	date := time.Now()

	res := response{e, nil, date, time.Time{}, "", Status{}, nil}

	return &res
}

// CreateServer create websocket server
func CreateServer() {
	http.HandleFunc("/ws/v1", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		if err != nil {
			Log.WithField("error", err).Error("websocket error")
		}

		for {
			e := event{Format: true}

			err := conn.ReadJSON(&e)

			r := newResponse(e)

			if err != nil {
				Log.WithField("error", err).Error("cannot read json message")
				r.SetError(errors.Convert(err))
			} else {
				err := e.validate()

				if err != nil {
					Log.WithField("error", err).Error("invalid event")
					r.SetError(err)
				} else {
					// Log message
					logfields := logrus.Fields{"remote": conn.RemoteAddr(), "event": e}
					Log.WithFields(logfields).Info("receive")

					// Dispatch action and return response
					data, err := dispatchEvent(e)

					r.EndDate = time.Now()
					r.Duration = r.EndDate.Sub(r.StartDate).String()

					if err != nil {
						Log.WithField("error", err).Error("cannot dispatch request")
						r.SetError(err)
					} else {
						Log.Debug(data)
						r.SetData(data)
					}
				}
			}

			Log.WithField("response", r).Debug("response before send to client")

			// Write response to client
			if err = conn.WriteJSON(r); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/websocket.html")
	})

	var httpErr error
	protocol := Config.GetString("protocol")
	host := Config.GetString("host")
	port := Config.GetInt(strings.Join([]string{protocol, "_port"}, ""))
	addr := strings.Join([]string{host, ":", strconv.Itoa(port)}, "")
	url := strings.Join([]string{protocol, "://", addr}, "")

	logfields := logrus.Fields{"host": host, "port": port, "url": url}
	Log.WithFields(logfields).Info("server started")

	if protocol == "https" {
		crt := Config.GetString("ssl_crt")
		key := Config.GetString("ssl_key")
		if _, err := os.Stat(crt); err != nil {
			Log.WithField("error", err).Fatal("could not find certificate file")
		}
		if _, err := os.Stat(key); err != nil {
			Log.WithField("error", err).Fatal("could not find key file")
		}
		httpErr = http.ListenAndServeTLS(addr, crt, key, nil)

	} else {
		httpErr = http.ListenAndServe(addr, nil)
	}

	if httpErr != nil {
		Log.WithField("error", httpErr).Fatal("listener fatal error")
	}
}
