package contract

import (
	contractpb "github.com/olegvelikanov/word-of-wisdom/internal/pkg/contract/pb"
	"google.golang.org/protobuf/proto"
)

func Serialize(m *contractpb.Message) ([]byte, error) {
	return proto.Marshal(m)
}

func Deserialize(b []byte) (contractpb.Message, error) {
	result := contractpb.Message{}
	err := proto.Unmarshal(b, &result)
	if err != nil {
		return contractpb.Message{}, err
	}
	return result, nil
}
