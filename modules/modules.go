package modules

type Module struct {
	Name       string
	Path       string
	Categories []Tag
}

var Modules = []Module{
	{"pterm", "github.com/pterm/pterm", Tags{CLI}},
	{"testza", "github.com/MarvinJWendt/testza", Tags{TEST_FRAMEWORK}},
	{"fiber", "github.com/gofiber/fiber/v2", Tags{WEB_FRAMEWORK}},
	{"survey", "github.com/AlecAivazis/survey", Tags{CLI}},
	{"gin", "github.com/gin-gonic/gin", Tags{WEB_FRAMEWORK}},
	{"gorm", "github.com/go-gorm/gorm", Tags{ORM}},
	{"logrus", "github.com/sirupsen/logrus", Tags{LOGGING}},
	{"zap", "go.uber.org/zap", Tags{LOGGING}},
	{"zerolog", "github.com/rs/zerolog/log", Tags{LOGGING}},
	{"cobra", "github.com/spf13/cobra", Tags{CLI}},
	{"testify", "github.com/stretchr/testify", Tags{TEST_FRAMEWORK}},
	{"fyne", "github.com/fyne-io/fyne", Tags{GUI}},
	{"webview", "github.com/webview/webview", Tags{GUI}},
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
)
