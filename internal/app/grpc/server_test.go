package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ravilushqa/music-bot/api"
)

func TestServer_Greet(t *testing.T) {
	s := New(zap.NewNop(), "")

	// set up test cases
	tests := []struct {
		name string
		want string
		err  error
	}{
		{
			name: "Ravilushqa",
			want: "Hello Ravilushqa",
		},
		{
			name: "",
			err:  status.Error(codes.InvalidArgument, "name cannot be empty"),
		},
	}

	for _, tt := range tests {
		req := &api.GreetRequest{Name: tt.name}
		resp, err := s.Greet(context.Background(), req)
		require.Equal(t, tt.err, err)

		if err == nil && resp.Message != tt.want {
			t.Errorf("Greet(%v)=%v, wanted %v", tt.name, resp.Message, tt.want)
		}
	}
}
