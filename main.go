package ngrok

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/mtag-io/ngrok/config"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"syscall"
)

//go:embed config.yaml
var rawConfig []byte

//go:embed pkg.info
var rawPkgInfo []byte

//go:embed kill-script.template
var scriptContent string

var cfg = config.New(rawConfig, rawPkgInfo)

func CreateKillScript(pid int) {
	cwd, _ := os.Getwd()
	dest := path.Join(cwd, cfg.ScriptName)
	script := fmt.Sprintf(scriptContent, pid)
	err := os.WriteFile(dest, []byte(script), 0755)
	if err != nil {
		fmt.Println("FILE-ERROR: Couldn't create Ngrok kill script.")
		return
	}
	fmt.Printf("Created Ngrok kill script in: %s\n", dest)
}

func Status(port string, route string) *Tunnel {
	u := fmt.Sprintf("%s://localhost:%s/%s", cfg.Protocol, port, route)
	res, err := http.Get(u)
	if err != nil {
		return nil
	}
	body, err := io.ReadAll(res.Body)
	ngrok := Ngrok{}
	err = json.Unmarshal(body, &ngrok)
	if err != nil {
		fmt.Println("NGROK-ERROR: Unable to parse tunnel data")
		return nil
	}
	return &ngrok.Tunnels[0]
}

func Start(port string) int {
	cmd := exec.Command(cfg.Command, cfg.Protocol, port)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	err := cmd.Start()
	pid := cmd.Process.Pid
	if err != nil {
		log.Printf("NGROK-ERROR: Unable to start Ngrok:\n %s", err.Error())
		return 0
	}
	err = cmd.Process.Release()
	if err != nil {
		log.Printf("NGROK-ERROR: Unable to detach Ngrok:\n %s", err.Error())
		return 0
	}
	fmt.Printf("Ngrok launched in background with PID: %d\n", pid)
	return pid
}
