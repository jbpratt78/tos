package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"google.golang.org/grpc"
)

var (
	listen = flag.String("listen", ":50051", "listen address")
)

type server struct{}

func (*server) GetMenu(ctx context.Context, empty *empty.Empty) (*mookiespb.Menu, error) {
	fmt.Println("Menu function was invoked")
	res := &mookiespb.Menu{
		Items: []*mookiespb.Item{
			{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Category: "Sandwich"},
			{Name: "Regular Smoked Pulled Pork", Id: 2, Price: 395, Category: "Sandwich"},
			{Name: "Large Smoked Chicken Breast", Id: 3, Price: 495, Category: "Sandwich"},
			{Name: "Regular Smoked Chicken Breast", Id: 4, Price: 395, Category: "Sandwich"},
			{Name: "Large Hamburger", Id: 5, Price: 495, Category: "Sandwich"},
			{Name: "Hamburger", Id: 6, Price: 395, Category: "Sandwich"},
			{Name: "Large Cheeseburger", Id: 7, Price: 555, Category: "Sandwich"},
			{Name: "Cheeseburger", Id: 8, Price: 425, Category: "Sandwich"},
			{Name: "Large Hamburger", Id: 9, Price: 495, Category: "Sandwich"},
			{Name: "Large Hamburger", Id: 10, Price: 495, Category: "Sandwich"},
		},
	}
	return res, nil
}

func (*server) SubmitOrder(ctx context.Context,
	req *mookiespb.SubmitOrderRequest) (*mookiespb.SubmitOrderResponse, error) {

	fmt.Printf("SubmitOrder function was invoked with %v\n", req)
	// req.Order needs to get sent over a channel to SubscribeToOrders
	res := &mookiespb.SubmitOrderResponse{
		Result: "Order was received.",
	}
	return res, nil
}

func (*server) SubscribeToOrders(req *mookiespb.ProcessOrderRequest,
	stream mookiespb.OrderService_SubscribeToOrdersServer) error {

	fmt.Printf("SubscribeToOrders function was invoked with %v\n", req)

	// client subscribes to receive orders
	// loop and on each receive of SubmitOrder, need to
	// call stream.Send(order)

	return nil

}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening on %q...\n", *listen)

	s := grpc.NewServer()
	mookiespb.RegisterMenuServiceServer(s, &server{})
	mookiespb.RegisterOrderServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
