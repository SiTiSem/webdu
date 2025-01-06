package actions

import (
	"embed"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net"
)

func HttpRoot(frontendEmbed embed.FS) {
	servPort := flag.String("port", "7000", "Server port")
	UserPassword = flag.String("pass", PasswordRandom(), "Login password")
	flag.Parse()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	Route(e)
	fmt.Println(`
__          __  _     _____  _    _
\ \        / / | |   |  __ \| |  | |
 \ \  /\  / /__| |__ | |  | | |  | |
  \ \/  \/ / _ \ '_ \| |  | | |  | |
   \  /\  /  __/ |_) | |__| | |__| |
    \/  \/ \___|_.__/|_____/ \____/
____________________________________
`)
	fmt.Printf("Login:\t\tadmin\nPassword:\t%v\nAddress:\thttp://%v:%v\n____________________________________\n\n", *UserPassword, GetOutboundIP(), *servPort)
	e.Logger.Fatal(e.Start(":" + *servPort))
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr).IP.String()

	return localAddr
}
