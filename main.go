package main

import (
	"context"
	"fmt"
	"os"
	openapi "pinger/api"
	"sync"
	"time"

	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	probing "github.com/prometheus-community/pro-bing"
)

var apiClient *client.Client
var sw *openapi.APIClient

func pingContainers(ctx context.Context) error {
	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return fmt.Errorf("could not retrieve containers list: %w", err)
	}

	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	pings := map[string]openapi.PingInfo{}

	successes := 0
	for _, ctr := range containers {
		for _, nw := range ctr.NetworkSettings.Networks {
			if ctr.State == "running" {
				pings[nw.IPAddress] = openapi.PingInfo{ContainerIp: nw.IPAddress, Timestamp: time.Now(), Success: false}

				pinger, err := probing.NewPinger(nw.IPAddress)
				if err != nil {
					log.Printf("could not create pinger for ip %s", nw.IPAddress)
					continue
				}
				pinger.Count = 1
				pinger.Timeout = 5 * time.Second
				wg.Add(1)
				go func() {
					defer wg.Done()

					err := pinger.RunWithContext(ctx)
					if err != nil {
						log.Printf("could not ping %s: %s", nw.IPAddress, err)
						return
					}
					if pinger.PacketsRecv == 0 {
						return
					}
					mx.Lock()
					cur := pings[nw.IPAddress]
					cur.Success = true
					pings[nw.IPAddress] = cur
					mx.Unlock()
					successes++
				}()
			}
		}
	}
	wg.Wait()

	log.Printf("pinged %d hosts, %d successes", len(pings), successes)

	pingslice := make([]openapi.PingInfo, 0, len(pings))
	for _, ping := range pings {
		pingslice = append(pingslice, ping)
	}

	_, err = sw.DefaultAPI.PingsPut(ctx).PingInfo(pingslice).Execute()
	if err != nil {
		return fmt.Errorf("could not post pings: %w", err)
	}

	return nil
}

func main() {
	var err error

	apiHost, ok := os.LookupEnv("API_HOST")
	if !ok {
		panic(fmt.Errorf("specify api host in environment variable API_HOST"))
	}

	apiClient, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	sw = openapi.NewAPIClient(&openapi.Configuration{Scheme: "http", Host: apiHost, Servers: []openapi.ServerConfiguration{{}}})

	for {
		err := pingContainers(context.Background())
		if err != nil {
			log.Printf("could not ping containers: %s", err)
		}

		time.Sleep(10 * time.Second)
	}
}
