package atccmd

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/concourse/concourse/atc"
	"github.com/jessevdk/go-flags"
	"github.com/concourse/concourse/flag"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/acme/autocert"
)

type CommandSuite struct {
	suite.Suite
	*require.Assertions
}

func (s *CommandSuite) TestLetsEncryptDefaultIsUpToDate() {
	cmd := &ATCCommand{}

	parser := flags.NewParser(cmd, flags.Default)
	parser.NamespaceDelimiter = "-"

	opt := parser.Find("run").FindOptionByLongName("lets-encrypt-acme-url")
	s.NotNil(opt)

	s.Equal(opt.Default, []string{autocert.DefaultACMEDirectory})
}

func (s *CommandSuite) TestInvalidConcurrentRequestLimitAction() {
	cmd := &RunCommand{}
	parser := flags.NewParser(cmd, flags.None)
	_, err := parser.ParseArgs([]string{
		"--client-secret",
		"client-secret",
		"--concurrent-request-limit",
		fmt.Sprintf("%s:2", atc.GetInfo),
	})

	s.Contains(
		err.Error(),
		fmt.Sprintf("action '%s' is not supported", atc.GetInfo),
	)
}

func (s *CommandSuite) TestValidateCustomHTTPHeaders() {
	tmp := s.T().TempDir()

	validPath := filepath.Join(tmp, "valid.yml")
	err := os.WriteFile(validPath, []byte("X-Foo: bar\n"), 0644)
	s.NoError(err)

	invalidPath := filepath.Join(tmp, "invalid.yml")
	err = os.WriteFile(invalidPath, []byte("X-Foo; bar\n"), 0644)
	s.NoError(err)

	tests := []struct {
		name        string
		path        flag.File
		errContains string
	}{
		{
			name: "flag not set",
			path: flag.File(""),
		},
		{
			name:        "file does not exist",
			path:        flag.File("/nonexistent/headers.yml"),
			errContains: "failed to open",
		},
		{
			name:        "invalid yaml",
			path:        flag.File(invalidPath),
			errContains: "failed to parse",
		},
		{
			name: "valid yaml",
			path: flag.File(validPath),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			cmd := &RunCommand{}
			cmd.Server.CustomHTTPHeaders = tt.path

			err := cmd.validateCustomHTTPHeaders()

			if tt.errContains == "" {
				s.NoError(err)
			} else {
				s.ErrorContains(err, tt.errContains)
			}
		})
	}
}

func (s *CommandSuite) TestParseCustomHTTPHeaders() {
	tmp := s.T().TempDir()

	yamlPath := filepath.Join(tmp, "headers.yml")
	err := os.WriteFile(yamlPath, []byte("X-Foo: bar\nX-Baz: \"hello; world\"\n"), 0644)
	s.NoError(err)

	jsonPath := filepath.Join(tmp, "headers.json")
	err = os.WriteFile(jsonPath, []byte(`{"X-Foo": "bar", "X-Baz": "hello; world"}`), 0644)
	s.NoError(err)

	tests := []struct {
		name     string
		path     flag.File
		expected map[string]string
	}{
		{
			name:     "flag not set",
			path:     flag.File(""),
			expected: map[string]string{},
		},
		{
			name: "valid yaml",
			path: flag.File(yamlPath),
			expected: map[string]string{
				"X-Foo": "bar",
				"X-Baz": "hello; world",
			},
		},
		{
			name: "valid json",
			path: flag.File(jsonPath),
			expected: map[string]string{
				"X-Foo": "bar",
				"X-Baz": "hello; world",
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			cmd := &RunCommand{}
			cmd.Server.CustomHTTPHeaders = tt.path

			mapping, err := cmd.parseCustomHTTPHeaders()

			s.NoError(err)
			s.Equal(tt.expected, mapping)
		})
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, &CommandSuite{
		Assertions: require.New(t),
	})
}
