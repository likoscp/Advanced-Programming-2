package mail

// import (
// 	"context"

// 	pb "github.com/likoscp/finalAddProgramming/finalproto/mail"
// 	"google.golang.org/grpc"
// )

// type Client struct {
// 	conn   *grpc.ClientConn
// 	client pb.MailServiceClient
// }

// func NewMailClient(grpcAddr string) (*Client, error) {
// 	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Client{
// 		conn:   conn,
// 		client: pb.NewMailServiceClient(conn),
// 	}, nil
// }

// func (c *Client) SendMail(ctx context.Context, req *pb.SendMailRequest) (*pb.Empty, error) {
// 	return c.client.SendMail(ctx, req)
// }

// func (c *Client) Close() {
// 	c.conn.Close()
// }
