package main

import (
  "strings"

  tea "charm.land/bubbletea/v2"
  lg  "charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
  // test loading indicator
  if false {
    v := tea.NewView(spinnerStyle.
      Width(m.width).
      Height(m.height).
      Render(m.spinner.View()))
    v.AltScreen = true
    return v
  }

  // left panel text content
  var lcontent strings.Builder
  lcontent.WriteString(defaultText.Render("Test content ") + subtleText.Render(":)"))

  // right panel
  var rcontent strings.Builder
  rcontent.WriteString(spinnerStyle.Render(m.spinner.View()))

  // set panel sizes
  leftW  := m.width / 2
  rightW := m.width / 2
  height := m.height - 2

  leftPanel := panelBorder(logPanel, m.activePanel).
    Width(leftW).
    Height(height).
    Render(lcontent.String())

  rightPanel := panelBorder(inputPanel, m.activePanel).
    Width(rightW).
    Height(height).
    Render(rcontent.String())

  // --- assemble ---
  header := headerStyle.
    Width(m.width).
    Render("godash")

  body := lg.JoinHorizontal(lg.Top, leftPanel, rightPanel)

  footer := footerStyle.
    Width(m.width).
    Render("ctrl +   navigate  •  q quit")

  v := tea.NewView(lg.JoinVertical(lg.Left, header, body, footer))
  v.AltScreen = true
  return v
}
