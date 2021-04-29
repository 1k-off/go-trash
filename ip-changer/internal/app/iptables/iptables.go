package iptables

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"ip_changer/internal/app/model"
	"ip_changer/internal/app/notifier"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var sh string
func init() {
	os := runtime.GOOS
	switch os {
	case "windows":
		sh = "C:\\Program Files\\Git\\bin\\bash.exe"
	case "darwin":
		sh = "/bin/bash"
	case "linux":
		sh = "/bin/bash"
	default:
		sh = "/bin/bash"
	}
}

func ChangeIp(networkName string, sc notifier.SlackClient) (*model.Ip, error) {
	var ip model.Ip
	netAddr := getDockerNetwork(networkName)
	res, err := exec.Command(sh, "bin/changer", netAddr).Output()
	if err != nil {
		return nil, err
	}
	ips := strings.Split(strings.ReplaceAll(string(res), "\n", ""), " ")
	ip.CurrentAddr = ips[0]
	ip.PreviousAddr = ips[1]
	changeTime := time.Now()
	ip.ChangeTime = changeTime.String()
	ip.NextChangeTime = (changeTime.Add(time.Minute * 5)).String()
	message := fmt.Sprintf("Current IP: %s. Previous IP: %s", ip.CurrentAddr, ip.PreviousAddr)
	log.Println(message)
	req := notifier.SlackJobNotification{
		Text: "IP changed",
		Details: message,
		Color: "success",
		IconEmoji: ":lobster:",
	}
	err = sc.SendJobNotification(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ips)
	return &ip, nil
}

func getDockerNetwork(networkName string) string{
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	net, err := cli.NetworkInspect(ctx, networkName, types.NetworkInspectOptions{})
	if err != nil {
		panic(err)
	}
	return net.IPAM.Config[0].Subnet
}