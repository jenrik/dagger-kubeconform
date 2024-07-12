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

type Schema struct {
	Pattern string
	Specs   *Directory
}

type Lint struct {
	schemas               []*Schema
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

//goland:noinspection GoUnusedConst
const (
	OutputFormatJson   = "json"
	OutputFormatJunit  = "junit"
	OutputFormatPretty = "pretty"
	OutputFormatTap    = "tap"
	OutputFormatText   = "text"
)

const kubeconformImage = "ghcr.io/yannh/kubeconform:v0.6.6@sha256:e4c69e6966a4842196ad3babc6f4c869d4ee51dc306fcf012faf10b25bb63a9c"

//goland:noinspection GoUnusedConst
const CRDSchemaPattern = "{{.Group}}_{{.ResourceKind}}_{{.ResourceAPIVersion}}.json"

func crdSchemaContainer() (*Container, error) {
	ctr := dag.Apko().Wolfi([]string{"bash", "curl", "git", "python3", "py3-pip", "yq"}).
		WithExec([]string{"pip", "install", "pyyaml"})

	// Download the openapi2jsonschema.py script and return a dagger *File
	openapi2jsonschemaScript := dag.HTTP("https://raw.githubusercontent.com/yannh/kubeconform/6ae8c45bc156ceeb1d421e9b217cfc0c7ba5828d/scripts/openapi2jsonschema.py")

	ctr = ctr.
		WithFile("/bin/openapi2jsonschema.py", openapi2jsonschemaScript, ContainerWithFileOpts{Permissions: 0750})
	return ctr, nil
}

//goland:noinspection ALL
func (m *Kubeconform) CRD_To_Schema(crdsDir *Directory) (*Schema, error) {
	ctx := context.Background()
	ctr, err := crdSchemaContainer()
	if err != nil {
		return nil, err
	}

	entries, err := crdsDir.Entries(ctx)
	if err != nil {
		return nil, err
	}

	var args = []string{"/bin/openapi2jsonschema.py"}
	for _, entry := range entries {
		args = append(args, "/crds/"+entry)
	}

	output := dag.Directory()
	ctr = ctr.
		WithDirectory("/crds", crdsDir).
		WithWorkdir("/output").
		WithEnvVariable("FILENAME_FORMAT", "{fullgroup}_{kind}_{version}").
		WithExec(args)

	output = output.WithDirectory("/", ctr.Directory("/output"))

	return &Schema{
		Pattern: CRDSchemaPattern,
		Specs:   output,
	}, nil
}

func (lint Lint) WithSchemas(schema *Schema) Lint {
	lint.schemas = append(lint.schemas, schema)
	return lint
}

func (m *Kubeconform) WithSchemas(schema *Schema) Lint {
	return Lint{}.WithSchemas(schema)
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

func (lint Lint) SkipGVK(gvk string) Lint {
	lint.skipGVK = append(lint.skipGVK, gvk)
	return lint
}

func (m *Kubeconform) SkipGVK(gvk string) Lint {
	return Lint{}.SkipGVK(gvk)
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

	if lint.skipGVK != nil {
		args = append(args, "--skip", strings.Join(lint.skipGVK, ","))
	}

	if lint.summary {
		args = append(args, "--summary")
	}

	if lint.verbose {
		args = append(args, "--verbose")
	}

	ctr := dag.Container().
		From(kubeconformImage)

	args = append(args, "--schema-location", "default")

	for i, schema := range lint.schemas {
		path := "/schemas/" + strconv.Itoa(i)
		ctr = ctr.WithDirectory(path, schema.Specs)
		args = append(args, "--schema-location", path+"/"+schema.Pattern)
	}

	args = append(args, "/manifests")

	return ctr.
		WithDirectory("/manifests", manifests).
		WithExec(args).
		Stdout(ctx)
}

func (m *Kubeconform) Lint(ctx context.Context, manifests *Directory) (string, error) {
	return Lint{}.Lint(ctx, manifests)
}
