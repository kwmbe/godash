package main

import (
  "time"

  tea "charm.land/bubbletea/v2"
  "charm.land/bubbles/v2/spinner"
)

type panel int

const (
  logPanel panel = iota
  inputPanel
  lastPanel // don't change this
)

type model struct {
  width      int
  height     int

  activePanel panel

  spinner    spinner.Model
  loading    bool
}

func initialModel() model {
  s := spinner.New()
  s.Spinner = spinner.Spinner{
    Frames: []string{"", "", "", "", "", ""},
    FPS:    time.Second / 12,
  }

  return model{
    activePanel: logPanel,
    spinner:     s,
    loading:     true,
  }
}

func (m model) Init() tea.Cmd {
  return m.spinner.Tick
}

