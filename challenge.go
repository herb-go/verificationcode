package verificationcode

type ChallengeType int

const ChallengeTypeText = ChallengeType(0)
const ChallengeTypeBinary = ChallengeType(1)
const ChallengeTypeJSON = ChallengeType(2)
const ChallengeTypeJPEG = ChallengeType(3)
const ChallengeTypePNG = ChallengeType(4)

type Challenge struct {
	Status Status
	Type   ChallengeType
	Body   []byte
}

func (r *Challenge) IsSuccess() bool {
	return r.Status == StatusSuccess
}

func NewChallenge() *Challenge {
	return &Challenge{}
}
