package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type (
	// Repo information.
	Repo struct {
		Owner string
		Name  string
	}
	// Build information.
	Build struct {
		Tag     string
		Event   string
		Number  int
		Commit  string
		Ref     string
		Branch  string
		Author  string
		Message string
		Pull    string
		Status  string
		Link    string
		Started int64
		Created int64
	}
	// Config for the plugin.
	Config struct {
		Method      string
		Username    string
		Password    string
		ContentType string
		Headers     []string
		URLs        []string
		SkipVerify  bool
		Debug       bool
		BearerToken string
	}
	Job struct {
		Started int64
	}
	// Plugin values.
	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {
	var (
		buf bytes.Buffer
		b   []byte
	)
	if len(p.Config.URLs) == 0 {
		return fmt.Errorf("You must provide at least one url")
	}
	data := struct {
		Repo  Repo  `json:"repo"`
		Build Build `json:"build"`
	}{p.Repo, p.Build}

	if err := json.NewEncoder(&buf).Encode(&data); err != nil {
		fmt.Printf("Error: Failed to encode JSON payload. %s\n", err)
		return err
	}

	b = buf.Bytes()
	for _, rawurl := range p.Config.URLs {
		uri, err := url.Parse(rawurl)

		if err != nil {
			fmt.Printf("Error: Failed to parse the hook URL. %s\n", err)
			os.Exit(1)
		}

		req, err := http.NewRequest(p.Config.Method, uri.String(), bytes.NewReader(b))

		if len(p.Config.Headers) > 0 {
			for _, value := range p.Config.Headers {
				header := strings.Split(value, "=")
				req.Header.Set(header[0], header[1])
			}
		}
		req.Header.Set("Content-Type", p.Config.ContentType)

		if p.Config.BearerToken != "" {
			req.Header.Set("Authorization", "Bearer "+p.Config.BearerToken)
		}
		if p.Config.Username != "" && p.Config.Password != "" {
			req.SetBasicAuth(p.Config.Username, p.Config.Password)
		}
		client := http.DefaultClient

		if p.Config.SkipVerify {
			client = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
		}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("Error: Failed to execute the HTTP request. %s\n", err)
			return err
		}

		defer resp.Body.Close()

		if p.Config.Debug || resp.StatusCode != http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				fmt.Printf("Error: Failed to read the HTTP response body. %s\n", err)
			}
			fmt.Printf(`
				URL: %s
				METHOD: %s
				HEADERS: %s
				REQUEST BODY: %s
				RESPONSE STATUS: %s
				RESPONSE BODY: %s`,
				req.URL,
				req.Method,
				req.Header,
				string(b),
				resp.Status,
				string(body),
			)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("invalid response codes %d", resp.StatusCode)
		}
	}

	return nil
}
