package main

import (
	"context"
	"encoding/base64"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"

	agent "gitlab.mobile-intra.com/cloud-next/schema/compiled/go/munic/grpc/agent"
	munic "gitlab.mobile-intra.com/cloud-next/schema/compiled/go/munic/types"
)

var evB64 = "COSCi7zW9eXVDRDkgou81vXl1Q0aACoAOg0KCwoDMTIzEgR0ZXN0QMiCgLKHr+TVDUoGCJaf88kFUgYIrKDzyQV6FwoKb2ZmZXJfbmFtZRIJbWV0cm9taWxlggEUChISCGNsb3VkLmluGAcgjLmdmwRasgEKiwEKClVOVFlQRURfTU0SfRJ7CnkKBk1NXzI0MBJvQm0KChD8DRjDv6usySsqByAIPQrXkz8yISCppvsDKNrF+wMyFa0PXpUCcys9QGQjhgSkAVTnA4cBIjosCKw7KBgwVUogNf7//////////wH+//////////8B/P//////////AQCABBBiBRDhoZhBCh8KBU1VTklDEhYKFAoSCSJPkq6Z5EJAEYE+kSdJm17AEJAP"
var state64 = "CsABEr0B5gHwPnsic3RhdGVfdmVyc2lvbiI6MS4wLCJiYXRjaF9maXJzdF9ldmVudF9pZCI6OTg3MTM3NTc5NTMyOTQzNzE2LA0qBGxhnikASGVuZ3RoIjoxLCJ0cmFjayI6eyIFPVhyZWNvcmRlZF9hdCI6MTQ5NzY1NTkwMwGNEHJ1bGVzASk8MF9ib290XzI2NS1FVkVOVAEUlGFjdGl2ZSI6dHJ1ZSwiY291bnQiOjAsInZlcnNpb24iOjJ9fX19EgsI/umeygUQuLOAHw=="

func main() {
	// Parse event
	ev := &munic.Event{}
	evBytes, _ := base64.StdEncoding.DecodeString(evB64)
	err := proto.Unmarshal(evBytes, ev)
	if err != nil {
		log.Fatal(err)
	}

	// Parse state
	state := &agent.State{}
	stateBytes, _ := base64.StdEncoding.DecodeString(state64)
	err = proto.Unmarshal(stateBytes, state)
	if err != nil {
		log.Fatal(err)
	}
	in := &agent.ProcessRequest{
		Key:    ev.GroupKey,
		Events: []*munic.Event{ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev, ev},
		State:  state,
	}

	conn, err := grpc.Dial("localhost:50070", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	cli := agent.NewProcessServiceClient(conn)

	goCount := 50
	wg := &sync.WaitGroup{}
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {

				res, err := cli.Process(ctx, in)
				if err != nil {
					log.Println(res, err)
				}
			}
		}()
	}
	wg.Wait()
}
