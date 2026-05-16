package main

import (
  "strings"

  tea "charm.land/bubbletea/v2"
  lg  "charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
  // style hasn't loaded yet
  if m.styles == nil {
    v := tea.NewView(m.spinner.View())
    v.AltScreen = true
    return v
  }

  var (
    s = m.styles
  )

  // left panel text content
  var lcontent strings.Builder
  lcontent.WriteString(s.text.Render("Test content ") + s.subtleText.Render(":)"))

  // right panel
  var rcontent strings.Builder
  rcontent.WriteString(s.spinner.Render(m.spinner.View()))

  // set panel sizes
  leftW  := m.width / 2
  rightW := m.width / 2
  height := m.height - 2

  leftPanel := panelBorder(logPanel, m.activePanel, s).
    Width(leftW).
    Height(height).
    Render(lcontent.String())

  rightPanel := panelBorder(inputPanel, m.activePanel, s).
    Width(rightW).
    Height(height).
    Render(rcontent.String())

  // --- assemble ---
  header := s.header.
    Width(m.width).
    Render("godash")

  body := lg.JoinHorizontal(lg.Top, leftPanel, rightPanel)

  footer := s.footer.
    Width(m.width).
    Render("ctrl +   navigate  •  q quit")

  v := tea.NewView(lg.JoinVertical(lg.Left, header, body, footer))
  v.AltScreen = true
  return v
}
