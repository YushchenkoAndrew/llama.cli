package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"

	"github.com/YushchenkoAndrew/llama-cli/src/lib"
	"github.com/YushchenkoAndrew/llama-cli/src/models"
)

func list(sources []string) {
	fmt.Printf("%-46s %-9s %-9s %-9s %s\n", "NAME", "STATUS", "SIZE", "OUTPUT", "INPUT")
	for _, m := range models.Load().Data {
		if slices.Contains(sources, m.Source) {
			var status string
			if model := models.Cached(m.ID); model != nil {
				status = "ready"
			} else {
				status = "---"
			}

			fmt.Printf("%-46s %-9s %-9s %-9s %s\n", m.ID, status, lib.GetSize(m.Meta), strings.Join(m.Architecture.OutputModalities, "+"), strings.Join(m.Architecture.InputModalities, "+"))
		}
	}
}

func ps() {
	fmt.Printf("%-46s %-9s %s\n", "NAME", "STATUS", "SIZE")
	for _, m := range models.Active() {
		status := m.Status.Value
		props := models.Props(m.ID)

		if props.IsSleeping {
			status = "sleeping"
		}

		fmt.Printf("%-46s %-9s %s\n", m.ID, status, lib.GetSize(m.Meta))
	}
}

func stop(ids []string) {
	for _, m := range models.Active() {
		if m.Source == "preset" && (len(ids) == 0 || slices.Contains(ids, m.ID)) {
			models.Stop(m.ID)
		}
	}
}

func delete(model string) {
	for _, m := range models.Load().Data {
		if cached := models.Cached(m.ID); m.ID == model && cached != nil {
			models.Delete(cached.ID)
			break
		}
	}
}

func status() {
	cmd := exec.Command("systemctl", "status", "llama-cpp", "--no-pager")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

func logs() {
	cmd := exec.Command("journalctl", "-u", "llama-cpp", "-n", "30", "--no-pager")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

func help() {
	fmt.Println(`llc — llama.cli router helper

Usage:
  llc ls              List models (preset entries or cache dir)
  llc cache           Show cached model disk usage
  llc ps              Show loaded/loading models
  llc stop <id>       Unload a specific model
  llc rm <id>         Unload + delete a cached model
  llc status          Show systemd service status
  llc logs            Tail the llama-cpp systemd journal
  llc help            Show this help`)
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "ls":
		list([]string{"preset"})

	case "cache":
		list([]string{"cache"})

	case "ps":
		ps()

	case "stop":
		if len(args) < 1 {
			fmt.Println("Usage: lc stop <model-id>")
			os.Exit(1)
		}
		stop(args)

	case "rm", "delete":
		if len(args) < 1 {
			fmt.Println("Usage: lc rm <model-id>")
			os.Exit(1)
		}
		delete(args[0])

	case "status":
		status()

	case "logs":
		logs()

	case "help", "-h", "--help":
		help()

	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		help()
		os.Exit(1)
	}
}
