package golib

import (
    "net/url"
)

type ParsedURL struct {
    Scheme   string
    Username string
    Password string
    Hostname string
    Port     string
    Path     string
    Fragment string
}

func URLParse(rawURL string) (ParsedURL, error) {
    u, err := url.Parse(rawURL)
    if err != nil {
        return ParsedURL{}, err
    }
    username := ""
    password := ""
    if u.User != nil {
        username = u.User.Username()
        password, _ = u.User.Password()
    }
    port := ""
    if u.Port() != "" {
        port = u.Port()
    }
    return ParsedURL{
        Scheme:   u.Scheme,
        Username: username,
        Password: password,
        Hostname: u.Hostname(),
        Port:     port,
        Path:     u.Path,
        Fragment: u.Fragment,
    }, nil
}
