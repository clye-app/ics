package ics

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Generate(prodId string, events ...*Event) (string, error) {
	obj := &generator{
		ProdId: prodId,
		Events: []string{},
	}

	eventTmpl, err := template.New("events").Parse(vevent)
	if err != nil {
		return "", err
	}
	for _, event := range events {

		for idx := range event.Attendees {
			if event.Attendees[idx].Rsvp == "" {
				event.Attendees[idx].Rsvp = Rsvp_False
			}
		}

		e := prepareEvent(event)

		buf := &bytes.Buffer{}
		if err := eventTmpl.Execute(buf, e); err != nil {
			return "", err
		}

		obj.Events = append(obj.Events, buf.String())
	}

	buf := &bytes.Buffer{}
	icsTmpl, err := template.New("ics").Parse(ics)
	if err != nil {
		return "", err
	}
	if err := icsTmpl.Execute(buf, obj); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (event *Event) Generate(prodId string) (string, error) {
	obj := &generator{
		ProdId: prodId,
		Events: []string{},
	}

	for idx := range event.Attendees {
		if event.Attendees[idx].Rsvp == "" {
			event.Attendees[idx].Rsvp = Rsvp_False
		}
	}

	e := prepareEvent(event)

	eventTmpl, err := template.New("events").Parse(vevent)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := eventTmpl.Execute(buf, e); err != nil {
		return "", err
	}

	obj.Events = append(obj.Events, buf.String())

	buf = &bytes.Buffer{}
	icsTmpl, err := template.New("ics").Parse(ics)
	if err != nil {
		return "", err
	}
	if err := icsTmpl.Execute(buf, obj); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func prepareEvent(event *Event) *vEvent {
	e := &vEvent{
		Event:        event,
		DtStamp:      FormatDateTime(event.DtStamp, event.DtStamp),
		DtEnd:        FormatDateTime(event.DtEnd, event.DtStamp),
		DtStart:      FormatDateTime(event.DtStart, event.DtStamp),
		ExDate:       make([]string, len(event.ExDate)),
		Created:      FormatDateTime(event.Created, event.DtStamp),
		LastModified: FormatDateTime(event.LastModified, event.DtStamp),
		Description:  strings.Join(strings.Split(event.Description, "\n"), `\n`),
	}
	for i, exd := range event.ExDate {
		e.ExDate[i] = FormatDateTime(exd, event.DtStamp)
	}
	event.UID = hex.EncodeToString([]byte(e.UID))
	return e
}

type generator struct {
	ProdId string
	Events []string
}

func FormatDateTime(t time.Time, now time.Time) string {
	// default to now if not specified
	if t == (time.Time{}) {
		t = now
	}
	dt := strconv.Itoa(t.Year())

	month := strconv.Itoa(int(t.Month()))
	if len(month) < 2 {
		dt += "0"
	}
	dt += month

	day := strconv.Itoa(t.Day())
	if len(day) < 2 {
		dt += "0"
	}
	dt += day + "T"

	hour := strconv.Itoa(t.Hour())
	if len(hour) < 2 {
		dt += "0"
	}
	dt += hour

	min := strconv.Itoa(t.Minute())
	if len(min) < 2 {
		dt += "0"
	}
	dt += min

	sec := strconv.Itoa(t.Second())
	if len(sec) < 2 {
		dt += "0"
	}
	dt += sec + "Z"

	return dt
}
