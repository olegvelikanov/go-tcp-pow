package contract

import (
	contractpb "github.com/olegvelikanov/word-of-wisdom/internal/pkg/contract/pb"
	"github.com/olegvelikanov/word-of-wisdom/internal/pkg/pow"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvPuzzleFromPb(p *contractpb.Puzzle) *pow.Puzzle {
	return &pow.Puzzle{
		Timestamp:        p.Timestamp.AsTime(),
		CoveredPreImage:  p.CoveredPreImage,
		CoveredBitsCount: int(p.CoveredBitsCount),
		Hash:             p.Hash,
	}
}

func ConvPuzzleToPb(p *pow.Puzzle) *contractpb.Puzzle {
	return &contractpb.Puzzle{
		Timestamp:        timestamppb.New(p.Timestamp),
		CoveredPreImage:  p.CoveredPreImage,
		CoveredBitsCount: int32(p.CoveredBitsCount),
		Hash:             p.Hash,
	}
}

func ConvPuzzleSolutionFromPb(s *contractpb.PuzzleSolution) *pow.Solution {
	return &pow.Solution{
		Timestamp: s.Timestamp.AsTime(),
		PreImage:  s.PreImage,
	}
}

func ConvPuzzleSolutionToPb(s *pow.Solution) *contractpb.PuzzleSolution {
	return &contractpb.PuzzleSolution{
		Timestamp: timestamppb.New(s.Timestamp),
		PreImage:  s.PreImage,
	}
}
