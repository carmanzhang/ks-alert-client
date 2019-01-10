package client

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestGetDispatcherGrpcLoadBalanceClient(t *testing.T) {
	Convey("test get executor grpc client", t, func() {
		conn, err := GetDispatcherGrpcLoadBalanceClient()
		if err != nil {
			panic(err)
		}
		// sleep a few millsecond for grpc dial etcd
		clientX := pb.NewAlertConfigHandlerClient(conn)
		time.Sleep(time.Millisecond * 500)

		resp, err := clientX.CreateAlertConfig(context.Background(), &pb.AlertConfig{AlertConfigId: "alert-config-xy7k034wv2yrwz"})
		fmt.Println(resp, err)
	})
}

func TestGetDispatcherGrpcClient(t *testing.T) {
	Convey("test get executor grpc client", t, func() {
		conn, err := grpc.Dial("127.0.0.1:50000", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		// sleep a few millsecond for grpc dial etcd
		clientX := pb.NewAlertConfigHandlerClient(conn)
		time.Sleep(time.Millisecond * 500)

		resp, err := clientX.CreateAlertConfig(context.Background(), &pb.AlertConfig{AlertConfigId: "alert-config-xy7k034wv2yrwz"})
		fmt.Println(resp, err)
	})
}
