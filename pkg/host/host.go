package host

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

type Host struct {
}

func makeTempFile(body []byte) (string, error) {
	hash := md5.New()
	hash.Write(body)
	hashed := hash.Sum(nil)
	name := fmt.Sprintf("%x", hashed)
	logrus.Info(name)
	codeDir := strings.TrimRight(viper.GetString("codedir"), "/")
	f, e := os.Create(fmt.Sprintf("%s/%s.go", codeDir, name))

	if e != nil {
		logrus.Error(e)
		return "", e
	}

	_, e = f.Write(body)
	if e != nil {
		return "", e
	}

	return f.Name(), nil

}

func Format(body []byte) ([]byte, error) {
	f, e := makeTempFile(body)
	if e != nil {
		return nil, e
	}
	logrus.Info(f)

	out, e := exec.Command("sudo", "-u", "eulerexec", "/usr/local/go/bin/go", "fmt", f).CombinedOutput()

	if e != nil {
		logrus.Info(e.Error)
		logrus.Info(string(out))
		return nil, fmt.Errorf("%s - %s", e.Error(), string(out))
	}
	logrus.Info(string(out))

	formatted, e := ioutil.ReadFile(f)
	if e != nil {
		return nil, e
	}

	return formatted, nil
}

func Run(body []byte) ([]byte, error) {
	f, e := makeTempFile(body)
	if e != nil {
		return nil, e
	}
	logrus.Info(f)

	out, e := exec.Command("sudo", "-u", "eulerexec", "/usr/local/go/bin/go", "run", f).CombinedOutput()

	if e != nil {
		logrus.Info(e.Error)
		logrus.Info(string(out))
		return nil, fmt.Errorf("%s - %s", e.Error(), string(out))
	}
	logrus.Info(string(out))

	return out, nil
}
