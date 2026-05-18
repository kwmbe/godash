package main

import lg "charm.land/lipgloss/v2"

var (
	black =     lg.Color("#000")
	white =     lg.Color("#fff")

	darkgrey =  lg.Color("#333")
	grey =      lg.Color("#585858")
	lightgrey = lg.Color("#ccc")
)

type styles struct {
  text,
  subtleText,

  border,
  activeBorder,
  defaultBorder,

  header,
  footer,

  spinner lg.Style
}

func newStyles(isDarkBg bool) (s *styles) {
  s = new(styles)

  lightDark := lg.LightDark(isDarkBg)

  s.text = lg.NewStyle().Foreground(lightDark(black, white))
  s.subtleText = lg.NewStyle().Foreground(grey)


  s.border = lg.NewStyle().
    BorderForeground(lightDark(darkgrey, lightgrey)).
    Padding(0, 1)

  s.activeBorder = s.border.
    BorderStyle(lg.ThickBorder())

  s.defaultBorder = s.border.
    BorderStyle(lg.NormalBorder())

  s.header = s.text.Bold(true).Align(lg.Center)
  s.footer = s.subtleText.Align(lg.Center)

  s.spinner = s.text.Align(lg.Center).AlignVertical(lg.Center)

  return s
}

func panelBorder(p panel, active panel, s *styles) lg.Style {
  if p == active {
    return s.activeBorder
  }
  return s.defaultBorder
}
