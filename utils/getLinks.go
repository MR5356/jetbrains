package utils

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"sort"
	"strings"
)

const (
	JetbrainsApiBaseUrl = "https://data.services.jetbrains.com/products/releases"
)

var names = map[string]string{"AC": "AppCode", "CL": "CLion", "RSU": "ReSharper Ultimate", "DG": "DataGrip",
	"GO": "Goland", "IIU": "IntelliJ IDEA", "PS": "PhpStorm", "PCP": "PyCharm", "RD": "Rider",
	"RM": "RubyMine", "WS": "WebStorm", "DC": "DC", "DM": "DM", "DP": "DP"}

type apiDataItem struct {
	Date      string `json:"date"`
	Type      string `json:"type"`
	Downloads struct {
		Linux struct {
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"linux"`
		ThirdPartyLibrariesJson struct {
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"thirdPartyLibrariesJson"`
		Windows struct {
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"windows"`
		WindowsZip struct {
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"windowsZip"`
		Mac struct {
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"mac"`
		MacM1 struct {
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"macM1"`
	} `json:"downloads"`
	Patches struct {
		Win []struct {
			FromBuild    string  `json:"fromBuild"`
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"win"`
		Mac []struct {
			FromBuild    string  `json:"fromBuild"`
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"mac"`
		Unix []struct {
			FromBuild    string  `json:"fromBuild"`
			Link         string  `json:"link"`
			Size         float64 `json:"size"`
			ChecksumLink string  `json:"checksumLink"`
		} `json:"unix"`
	} `json:"patches"`
	NotesLink              string `json:"notesLink"`
	LicenseRequired        bool   `json:"licenseRequired"`
	Version                string `json:"version"`
	MajorVersion           string `json:"majorVersion"`
	Build                  string `json:"build"`
	Whatsnew               string `json:"whatsnew"`
	UninstallFeedbackLinks struct {
		WindowsJBR8             string `json:"windowsJBR8"`
		WindowsZipJBR8          string `json:"windowsZipJBR8"`
		Linux                   string `json:"linux"`
		ThirdPartyLibrariesJson string `json:"thirdPartyLibrariesJson"`
		Windows                 string `json:"windows"`
		WindowsZip              string `json:"windowsZip"`
		LinuxJBR8               string `json:"linuxJBR8"`
		Mac                     string `json:"mac"`
		MacJBR8                 string `json:"macJBR8"`
		MacM1                   string `json:"macM1"`
	} `json:"uninstallFeedbackLinks"`
	PrintableReleaseType interface{} `json:"printableReleaseType"`
}

type ApiDataSet map[string][]apiDataItem

type LinkItem struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Size     string `json:"size"`
	Build    string `json:"build"`
	Date     string `json:"date"`
	PlatFrom string `json:"platFrom"`
	Link     string `json:"link"`
}

type LinkItemSort []LinkItem

func (s LinkItemSort) Len() int {
	return len(s)
}

func (s LinkItemSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s LinkItemSort) Less(i, j int) bool {
	if s[i].PlatFrom == s[j].PlatFrom {
		return s[i].Name < s[j].Name
	} else {
		return s[i].PlatFrom < s[j].PlatFrom
	}
}

func getLinks() ApiDataSet {
	codes := make([]string, len(names))

	for k := range names {
		codes = append(codes, k)
	}

	client := resty.New()

	res, err := client.R().
		SetQueryParams(map[string]string{
			"code":   strings.Join(codes, ","),
			"latest": "true",
			"type":   "release",
		}).Get(JetbrainsApiBaseUrl)

	if err != nil {
		panic(err)
	}

	apiResult := ApiDataSet{}
	err = json.Unmarshal(res.Body(), &apiResult)
	if err != nil {
		return nil
	}
	return apiResult
}

func GetLinks(os string) LinkItemSort {
	apiResult := getLinks()
	windowsLinks := make(LinkItemSort, 0)
	linuxLinks := make(LinkItemSort, 0)
	macLinks := make(LinkItemSort, 0)
	macM1Links := make(LinkItemSort, 0)
	allLinks := make(LinkItemSort, 0)

	for ks, vs := range apiResult {
		for _, v := range vs {
			if len(v.Downloads.Windows.Link) > 0 {
				windowsLinks = append(windowsLinks, LinkItem{names[ks], v.Version, SizeConversion(v.Downloads.Windows.Size), v.Build, v.Date, "windows", v.Downloads.Windows.Link})
			}
			if len(v.Downloads.Linux.Link) > 0 {
				linuxLinks = append(linuxLinks, LinkItem{names[ks], v.Version, SizeConversion(v.Downloads.Linux.Size), v.Build, v.Date, "linux", v.Downloads.Linux.Link})
			}
			if len(v.Downloads.Mac.Link) > 0 {
				macLinks = append(macLinks, LinkItem{names[ks], v.Version, SizeConversion(v.Downloads.Mac.Size), v.Build, v.Date, "mac", v.Downloads.Mac.Link})
			}
			if len(v.Downloads.MacM1.Link) > 0 {
				macM1Links = append(macM1Links, LinkItem{names[ks], v.Version, SizeConversion(v.Downloads.MacM1.Size), v.Build, v.Date, "mac m1", v.Downloads.MacM1.Link})
			}
		}
	}
	allLinks = append(append(append(windowsLinks, linuxLinks...), macLinks...), macM1Links...)
	sort.Sort(windowsLinks)
	sort.Sort(linuxLinks)
	sort.Sort(macLinks)
	sort.Sort(macM1Links)
	sort.Sort(allLinks)
	if os == "windows" {
		return windowsLinks
	} else if os == "linux" {
		return linuxLinks
	} else if os == "mac" {
		return macLinks
	} else if os == "mac-m1" {
		return macM1Links
	} else if os == "all" {
		return allLinks
	} else {
		return []LinkItem{}
	}
}
