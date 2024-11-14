package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dapr/go-sdk/actor"
	dapr "github.com/dapr/go-sdk/client"
	daprd "github.com/dapr/go-sdk/service/http"
)

type TestActor1 struct {
	actor.ServerImplBaseCtx
	daprClient dapr.Client
}

func (t *TestActor1) Type() string {
	return "testActor2Type"
}

func testActorFactory() actor.ServerContext {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	return &TestActor1{
		daprClient: client,
	}
}

type carState struct {
	EnterTime time.Time
	Plate     string
}

func (t *TestActor1) EnterLane(ctx context.Context, plate string) (string, error) {
	log.Printf("Car %s enter lane", plate)
	c := carState{
		EnterTime: time.Now(),
		Plate:     plate,
	}

	t.GetStateManager().Set(ctx, "plate", c)
	err := t.daprClient.RegisterActorReminder(ctx, &dapr.RegisterActorReminderRequest{
		ActorType: t.Type(),
		ActorID:   t.ID(),
		Name:      "VehiLOST",
		DueTime:   "10s",
		Data:      []byte(t.ID()),
	})
	if err != nil {
		log.Printf("Error register reminder: %s", err)
	} else {
		log.Printf("Reminder reg OK")
	}
	return plate, nil
}

func (t *TestActor1) ExitLane(ctx context.Context, plate string) (string, error) {
	log.Printf("ID %s: Car %s exit lane\n", t.ID(), plate)
	t.daprClient.UnregisterActorReminder(ctx, &dapr.UnregisterActorReminderRequest{
		ActorType: t.Type(),
		ActorID:   t.ID(),
		Name:      "VehiLOST",
	})
	c := carState{}
	t.GetStateManager().Get(ctx, "plate", &c)
	log.Printf("ID %s: Time spent in: %s\n", t.ID(), time.Since(c.EnterTime))
	return time.Since(c.EnterTime).String(), nil
}

func (t *TestActor1) ReminderCall(reminderName string, state []byte, dueTime string, period string) {
	log.Printf("receive reminder = %s state=%s duetime=%s period=%s", reminderName, string(state), dueTime, period)
}
func main() {
	s := daprd.NewService(":8080")

	s.RegisterActorImplFactoryContext(testActorFactory)

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listening: %v", err)
	}
}
