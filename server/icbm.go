package server

import (
	"context"
	"io"
	"log/slog"

	"github.com/mkaminski/goaim/oscar"
	"github.com/mkaminski/goaim/state"
)

type ICBMHandler interface {
	ChannelMsgToHostHandler(ctx context.Context, sess *state.Session, snacPayloadIn oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost) (*oscar.SNACMessage, error)
	ClientEventHandler(ctx context.Context, sess *state.Session, snacPayloadIn oscar.SNAC_0x04_0x14_ICBMClientEvent) error
	EvilRequestHandler(ctx context.Context, sess *state.Session, snacPayloadIn oscar.SNAC_0x04_0x08_ICBMEvilRequest) (oscar.SNACMessage, error)
	ParameterQueryHandler(context.Context) oscar.SNACMessage
}

func NewICBMRouter(logger *slog.Logger, handler ICBMHandler) ICBMRouter {
	return ICBMRouter{
		ICBMHandler: handler,
		RouteLogger: RouteLogger{
			Logger: logger,
		},
	}
}

type ICBMRouter struct {
	ICBMHandler
	RouteLogger
}

func (rt *ICBMRouter) RouteICBM(ctx context.Context, sess *state.Session, SNACFrame oscar.SNACFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	switch SNACFrame.SubGroup {
	case oscar.ICBMAddParameters:
		inSNAC := oscar.SNAC_0x04_0x02_ICBMAddParameters{}
		rt.logRequest(ctx, SNACFrame, inSNAC)
		return oscar.Unmarshal(&inSNAC, r)
	case oscar.ICBMParameterQuery:
		outSNAC := rt.ParameterQueryHandler(ctx)
		rt.logRequestAndResponse(ctx, SNACFrame, outSNAC, outSNAC.Frame, outSNAC.Body)
		return sendSNAC(SNACFrame, outSNAC.Frame, outSNAC.Body, sequence, w)
	case oscar.ICBMChannelMsgToHost:
		inSNAC := oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost{}
		if err := oscar.Unmarshal(&inSNAC, r); err != nil {
			return err
		}
		outSNAC, err := rt.ChannelMsgToHostHandler(ctx, sess, inSNAC)
		if err != nil {
			return err
		}
		rt.Logger.InfoContext(ctx, "user sent an IM", slog.String("recipient", inSNAC.ScreenName))
		if outSNAC == nil {
			return nil
		}
		rt.logRequestAndResponse(ctx, SNACFrame, inSNAC, outSNAC.Frame, outSNAC.Body)
		return sendSNAC(SNACFrame, outSNAC.Frame, outSNAC.Body, sequence, w)
	case oscar.ICBMEvilRequest:
		inSNAC := oscar.SNAC_0x04_0x08_ICBMEvilRequest{}
		if err := oscar.Unmarshal(&inSNAC, r); err != nil {
			return err
		}
		outSNAC, err := rt.EvilRequestHandler(ctx, sess, inSNAC)
		if err != nil {
			return err
		}
		rt.logRequestAndResponse(ctx, SNACFrame, inSNAC, outSNAC.Frame, outSNAC.Body)
		return sendSNAC(SNACFrame, outSNAC.Frame, outSNAC.Body, sequence, w)
	case oscar.ICBMClientErr:
		inSNAC := oscar.SNAC_0x04_0x0B_ICBMClientErr{}
		rt.logRequest(ctx, SNACFrame, inSNAC)
		return oscar.Unmarshal(&inSNAC, r)
	case oscar.ICBMClientEvent:
		inSNAC := oscar.SNAC_0x04_0x14_ICBMClientEvent{}
		if err := oscar.Unmarshal(&inSNAC, r); err != nil {
			return err
		}
		rt.logRequest(ctx, SNACFrame, inSNAC)
		return rt.ClientEventHandler(ctx, sess, inSNAC)
	default:
		return ErrUnsupportedSubGroup
	}
}
