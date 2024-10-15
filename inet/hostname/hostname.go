package hostname

import (
	"maps"
	"net"
	"strings"

	"github.com/goexts/ggb/settings"
)

const (
	defaultMatchPrefix = "@"
	defaultMatchSuffix = ":" // Default suffix
)

// Replacer Interface definition
type Replacer interface {
	Match(name string) (string, bool)
	Replace(name string) string
}

// Hostname Default implementation
type Hostname struct {
	hosts           map[string]net.IP
	prefix          string
	suffix          string
	useCustomSuffix bool
}

type Option = settings.Setting[Hostname]

// UseCustomPrefix Customize the prefix
func UseCustomPrefix(prefix string) Option {
	return func(h *Hostname) {
		h.prefix = prefix
	}
}

// UseCustomSuffix Customize the suffix
func UseCustomSuffix(suffix string) Option {
	return func(h *Hostname) {
		h.useCustomSuffix = true
		h.suffix = suffix
	}
}

// WithHostMap Parse the name and return the corresponding IP address
func WithHostMap(hosts map[string]net.IP) Option {
	return func(h *Hostname) {
		if h.hosts == nil {
			h.hosts = make(map[string]net.IP)
		}
		maps.Copy(h.hosts, hosts)
	}
}

// WithHosts Parse the name and return the corresponding IP address
func WithHosts(list []string, sep string) Option {
	hosts := make(map[string]net.IP)
	for i := range list {
		nameip := strings.Split(list[i], sep)
		if len(nameip) != 2 {
			continue
		}
		if ip := net.ParseIP(nameip[1]); ip != nil {
			hosts[nameip[0]] = net.ParseIP(nameip[1])
		}
	}
	return func(h *Hostname) {
		if h.hosts == nil {
			h.hosts = hosts
			return
		}
		maps.Copy(h.hosts, hosts)
	}
}

// New Create a new instance of HostnameReplacer
func New(opts ...Option) *Hostname {
	return settings.Apply(&Hostname{
		prefix: defaultMatchPrefix,
		suffix: defaultMatchSuffix,
	}, opts)
}

func (r *Hostname) matchKey(key string) (string, bool) {
	if !strings.HasPrefix(key, r.prefix) {
		return "", false
	}
	key = key[len(r.prefix):]
	if i := strings.Index(key, r.suffix); i >= 0 {
		key = key[:i]
	}
	return key, true
}

func (r *Hostname) match(key string) (string, bool) {
	if ip, ok := r.hosts[key]; ok && ip != nil {
		return ip.String(), true
	}
	return "", false
}

// Match Return the corresponding IP address if the name ends with the specified suffix
func (r *Hostname) Match(name string) (string, bool) {
	if r.hosts == nil {
		return "", false
	}

	key, truth := r.matchKey(name)
	if !truth {
		return "", false
	}
	if ip, ok := r.match(key); ok {
		return ip, true
	}
	return "", false
}

// Replace is a hook that will Replace the name with the corresponding IP address
func (r *Hostname) Replace(name string) string {
	if r.hosts == nil {
		return name
	}
	key, truth := r.matchKey(name)
	if !truth {
		return name
	}
	if ip, matched := r.match(key); matched {
		origin := r.prefix + key
		if r.useCustomSuffix {
			origin += r.suffix
		}
		return strings.Replace(name, origin, ip, 1)
	}
	return name // If there is no match, return the original name
}
