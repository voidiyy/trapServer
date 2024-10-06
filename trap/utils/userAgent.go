package utils

import (
	"fmt"
	"github.com/mileusna/useragent"
)

type UserAgent struct {
	Raw     string
	OS      *OS
	Browser *Browser
	Device  *Device
}

type Browser struct {
	Name    string
	Version string
}

type OS struct {
	Name    string
	Version string
}

type Device struct {
	Name string
	Type string
}

func ParseUserAgent(agent string) *UserAgent {
	if agent == "" {
		return nil
	}

	parse := useragent.Parse(agent)

	ua := &UserAgent{
		Raw: agent,
		Device: &Device{
			Name: parse.Device,
			Type: deviceType(parse),
		},
	}

	if parse.Name != "" {
		ua.Browser = &Browser{
			Name:    parse.Name,
			Version: parse.Version,
		}
	}

	if parse.OS != "" {
		ua.OS = &OS{
			Name:    parse.OS,
			Version: parse.Version,
		}
	}

	return ua
}

func deviceType(agent useragent.UserAgent) string {
	if agent.Mobile {
		if agent.IsAndroid() {
			return "Android"
		}
		if agent.IsIOS() {
			return "iOS"
		}

		return "Mobile"
	}

	if agent.Desktop {
		return "Desktop"
	}

	if agent.Tablet {
		return "Tablet"
	}

	if agent.Bot {
		if agent.IsFacebookbot() {
			return "FacebookBot"
		}
		if agent.IsGooglebot() {
			return "GoogleBot"
		}
		if agent.IsTwitterbot() {
			return "TwitterBot"
		}

		if agent.IsYandexbot() {
			return "YandexBot"
		}

		return "Bot"
	}

	return "Unknown"
}

func (b *Browser) string() string {
	if b == nil {
		return "Browser: N/A"
	}
	return fmt.Sprintf("Browser: %s (Version: %s)", b.Name, b.Version)
}

func (os *OS) string() string {
	if os == nil {
		return "OS: N/A"
	}
	return fmt.Sprintf("OS: %s (Version: %s)", os.Name, os.Version)
}

func (d *Device) string() string {
	if d == nil {
		return "Device: N/A"
	}
	return fmt.Sprintf("Device: %s (Type: %s)", d.Name, d.Type)
}

func (ua *UserAgent) String() string {
	return fmt.Sprintf(
		"User Agent Information:\nRaw: %s\n%s\n%s\n%s\n",
		ua.Raw, ua.OS.string(), ua.Browser.string(), ua.Device.string(),
	)
}
