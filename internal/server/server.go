package server

import (
	"fmt"
	"log"
	"net/http"
	"system_monitor/internal/monitor"
	"text/template"
)

func HandleServer() {
	todoHandler := func(res http.ResponseWriter, req *http.Request) {
		templ := template.Must(template.ParseFiles("htmx/index.html"))
		templ.Execute(res, nil)
	}

	http.HandleFunc("/", todoHandler)
	http.HandleFunc("/ram", func(w http.ResponseWriter, r *http.Request) {
		ramInfo, err := monitor.GetRamInfo()
		if err != nil {
			http.Error(w, "Failed to get RAM info", http.StatusInternalServerError)
			return
		}

		ramTmpl := template.Must(template.New("ram").Parse(`
			<div hx-get="/ram" hx-trigger="every 2s" hx-swap="outerHTML">
			<ul>
				<li><strong>RAM Free:</strong> <span>{{.Free}} MB</span></li>
				<li><strong>RAM Used:</strong> <span>{{.Used}} MB</span></li>
				<li><strong>RAM Total:</strong> <span>{{.Total}} MB</span></li>
				<li><strong>RAM Usage:</strong> <span>{{.Percentage}}%</span></li>
			</ul>
			</div>
		`))

		ramTmpl.Execute(w, ramInfo)
	})
	http.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
		cpuInfo, err := monitor.GetCpuInfo()
		if err != nil {
			http.Error(w, "Failed to get RAM info", http.StatusInternalServerError)
			return
		}
		cpuTmpl := template.Must(template.New("cpu").Parse(`
		<div hx-get="/cpu" hx-trigger="every 2s" hx-swap="outerHTML">
          {{range .}}
        <div class="cpu-core">
          <strong>Core:</strong> <span>{{.Core}}</span> — <strong></strong>
          <span>{{.Model}}</span> — <strong>Usage:</strong>
          <span style="margin-left: 50px">{{.Usage}}%</span>
        </div>
        {{end}}
        </div>
		`))

		cpuTmpl.Execute(w, cpuInfo)
	})
	http.HandleFunc("/host", func(w http.ResponseWriter, r *http.Request) {
		hostInfo, err := monitor.GetHostInfo()
		if err != nil {
			http.Error(w, "Failed to get RAM info", http.StatusInternalServerError)
			return
		}
		hostTmpl := template.Must(template.New("host").Parse(`
          	<div hx-get="/host" hx-trigger="every 2s" hx-swap="outerHTML">
            	<ul>
					<li><strong>Hostname:</strong> <span>{{.Host}}</span></li>
					<li><strong>OS:</strong> <span>{{.Os}}</span></li>
					<li>
					<strong>Uptime:</strong>
					<span data-uptime>{{.Uptime}} min</span>
              		</li>
            	</ul>
          	</div>
		`))
		hostTmpl.Execute(w, hostInfo)
	})

	fmt.Println("Server Started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
