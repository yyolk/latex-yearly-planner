package common

import (
	"time"
)

type Params struct {
	Year       int
	Weekday    time.Weekday
	DeviceName string
	Sections   []string
	Hand       MainHand

	ShowFrames bool

	ArrayStretch string
	ShowLinks    bool
}

type ApplyParameterOption func(*Params)

func NewParams(options ...ApplyParameterOption) Params {
	params := Params{
		Weekday: time.Monday,
	}

	for _, option := range options {
		option(&params)
	}

	return params
}

func ParamWithYear(year int) ApplyParameterOption {
	return func(params *Params) {
		params.Year = year
	}
}

func ParamWithSections(sections []string) ApplyParameterOption {
	return func(params *Params) {
		params.Sections = sections
	}
}

func ParamWithWeekday(weekday time.Weekday) ApplyParameterOption {
	return func(params *Params) {
		params.Weekday = weekday
	}
}

func ParamWithMainHand(hand MainHand) ApplyParameterOption {
	return func(params *Params) {
		params.Hand = hand
	}
}

func ParamWithFrames(showFrames bool) ApplyParameterOption {
	return func(params *Params) {
		params.ShowFrames = showFrames
	}
}

func ParamWithLinks(showLinks bool) ApplyParameterOption {
	return func(params *Params) {
		params.ShowLinks = showLinks
	}
}

func ParamWithDeviceName(deviceName string) ApplyParameterOption {
	return func(params *Params) {
		params.DeviceName = deviceName
	}
}
