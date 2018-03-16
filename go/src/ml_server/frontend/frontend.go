package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"ml_server/proto/learner"
	"net/http"
)

const (
	address = "ml-server-backend:50052"
)

func handler(w http.ResponseWriter, r *http.Request) {
	test := &learner.RegressRequest{
		Train: &learner.Matrix{
			DataPoint: []*learner.DataPoint{
				{
					Obs:  8.1,
					Vars: []float32{10294, 841, 10},
				}, {
					Obs:  8.1,
					Vars: []float32{10294, 841, 10},
				}, {
					Obs:  8.1,
					Vars: []float32{10294, 841, 10},
				},
			},
		},
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	c := learner.NewLearnerClient(conn)

	resp, err := c.Regress(context.Background(), test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Prediction: %f", resp.Prediction)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":50051", nil)
}
