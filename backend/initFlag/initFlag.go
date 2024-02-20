package initFlag

import "flag"

var (
	RestAPIPort        string
	RestAPICorsMode    string
	JsonServerHost     string
	JsonServerPort     string
	FrontendPort       string
	Standalone         int
	JsonServerFullPath string
)

func InitFlag() {
	flag.StringVar(&RestAPIPort, "restapiport", "3001", "REST API port")
	flag.StringVar(&RestAPICorsMode, "hardcodecors", "*", "Hardcode frontend url for better protection")
	flag.StringVar(&JsonServerHost, "jsonserverhost", "127.0.0.1", "JSON server host address")
	flag.StringVar(&JsonServerPort, "jsonserverport", "3000", "JSON server port")
	flag.StringVar(&FrontendPort, "frontendport", "80", "Frontend port")
	flag.IntVar(&Standalone, "standalone", 0, "standalone mode 0 for off. 1 for only frontend host. 2 for only rest api host.")

	flag.Parse()
}
