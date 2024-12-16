package proxy

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
	"tokencarrier/internal/proxy/oidc"
)

func GetProxyHandler() (*http.ServeMux, error) {
	configuration, err := getConfiguration()
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", configuration.proxyHandler)
	mux.HandleFunc(configuration.LoginPath, configuration.login)
	mux.HandleFunc(configuration.LogoutPath, configuration.logout)
	mux.HandleFunc(configuration.ProfilePath, configuration.profile)
	mux.HandleFunc(configuration.CallbackPath, configuration.callback)
	mux.HandleFunc(configuration.BackChannelLogoutPath, configuration.backchannel)
	return mux, err
}

func (conf configuration) proxyHandler(w http.ResponseWriter, r *http.Request) {
	parsedURL, _ := url.Parse(r.URL.String())
	parsedURL.Host = conf.UpstreamServer
	parsedURL.Scheme = conf.UpstreamServerSchema
	proxyReq, err := http.NewRequest(r.Method, parsedURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	copyHeader(proxyReq.Header, r.Header)
	// add client ip to x-forwarded-for
	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		appendHostToXForwardHeader(proxyReq.Header, clientIP)
	}

	// add authorization header
	if proxyReq.Header.Get("Authorization") == "" {
		tokens, err := oidc.GetTokens(r)
		if err == nil {
			proxyReq.Header.Set("Authorization", "Bearer "+tokens.AccessToken)
		}
	}

	// do request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	// Copy the response body
	_, _ = io.Copy(w, resp.Body)
}

func (conf configuration) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (conf configuration) logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}

func (conf configuration) profile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (conf configuration) callback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (conf configuration) backchannel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
