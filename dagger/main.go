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
	// +private
	Schemas []*Schema
	// +private
	ExitOnError bool
	// +private
	IgnoreFilenamePattern []string
	// +private
	IgnoreMissingSchemas bool
	// +private
	KubernetesVersion string
	// +private
	Parallelism *int
	// +private
	OutputFormat string
	// +private
	Reject []string
	// +private
	SkipGVK []string
	// +private
	Strict bool
	// +private
	Summary bool
	// +private
	Verbose bool
}

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
	lint.Schemas = append(lint.Schemas, schema)
	return lint
}

func (m *Kubeconform) WithSchemas(schema *Schema) Lint {
	return Lint{}.WithSchemas(schema)
}

func (lint Lint) WithExitOnError() Lint {
	lint.ExitOnError = true
	return lint
}

func (m *Kubeconform) WithExitOnError() Lint {
	return Lint{}.WithExitOnError()
}

func (lint Lint) WithIgnoreFilenamePattern(pattern string) Lint {
	lint.IgnoreFilenamePattern = append(lint.IgnoreFilenamePattern, pattern)
	return lint
}

func (m *Kubeconform) WithIgnoreFilenamePattern(pattern string) Lint {
	return Lint{}.WithIgnoreFilenamePattern(pattern)
}

func (lint Lint) WithIgnoreMissingSchemas() Lint {
	lint.IgnoreMissingSchemas = true
	return lint
}

func (m *Kubeconform) WithIgnoreMissingSchemas() Lint {
	return Lint{}.WithIgnoreMissingSchemas()
}

func (lint Lint) WithKubernetesVersion(version string) Lint {
	lint.KubernetesVersion = version
	return lint
}

func (m *Kubeconform) WithKubernetesVersion(version string) Lint {
	return Lint{}.WithKubernetesVersion(version)
}

func (lint Lint) WithParallelism(n int) Lint {
	lint.Parallelism = &n
	return lint
}

func (m *Kubeconform) WithParallelism(n int) Lint {
	return Lint{}.WithParallelism(n)
}

func (lint Lint) WithOutputFormat(format string) Lint {
	lint.OutputFormat = format
	return lint
}

func (m *Kubeconform) WithOutputFormat(format string) Lint {
	return Lint{}.WithOutputFormat(format)
}

func (lint Lint) WithRejectGVKs(gvks []string) Lint {
	lint.Reject = append(lint.Reject, gvks...)
	return lint
}

func (m *Kubeconform) WithRejectGVKs(gvks []string) Lint {
	return Lint{}.WithRejectGVKs(gvks)
}

func (lint Lint) WithStrict() Lint {
	lint.Strict = true
	return lint
}

func (m *Kubeconform) WithStrict() Lint {
	return Lint{}.WithStrict()
}

func (lint Lint) WithSkipGVK(gvk string) Lint {
	lint.SkipGVK = append(lint.SkipGVK, gvk)
	return lint
}

func (m *Kubeconform) WithSkipGVK(gvk string) Lint {
	return Lint{}.WithSkipGVK(gvk)
}

func (lint Lint) WithSummary() Lint {
	lint.Summary = true
	return lint
}

func (m *Kubeconform) WithSummary() Lint {
	return Lint{}.WithSummary()
}

func (lint Lint) WithVerbose() Lint {
	lint.Verbose = true
	return lint
}

func (m *Kubeconform) WithVerbose() Lint {
	return Lint{}.WithVerbose()
}

//goland:noinspection GoMixedReceiverTypes
func (lint Lint) Lint(ctx context.Context, manifests *Directory) (string, error) {
	var args []string

	if lint.ExitOnError {
		args = append(args, "--exit-on-error")
	}

	if len(lint.IgnoreFilenamePattern) > 0 {
		for _, pattern := range lint.IgnoreFilenamePattern {
			args = append(args, "--ignore-filename-pattern", pattern)
		}
	}

	if lint.IgnoreMissingSchemas {
		args = append(args, "--ignore-missing-schemas")
	}

	if lint.KubernetesVersion != "" {
		args = append(args, "--kubernetes-version", lint.KubernetesVersion)
	}

	if lint.Parallelism != nil {
		args = append(args, "-n", strconv.Itoa(*lint.Parallelism))
	}

	if lint.OutputFormat != "" {
		args = append(args, "--output", lint.OutputFormat)
	}

	if lint.Reject != nil {
		args = append(args, "--reject", strings.Join(lint.Reject, ","))
	}

	if lint.SkipGVK != nil {
		args = append(args, "--skip", strings.Join(lint.SkipGVK, ","))
	}

	if lint.Summary {
		args = append(args, "--summary")
	}

	if lint.Verbose {
		args = append(args, "--Verbose")
	}

	ctr := dag.Container().
		From(kubeconformImage)

	args = append(args, "--schema-location", "default")

	for i, schema := range lint.Schemas {
		path := "/schemas/" + strconv.Itoa(i)
		ctr = ctr.WithDirectory(path, schema.Specs)
		args = append(args, "--schema-location", path+"/"+schema.Pattern)
	}

	args = append(args, ".")

	return ctr.
		WithDirectory("/manifests", manifests).
		WithWorkdir("/manifests").
		WithExec(args).
		Stdout(ctx)
}

//goland:noinspection GoMixedReceiverTypes
func (m *Kubeconform) Lint(ctx context.Context, manifests *Directory) (string, error) {
	return Lint{}.Lint(ctx, manifests)
}
