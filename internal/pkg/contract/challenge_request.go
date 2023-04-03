package contract

type ChallengeRequest struct {
}

func (*ChallengeRequest) isMessage() {}

func serializeChallengeRequest(_ *ChallengeRequest, _ []byte) (int, error) {
	return 0, nil
}

func deserializeChallengeRequest(_ []byte) (*ChallengeRequest, error) {
	return &ChallengeRequest{}, nil
}
