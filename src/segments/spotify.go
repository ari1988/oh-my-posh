package segments

import (
	"github.com/jandedobbeleer/oh-my-posh/src/properties"
)

type Spotify struct {
	base

	MusicPlayer
}

type MusicPlayer struct {
	Status string
	Artist string
	Track  string
	Icon   string
}

const (
	// PlayingIcon indicates a song is playing
	PlayingIcon properties.Property = "playing_icon"
	// PausedIcon indicates a song is paused
	PausedIcon properties.Property = "paused_icon"
	// StoppedIcon indicates a song is stopped
	StoppedIcon properties.Property = "stopped_icon"
	// AdIcon indicates an advertisement is playing
	AdIcon properties.Property = "ad_icon"

	playing     = "playing"
	stopped     = "stopped"
	paused      = "paused"
	advertising = "advertising"
)

func (s *Spotify) Template() string {
	return " {{ .Icon }}{{ if ne .Status \"stopped\" }}{{ .Artist }} - {{ .Track }}{{ end }} "
}

func (s *Spotify) resolveIcon() {
	switch s.Status {
	case stopped:
		// in this case, no artist or track info
		s.Icon = s.props.GetString(StoppedIcon, "\uF04D ")
	case paused:
		s.Icon = s.props.GetString(PausedIcon, "\uF8E3 ")
	case playing:
		s.Icon = s.props.GetString(PlayingIcon, "\uE602 ")
	}
}
