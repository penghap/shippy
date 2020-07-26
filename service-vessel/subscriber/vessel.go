package subscriber

import (
	pb "github.com/penghap/shippy/service-vessel/proto/vessel"
	"golang.org/x/net/context"

	"fmt"
)

func Handler(ctx context.Context, msg *pb.Response) error {
	fmt.Printf("Received message: %t \n", msg.Created)
	return nil
}
