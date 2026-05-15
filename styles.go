package main

import lg "charm.land/lipgloss/v2"

var (
  defaultText = lg.NewStyle().Foreground(lg.Color("#fff"))
  subtleText = lg.NewStyle().Foreground(lg.Color("#585858"))

  border = lg.NewStyle().
    BorderForeground(lg.Color("#ccc")).
    Padding(0, 1)

  selectedBorder = border.
    BorderStyle(lg.ThickBorder())

  defaultBorder = border.
    BorderStyle(lg.NormalBorder())

  headerStyle = defaultText.Bold(true).Align(lg.Center)
  footerStyle = subtleText.Align(lg.Center)

  spinnerStyle = defaultText.Align(lg.Center).AlignVertical(lg.Center)
)

func panelBorder(p panel, active panel) lg.Style {
  if p == active {
    return selectedBorder
  }
  return defaultBorder
}
