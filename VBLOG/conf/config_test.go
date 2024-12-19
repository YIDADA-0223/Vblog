package conf_test

import (
	"os"
	"testing"

	"gitee.com/VBLOG/conf"
	"github.com/go-playground/assert/v2"
)

func TestToYaml(t *testing.T) {
	t.Log(conf.Defalut().ToYaml())
}
func TestToLoadFromEnv(t *testing.T) {
	os.Setenv("DATASOURCE_USERNAME", "env test")
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().ToYaml())
	assert.Equal(t, conf.C().MySQL.Username, "env test")
}
func TestToLoadFromYAML(t *testing.T) {
	err := conf.LoadConfigFromYaml("./application.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().ToYaml())
}
