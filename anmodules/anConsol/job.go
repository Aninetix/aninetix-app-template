package anconsol

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Aninetix/core/anware"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func (m *AnModule) ProcessModule() {
	reader := bufio.NewReader(os.Stdin)
	headerConsol()
	fmt.Println(m.Config.Host)
	for {
		select {
		case <-m.ctx.Done():
			return
		default:
			fmt.Print("aninetix> ")
			cmd, _ := reader.ReadString('\n')
			cmd = strings.TrimSpace(cmd)
			m.handleCommand(cmd)
		}
	}
}

func headerConsol() {
	fmt.Println(ColorGreen, `
 ______    __   __    __    __   __   ______   ______   __    __  __
/\  __ \  /\ "-.\ \  /\ \  /\ "-.\ \ /\  ___\ /\__  _\ /\ \  /\_\_\_\
\ \  __ \ \ \ \-.  \ \ \ \ \ \ \-.  \\ \  __\ \/_/\ \/ \ \ \ \/_/\_\/_
 \ \_\ \_\ \ \_\\"\_\ \ \_\ \ \_\\"\_\\ \_____\  \ \_\  \ \_\  /\_\/\_\
  \/_/\/_/  \/_/ \/_/  \/_/  \/_/ \/_/ \/_____/   \/_/   \/_/  \/_/\/_/
  `, ColorReset)
}

func (m *AnModule) handleCommand(cmd string) {
	switch cmd {
	case "exit":
		m.mw.Send(anware.AnWareEvent{
			Source: m.name,
			Target: "anWare",
			Type:   "exit",
		})
	default:
		fmt.Println("Commande inconnue:", cmd)
		fmt.Print("\n")
	}
}
