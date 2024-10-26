package main

import (
	"context"
	pb "llm-simulation/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWorldServerClient(conn)

	// Register agent
	agent := &pb.Agent{
		Id:   "agent1",
		Type: "basic",
		Attributes: map[string]string{
			"name": "TestAgent",
		},
	}

	_, err = client.RegisterAgent(context.Background(), agent)
	if err != nil {
		log.Fatalf("failed to register: %v", err)
	}

	// Start streaming world state
	stream, err := client.StreamWorldState(context.Background(), agent)
	if err != nil {
		log.Fatalf("failed to start streaming: %v", err)
	}

	// Handle world state updates
	go func() {
		for {
			state, err := stream.Recv()
			if err != nil {
				log.Printf("stream error: %v", err)
				return
			}
			log.Printf("Received state update: %v", state)
		}
	}()

	// Periodically submit actions
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		action := &pb.Action{
			AgentId:    agent.Id,
			ActionType: "move",
			Parameters: map[string]string{
				"direction": "north",
			},
		}

		_, err := client.SubmitAction(context.Background(), action)
		if err != nil {
			log.Printf("failed to submit action: %v", err)
		}
	}
}
