package traefik

type Router struct {
	EntryPoints []string `json:"entryPoints"`
	Service     string   `json:"service"`
	Rule        string   `json:"rule"`
	Priority    int      `json:"priority,omitempty"`
	Status      string   `json:"status"`
	Using       []string `json:"using"`
	Name        string   `json:"name"`
	Provider    string   `json:"provider"`
	Middlewares []string `json:"middlewares,omitempty"`
}
