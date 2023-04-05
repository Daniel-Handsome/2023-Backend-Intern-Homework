package errorAssert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsGrpcInternalError(t *testing.T, err error) {
	assert.Equal(t, codes.Internal, status.Convert(err).Code())
}
