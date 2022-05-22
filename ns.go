package main

import (
	"fmt"
	"github.com/vishvananda/netns"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	curr, _ := netns.Get();
	fmt.Println(curr);

	data, err := ioutil.ReadFile("/proc/sys/net/netfilter/nf_conntrack_count")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	value, err := strconv.ParseUint(strings.TrimSpace(string(data)), 10, 64)
	if err != nil {
		fmt.Println("Error converting contents to int")
		os.Exit(1)
	}

	fmt.Println(value)


	dir, err := os.Open("/var/run/netns/")
	if err != nil {
		fmt.Println("Failed to open /var/run/netns directory")
	}
	files, err := dir.Readdir(0)

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())

		ns, _ := netns.GetFromName(v.Name());
		netns.Set(ns);

		curr, _ := netns.Get();
		fmt.Println(curr);

		data, err := ioutil.ReadFile("/proc/sys/net/netfilter/nf_conntrack_count")
		if err != nil {
			fmt.Println("Error reading file")
			os.Exit(1)
		}

		value, err := strconv.ParseUint(strings.TrimSpace(string(data)), 10, 64)
		if err != nil {
			fmt.Println("Error converting contents to int")
			os.Exit(1)
		}

		fmt.Println(value)
	}
}
