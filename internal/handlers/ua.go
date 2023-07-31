package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mssola/user_agent"
)

type ClientInfo struct {
	UserAgent      string `json:"user-agent"`
	IP             string `json:"ip"`
	Method         string `json:"method"`
	URI            string `json:"uri"`
	HostName       string `json:"host"`
	Protocol       string `json:"protocol"`
	Referer        string `json:"referer"`
	Language       string `json:"language"`
	ContentType    string `json:"content-type"`
	OS             string `json:"os"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser-version"`
	Engine         string `json:"engine"`
	EngineVersion  string `json:"engine-version"`
}

func UAHandler(w http.ResponseWriter, r *http.Request) {
	userAgentStr := r.Header.Get("User-Agent")
	ip := r.RemoteAddr
	method := r.Method
	uri := r.RequestURI
	hostName := r.Host
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}
	referer := r.Header.Get("Referer")
	language := r.Header.Get("Accept-Language")
	contentType := r.Header.Get("Content-Type")

	ua := user_agent.New(userAgentStr)
	browser, browserVersion := ua.Browser()
	engine, engineVersion := ua.Engine()
	osInfo := ua.OS()

	info := ClientInfo{
		UserAgent:      userAgentStr,
		IP:             ip,
		Method:         method,
		URI:            uri,
		HostName:       hostName,
		Protocol:       protocol,
		Referer:        referer,
		Language:       language,
		ContentType:    contentType,
		OS:             osInfo,
		Browser:        browser,
		BrowserVersion: browserVersion,
		Engine:         engine,
		EngineVersion:  engineVersion,
	}

	jsonData, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
