package ics

import (
	"time"
)

type Event struct {
	Class        CLASS
	Summary      string
	Description  string
	Status       EventStatus
	Geo          *GeoLocation
	Location     string
	DtEnd        time.Time
	DtStart      time.Time
	RRule        []string
	ExRule       []string
	ExDate       []time.Time
	Transparency Transparency
	Attendees    []Attendee
	Organizer    Attendee
	UID          string

	// This property defines the revision sequence number of the calendar component within a sequence of revisions.
	// https://www.kanzaki.com/docs/ical/sequence.html
	Sequence int
	// The property indicates the date/time that the instance of the iCalendar object was created.
	// https://www.kanzaki.com/docs/ical/dtstamp.html
	DtStamp time.Time
	// This property specifies the date and time that the calendar information was created by
	// the calendar user agent in the calendar store. Note: This is analogous to the creation
	// date and time for a file in the file system.
	// https://www.kanzaki.com/docs/ical/created.html
	Created time.Time
	// This property specifies the date and time that the calendar information was last
	// modified by the calendar user agent in the calendar store. Note: This is analogous to
	// the modification date and time for a file in the file system.
	// https://www.kanzaki.com/docs/ical/lastModified.html
	LastModified time.Time
}

type EventStatus string

const (
	EventStatus_CONFIRMED EventStatus = "CONFIRMED"
	EventStatus_CANCELLED EventStatus = "CANCELLED"
	EventStatus_TENTATIVE EventStatus = "TENTATIVE"
)

type CLASS string

const (
	Classification_PUBLIC       CLASS = "PUBLIC"
	Classification_PRIVATE      CLASS = "PRIVATE"
	Classification_CONFIDENTIAL CLASS = "CONFIDENTIAL"
)

type GeoLocation struct {
	Latitude  float32
	Longitude float32
}

type Transparency string

const (
	TRANSAPARENT Transparency = "TRANSPARENT"
	OPAQUE       Transparency = "OPAQUE"
)

type Attendee struct {
	CommonName   string
	EmailAddress string
	Role         Role
	PartStatus   AttendeeStatus
	//RSVP is by default NO
	CuType CalendarUserType
	Rsvp   Rsvp
}

type Role string

const (
	REQUIRED Role = "REQ-PARTICIPANT"
)

type AttendeeStatus string

const (
	AttendeeStatus_NEEDACTION AttendeeStatus = "NEEDS-ACTION"
	AttendeeStatus_TENTATIVE  AttendeeStatus = "TENTATIVE"
	AttendeeStatus_ACCEPTED   AttendeeStatus = "ACCEPTED"
	AttendeeStatus_DECLINED   AttendeeStatus = "DECLINED"
)

type CalendarUserType string

const (
	INDIVIDUAL CalendarUserType = "INDIVIDUAL"
)

// This is related to PartStat
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17
type Rsvp string

const (
	Rsvp_False Rsvp = "FALSE"
	Rsvp_True  Rsvp = "TRUE"
)
