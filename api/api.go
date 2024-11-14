package api

import (
	"context"
	"log"
)

type ClientStub struct {
	plate           string
	GetUser         func(context.Context, *User) (*User, error)
	Invoke          func(context.Context, string) (string, error)
	EnterLane       func(context.Context, string) (string, error)
	ExitLane        func(context.Context, string) (string, error)
	Get             func(context.Context) (string, error)
	Post            func(context.Context, string) error
	StartTimer      func(context.Context, *TimerRequest) error
	StopTimer       func(context.Context, *TimerRequest) error
	StartReminder   func(context.Context, *ReminderRequest) error
	StopReminder    func(context.Context, *ReminderRequest) error
	IncrementAndGet func(ctx context.Context, stateKey string) (*User, error)
}

func (a *ClientStub) SetPlate(plate string) {
	a.plate = plate
}
func (a *ClientStub) Type() string {
	return "testActor2Type"
}

func (a *ClientStub) ID() string {
	log.Printf("Get ID")
	return a.plate
}

type User struct {
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

type TimerRequest struct {
	TimerName string `json:"timer_name"`
	CallBack  string `json:"call_back"`
	Duration  string `json:"duration"`
	Period    string `json:"period"`
	Data      string `json:"data"`
}

type ReminderRequest struct {
	ReminderName string `json:"reminder_name"`
	Duration     string `json:"duration"`
	Period       string `json:"period"`
	Data         string `json:"data"`
}
