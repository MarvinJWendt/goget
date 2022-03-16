package modules

type Module struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Tags []Tag  `json:"tags"`
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
}

type Tag string
type Tags []Tag

const (
	WEB_FRAMEWORK  Tag = "Web Framework"
	CLI            Tag = "CLI"
	TEST_FRAMEWORK Tag = "Test Framework"
	ORM            Tag = "ORM"
	LOGGING        Tag = "Logging"
	GUI            Tag = "GUI"
	UTILS          Tag = "Utils"
)
