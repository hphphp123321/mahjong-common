package player

import pb "github.com/hphphp123321/mahjong-common/services/mahjong/v1"

type GameAgent interface {
	ChooseAction() (*pb.Action, error)
}
