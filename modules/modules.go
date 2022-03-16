package modules

import (
	"sort"
	"strings"
)

func init() {
	// Sort the modules by the first tag.
	sort.Slice(Modules, func(i, j int) bool {
		return Modules[i].Tags[0] < Modules[j].Tags[0]
	})
}

type Module struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Tags Tags   `json:"tags"`
}

var Modules = []Module{
	{"cobra", "github.com/spf13/cobra", Tags{CLI}},
	{"pterm", "github.com/pterm/pterm", Tags{CLI}},
	{"survey", "github.com/AlecAivazis/survey/v2", Tags{CLI}},
	{"fyne", "github.com/fyne-io/fyne/v2", Tags{GUI}},
	{"webview", "github.com/webview/webview", Tags{GUI}},
	{"logrus", "github.com/sirupsen/logrus", Tags{LOGGING}},
	{"zap", "go.uber.org/zap", Tags{LOGGING}},
	{"zerolog", "github.com/rs/zerolog/log", Tags{LOGGING}},
	{"gorm", "github.com/go-gorm/gorm", Tags{ORM}},
	{"testify", "github.com/stretchr/testify", Tags{TEST_FRAMEWORK}},
	{"testza", "github.com/MarvinJWendt/testza", Tags{TEST_FRAMEWORK}},
	{"beego", "github.com/beego/beego/v2", Tags{WEB_FRAMEWORK}},
	{"fasthttp", "github.com/valyala/fasthttp", Tags{WEB_FRAMEWORK}},
	{"fiber", "github.com/gofiber/fiber/v2", Tags{WEB_FRAMEWORK}},
	{"gin", "github.com/gin-gonic/gin", Tags{WEB_FRAMEWORK}},
	{"lo", "github.com/samber/lo", Tags{UTILS}},
	{"colly", "github.com/gocolly/colly/v2", Tags{WEB_SCRAPER}},
	{"excelize", "github.com/qax-os/excelize/v2", Tags{DOCUMENT_PARSER}},
	{"gjson", "github.com/tidwall/gjson", Tags{PARSER}},
}

type Tag string
type Tags []Tag

const (
	WEB_FRAMEWORK   Tag = "Web Framework"
	WEB_SCRAPER     Tag = "Web Scraper"
	CLI             Tag = "CLI"
	TEST_FRAMEWORK  Tag = "Test Framework"
	ORM             Tag = "ORM"
	LOGGING         Tag = "Logging"
	GUI             Tag = "GUI"
	UTILS           Tag = "Utils"
	DOCUMENT_PARSER Tag = "Document Parser"
	PARSER          Tag = "Parser"
)

func (t Tags) String() string {
	var tags []string
	for _, tag := range t {
		tags = append(tags, string(tag))
	}

	return strings.Join(tags, ", ")
}
