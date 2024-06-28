// A generated module for Kubeconform functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"strconv"
	"strings"
)

type Kubeconform struct {
}

type Lint struct {
	schemas []struct {
		pattern string
		dir     *Directory
	}
	exitOnError           bool
	ignoreFilenamePattern []string
	ignoreMissingSchemas  bool
	kubernetesVersion     string
	parallelism           *int
	outputFormat          string
	reject                []string
	skipGVK               []string
	strict                bool
	summary               bool
	verbose               bool
}

const (
	OutputFormatJson   = "json"
	OutputFormatJunit  = "junit"
	OutputFormatPretty = "pretty"
	OutputFormatTap    = "tap"
	OutputFormatText   = "text"
)

const CRDSchemaPattern = ""

//goland:noinspection ALL
func (m *Kubeconform) CRD_To_Schema(crdsDir *Directory) (*Directory, error) {
	panic("implement me")
}

func (lint Lint) WithSchemas(pattern string, schemaDir *Directory) Lint {
	lint.schemas = append(lint.schemas, struct {
		pattern string
		dir     *Directory
	}{pattern: pattern, dir: schemaDir})
	return lint
}

func (m *Kubeconform) WithSchemas(pattern string, schemaDir *Directory) Lint {
	return Lint{}.WithSchemas(pattern, schemaDir)
}

func (lint Lint) ExitOnError() Lint {
	lint.exitOnError = true
	return lint
}

func (m *Kubeconform) ExitOnError() Lint {
	return Lint{}.ExitOnError()
}

func (lint Lint) IgnoreFilenamePattern(pattern string) Lint {
	lint.ignoreFilenamePattern = append(lint.ignoreFilenamePattern, pattern)
	return lint
}

func (m *Kubeconform) IgnoreFilenamePattern(pattern string) Lint {
	return Lint{}.IgnoreFilenamePattern(pattern)
}

func (lint Lint) IgnoreMissingSchemas() Lint {
	lint.ignoreMissingSchemas = true
	return lint
}

func (m *Kubeconform) IgnoreMissingSchemas() Lint {
	return Lint{}.IgnoreMissingSchemas()
}

func (lint Lint) WithKubernetesVersion(version string) Lint {
	lint.kubernetesVersion = version
	return lint
}

func (m *Kubeconform) WithKubernetesVersion(version string) Lint {
	return Lint{}.WithKubernetesVersion(version)
}

func (lint Lint) WithParallelism(n int) Lint {
	lint.parallelism = &n
	return lint
}

func (m *Kubeconform) WithParallelism(n int) Lint {
	return Lint{}.WithParallelism(n)
}

func (lint Lint) WithOutputFormat(format string) Lint {
	lint.outputFormat = format
	return lint
}

func (m *Kubeconform) WithOutputFormat(format string) Lint {
	return Lint{}.WithOutputFormat(format)
}

func (lint Lint) RejectGVKs(gvks []string) Lint {
	lint.reject = append(lint.reject, gvks...)
	return lint
}

func (m *Kubeconform) RejectGVKs(gvks []string) Lint {
	return Lint{}.RejectGVKs(gvks)
}

func (lint Lint) Strict() Lint {
	lint.strict = true
	return lint
}

func (m *Kubeconform) Strict() Lint {
	return Lint{}.Strict()
}

func (lint Lint) WithSummary() Lint {
	lint.summary = true
	return lint
}

func (m *Kubeconform) WithSummary() Lint {
	return Lint{}.WithSummary()
}

func (lint Lint) Verbose() Lint {
	lint.verbose = true
	return lint
}

func (m *Kubeconform) Verbose() Lint {
	return Lint{}.Verbose()
}

func (lint Lint) Lint(ctx context.Context, manifests *Directory) (string, error) {
	var args []string

	if lint.exitOnError {
		args = append(args, "--exit-on-error")
	}

	if len(lint.ignoreFilenamePattern) > 0 {
		for _, pattern := range lint.ignoreFilenamePattern {
			args = append(args, "--ignore-filename-pattern", pattern)
		}
	}

	if lint.ignoreMissingSchemas {
		args = append(args, "--ignore-missing-schemas")
	}

	if lint.kubernetesVersion != "" {
		args = append(args, "--kubernetes-version", lint.kubernetesVersion)
	}

	if lint.parallelism != nil {
		args = append(args, "-n", strconv.Itoa(*lint.parallelism))
	}

	if lint.outputFormat != "" {
		args = append(args, "--output", lint.outputFormat)
	}

	if lint.reject != nil {
		args = append(args, "--reject", strings.Join(lint.reject, ","))
	}

	if lint.summary {
		args = append(args, "--summary")
	}

	if lint.verbose {
		args = append(args, "--verbose")
	}

	args = append(args, "/manifests")

	ctr := dag.Container().
		From("ghcr.io/yannh/kubeconform:v0.6.6").
		WithDirectory("/manifests", manifests)

	for i, schema := range lint.schemas {
		path := "/schemas/" + strconv.Itoa(i)
		ctr = ctr.WithDirectory(path, schema.dir)
		args = append(args, "--schema-location", path+"/"+schema.pattern)
	}

	return ctr.WithExec(args).
		Stdout(ctx)
}

func (m *Kubeconform) Lint(ctx context.Context, manifests *Directory) (string, error) {
	return Lint{}.Lint(ctx, manifests)
}
