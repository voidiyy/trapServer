package utils

import (
	"context"
	"google.golang.org/grpc/peer"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
)

const (
	XClientIp        = "X-Client-Ip"         // Standard headers used by Amazon EC2, Heroku, and others.
	XForwardedFor    = "X-Forwarded-For"     // Load-balancers (AWS ELB) or proxies.
	CFConnectingIp   = "CF-Connecting-Ip"    // @see https://support.cloudflare.com/hc/en-us/articles/200170986-How-does-Cloudflare-handle-HTTP-Request-headers-
	FastlyClientIp   = "Fastly-Client-Ip"    // Fastly and Firebase hosting header (When forwarded to cloud function)
	TrueClientIp     = "True-Client-Ip"      // Akamai and Cloudflare: True-Client-IP.
	XRealIp          = "X-Real-Ip"           // Default nginx proxy/fcgi; alternative to x-forwarded-for, used by some proxies.
	XClusterClientIp = "X-Cluster-Client-Ip" // (Rackspace LB and Riverbed's Stingray) http://www.rackspace.com/knowledge_center/article/controlling-access-to-linux-cloud-sites-based-on-the-client-ip-address
	XForwarded       = "X-Forwarded"
	ForwardedFor     = "Forwarded-For"
	Forwarded        = "Forwarded"
)

var requestHeaders = []string{
	XClientIp,
	XForwardedFor,
	CFConnectingIp,
	FastlyClientIp,
	TrueClientIp,
	XRealIp,
	XClusterClientIp,
	XForwarded,
	ForwardedFor,
	Forwarded,
}

func CheckHeaders(r *http.Request) map[string]string {

	matched := make(map[string]string)
	for _, header := range requestHeaders {
		headerValue := r.Header.Get(header)
		if headerValue != "" {
			matched[header] = headerValue
		}
	}
	if len(matched) > 0 {
		return matched
	}

	return nil
}

type IpAddress struct {
	Value string `json:"value"`
}

func ParseIP(r *http.Request) *IpAddress {

	ip := getIP(r)
	if ip == "" {
		return &IpAddress{Value: r.RemoteAddr}
	}
	return &IpAddress{Value: ip}
}

func getIP(r *http.Request) string {
	if len(r.Header) > 0 {
		for _, header := range requestHeaders {
			switch header {
			case XForwarded:
				host, isIP := getForwarded(r.Header.Get(header))
				if !isIP {
					return host
				}
			default:
				if host := r.Header.Get(header); isCorrectIP(host) {
					return host
				}
			}
		}
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && isCorrectIP(host) {
		return host
	}

	return r.RemoteAddr
}

func ParseIPContext(ctx context.Context) *IpAddress {
	ip := getIPContext(ctx)

	if ip == "" {
		return nil
	}

	return &IpAddress{Value: parseLocalHost(ip)}
}

func getIPContext(ctx context.Context) string {
	md := metadata.ExtractIncoming(ctx)

	if len(md) > 0 {
		for _, header := range requestHeaders {
			switch header {
			case XForwardedFor:
				host, isIP := getForwarded(md.Get(header))
				if isIP {
					return host
				}
			default:
				if host := md.Get(header); isCorrectIP(host) {
					return host
				}
			}
		}
	}

	peerr, ok := peer.FromContext(ctx)
	if ok && peerr.Addr != nil {
		host, _, err := net.SplitHostPort(peerr.Addr.String())
		if err == nil && isCorrectIP(host) {
			return host
		}
	}

	return ""
}

func getForwarded(headers string) (string, bool) {
	if headers == "" {
		return "", false
	}

	forwardedIPs := strings.Split(headers, ",")

	for _, ip := range forwardedIPs {
		ip = strings.TrimSpace(ip)

		if split := strings.Split(ip, ":"); len(split) == 2 {
			ip = split[0]
		}
		if isCorrectIP(ip) {
			return ip, true
		}
	}

	return "", false
}

func isCorrectIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func parseLocalHost(ip string) string {
	if ip == "localhost" || ip == "127.0.0.1" {
		return "::1"
	}

	return ip
}
