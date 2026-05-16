package main

import lg "charm.land/lipgloss/v2"

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

  s.text = lg.NewStyle().Foreground(lightDark(lg.Color("#000"), lg.Color("#fff")))
  s.subtleText = lg.NewStyle().Foreground(lg.Color("#585858"))


  s.border = lg.NewStyle().
    BorderForeground(lightDark(lg.Color("#333"), lg.Color("#ccc"))).
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
