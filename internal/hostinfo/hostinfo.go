package hostinfo

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"
)

type Info struct {
	Hostname    string   `json:"hostname"`
	OS          string   `json:"os"`
	Arch        string   `json:"arch"`
	CPUs        int      `json:"cpus"`
	GoVersion   string   `json:"goVersion"`
	LocalIPs    []string `json:"localIps"`
	CollectedAt string   `json:"collectedAt"`
}

func Collect() (Info, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return Info{}, fmt.Errorf("get hostname: %w", err)
	}

	ips, err := localIPs()
	if err != nil {
		return Info{}, fmt.Errorf("get local IPs: %w", err)
	}

	return Info{
		Hostname:    hostname,
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		CPUs:        runtime.NumCPU(),
		GoVersion:   runtime.Version(),
		LocalIPs:    ips,
		CollectedAt: time.Now().UTC().Format(time.RFC3339),
	}, nil
}

func WriteHTTP(w http.ResponseWriter, r *http.Request, info Info) error {
	if wantsText(r) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		_, err := w.Write([]byte(FormatText(info)))
		if err != nil {
			return fmt.Errorf("write text response: %w", err)
		}

		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(info); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}

	return nil
}

func FormatText(info Info) string {
	return fmt.Sprintf(
		"hostname: %s\nos: %s\narch: %s\ncpus: %d\ngo_version: %s\nlocal_ips: %s\ncollected_at: %s\n",
		info.Hostname,
		info.OS,
		info.Arch,
		info.CPUs,
		info.GoVersion,
		strings.Join(info.LocalIPs, ", "),
		info.CollectedAt,
	)
}

func wantsText(r *http.Request) bool {
	if strings.EqualFold(r.URL.Query().Get("format"), "text") {
		return true
	}

	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "text/plain")
}

func localIPs() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("list interfaces: %w", err)
	}

	ips := make([]string, 0)

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addresses, err := iface.Addrs()
		if err != nil {
			return nil, fmt.Errorf("list addresses for %s: %w", iface.Name, err)
		}

		for _, address := range addresses {
			ipNet, ok := address.(*net.IPNet)
			if !ok || ipNet.IP == nil || ipNet.IP.IsLoopback() {
				continue
			}

			ips = append(ips, ipNet.IP.String())
		}
	}

	sort.Strings(ips)
	return ips, nil
}
