package main

import (
	"context"
	"fmt"
	swagger "pinger/api"
	"sync"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	probing "github.com/prometheus-community/pro-bing"
)

func main() {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	sw := swagger.NewAPIClient(&swagger.Configuration{BasePath: "http://localhost:3001"})
	wg := &sync.WaitGroup{}

	for _, ctr := range containers {
		for _, nw := range ctr.NetworkSettings.Networks {
			if ctr.State == "running" {
				pinger, err := probing.NewPinger(nw.IPAddress)
				if err != nil {
					continue
				}
				pinger.Count = 1
				pinger.Timeout = 5 * time.Second
				wg.Add(1)
				go func() {
					err := pinger.Run()
					success := true
					if err != nil || pinger.PacketsRecv == 0 {
						success = false
					}
					_, err = sw.DefaultApi.PingsPut(context.Background(), []swagger.PingInfo{{ContainerIp: nw.IPAddress, Timestamp: time.Now(), Success: success}})
					if err != nil {
						fmt.Println(err)
					}
					wg.Done()
				}()
			}
		}
	}
	wg.Wait()
}
