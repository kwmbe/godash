package main

import (
  tea "charm.land/bubbletea/v2"
  "charm.land/bubbles/v2/spinner"
)

func clamp(val, min, max int) int {
  if val < min { return min }
  if val > max { return max }
  return val
}

func (m model) cyclePanel(dir int) (tea.Model, tea.Cmd) {
  m.activePanel = panel(clamp(int(m.activePanel) + dir, 0, int(lastPanel) - 1))
  return m, nil
}

func (m model) handleStyle(msg tea.BackgroundColorMsg) (tea.Model, tea.Cmd) {
  m.styles = newStyles(msg.IsDark())
  return m, nil
}

func (m model) handleWindowSize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
  m.width = msg.Width
  m.height = msg.Height
  return m, nil
}

func (m model) handleSpinnerTick(msg spinner.TickMsg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  m.spinner, cmd = m.spinner.Update(msg)
  return m, cmd
}

func (m model) handleKeyPress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
  switch msg.String() {
  case "ctrl+c", "q":
    return m, tea.Quit
  case "ctrl+left", "ctrl+h":
    return m.cyclePanel(-1)
  case "ctrl+right", "ctrl+l":
    return m.cyclePanel(1)
  }
  return m, nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.BackgroundColorMsg:
    return m.handleStyle(msg)
  case tea.WindowSizeMsg:
    return m.handleWindowSize(msg)
  case tea.KeyPressMsg:
    return m.handleKeyPress(msg)
  case spinner.TickMsg:
    return m.handleSpinnerTick(msg)
  }
  return m, nil
}

