package phone

import "github.com/twilio/twilio-go/twiml"

type Message interface {
	Element() twiml.Element
	JSON() ([]byte, error)
}

type TTS struct {
	Text string `json:"text"`
}

func (t *TTS) Element() twiml.Element {
	return &twiml.VoiceSay{
		Message: t.Text,
	}
}

type Audio struct {
	URL string `json:"url"`
}

func (a *Audio) Element() twiml.Element {
	return twiml.VoicePlay{
		Url: a.URL,
	}
}
