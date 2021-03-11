package message

import (
	"context"
	"testing"

	"github.com/gotd/td/tg"
)

func TestMediaDice(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)

	expectDice := func(emoticon string) {
		expectSendMedia(&tg.InputMediaDice{Emoticon: emoticon}, mock)
	}

	expectDice(DiceEmoticon)
	expectDice(DartsEmoticon)
	expectDice(BasketballEmoticon)

	_, err := sender.Self().Dice(ctx)
	mock.NoError(err)
	_, err = sender.Self().Darts(ctx)
	mock.NoError(err)
	_, err = sender.Self().Basketball(ctx)
	mock.NoError(err)
}