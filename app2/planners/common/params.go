package common

import (
	"time"
)

type Params[T any] struct {
	Year       int
	Weekday    time.Weekday
	DeviceName string
	Sections   []string
	Hand       MainHand

	ShowFrames bool

	ArrayStretch string
	ShowLinks    bool
}

type ApplyParameterOption[T any] func(*Params[T])

func NewParams[T any](options ...ApplyParameterOption[T]) Params[T] {
	params := Params[T]{
		Weekday: time.Monday,
	}

	for _, option := range options {
		option(&params)
	}

	return params
}

func ParamWithYear[T any](year int) ApplyParameterOption[T] {
	return func(params *Params[T]) {
		params.Year = year
	}
}

func ParamWithSections[T any](sections []string) ApplyParameterOption[T] {
	return func(params *Params[T]) {
		params.Sections = sections
	}
}

func ParamWithWeekday[T any](weekday time.Weekday) ApplyParameterOption[T] {
	return func(params *Params[T]) {
		params.Weekday = weekday
	}
}

func ParamWithMainHand[T any](hand MainHand) ApplyParameterOption[T] {
	return func(params *Params[T]) {
		params.Hand = hand
	}
}

func ParamWithFrames[T any](showFrames bool) ApplyParameterOption[T] {
	return func(params *Params[T]) {
		params.ShowFrames = showFrames
	}
}

func ParamWithLinks[T any](showLinks bool) ApplyParameterOption[T] {
	return func(params *Params[T]) {
		params.ShowLinks = showLinks
	}
}

func ParamWithDeviceName[T any](deviceName string) ApplyParameterOption[T] {
	return func(params *Params[T]) {
		params.DeviceName = deviceName
	}
}
