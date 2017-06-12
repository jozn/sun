package models

import (
	rpc "ms/sun/play/rpc/mesg"
	"net"

	"google.golang.org/grpc/reflection"

	"fmt"
	"log"
	"ms/sun/helper"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":4000"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

var cert = "server.crt"

// SayHello implements helloworld.GreeterServer
func (s *server) GetUsers(context.Context, *rpc.UserReq) (*rpc.UserRes, error) {
	return &rpc.UserRes{Users: []*rpc.User{
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Name: "da hsdkjhsd sdsdjhsdkjhsd sdjhkasdk sdjs ksjdh sskjhsdjhjhkj sتانتانتاسنتشساسنیتاستناسین سیسی sdas", Email: "sdsadsd jshdbjsd sdjsdbkjsds sdkjsdjdshsd sdj"},
		&rpc.User{Email: "dddsd@sd"}}}, nil
}

func SERVER_GRPC() {
	_ = cert
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Create the TLS credentials
	//creds, err := credentials.NewServerTLSFromFile(cert, "server.key")
	//if err != nil {
	//    return fmt.Errorf("could not load TLS keys: %s", err)
	//}

	//s := grpc.NewServer(grpc.Creds(creds))

	rpc.RegisterUserServiceServer(s, &server{})
	//pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Client_GRPC() {

	const (
		address     = "localhost:4000"
		defaultName = "world"
	)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	c := rpc.NewUserServiceClient(conn) //pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
		_ = name
	}
	r, err := c.GetUsers(context.Background(), &rpc.UserReq{Limit: 50, Page: 69}) //&pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Users)

	go func() {

		time.Sleep(time.Microsecond * 1000)

		arr := make([]*rpc.UserRes, 0, 1000000)
		fmt.Println("start")

		t := helper.TimeNowMs()
		for i := 0; i < 1000000; i++ {
			a, err := c.GetUsers(context.Background(), &rpc.UserReq{Limit: 50, Page: 69})
			if err != nil {
				log.Fatalf("could not greet: %v  %d", err, i)
			}
			//if err == nil{
			arr = append(arr, a)
			//}
			if i%1000 == 0 {
				fmt.Printf("1000 of grpc %d : %d , len = %d\n", i, helper.TimeNowMs()-t, len(arr))
			}
		}

		fmt.Println("100000 of grpc :", helper.TimeNowMs()-t)
		fmt.Println("len :", len(arr))

	}()

}
