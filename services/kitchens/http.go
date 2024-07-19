package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/Asad2730/Kitchen_Micro/generated/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func NewGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGrpcClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  3123,
			Quantity:   2,
		})

		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{CustomerID: 42})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}

	})

	log.Println("Starting server on", h.addr)
	return http.ListenAndServe(h.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
