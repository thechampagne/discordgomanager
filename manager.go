/*
 * Copyright (c) 2022 XXIV
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package discordgomanager

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Command interface {
	Run([]string,*discordgo.Session,*discordgo.MessageCreate)
	GetCommand() string
}

type Manager struct {
	Prefix string
	Commands map[string]Command
}

func New(prefix string,) *Manager {
	return &Manager{Prefix: prefix, Commands: make(map[string]Command)}
}

func (m *Manager) AddCommand(command Command) {
	if _, ok := m.Commands[command.GetCommand()]; !ok {
		m.Commands[command.GetCommand()] = command
	}
}

func (m *Manager) AddCommands(command []Command) {
	for _, v := range command {
		if _, ok := m.Commands[v.GetCommand()]; !ok {
			m.Commands[v.GetCommand()] = v
		}
	}
}

func (m *Manager) Handler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	message := msg.Content
	if !strings.HasPrefix(message, m.Prefix) {
		return
	}
	args := strings.Split(strings.Replace(message, m.Prefix, "", 1), " ")
	command := strings.ToLower(args[0])

	if v, ok := m.Commands[command]; ok {
		v.Run(args[1:], s, msg)
	}
}