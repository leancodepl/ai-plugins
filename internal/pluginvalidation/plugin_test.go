package pluginvalidation

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testPluginName = "demo-plugin"

// fixture builds a minimal but fully valid repository on disk, so each test can
// break exactly one thing and attribute the resulting errors to that change.
type fixture struct {
	t    *testing.T
	root string
}

func newFixture(t *testing.T) *fixture {
	t.Helper()
	f := &fixture{t: t, root: t.TempDir()}
	f.write("plugin-tooling.json", `{ "pluginRoot": "plugins" }`)
	f.write(".claude-plugin/marketplace.json", `{
  "plugins": [
    { "name": "`+testPluginName+`", "source": "./plugins/`+testPluginName+`" },
    { "name": "external-plugin", "source": { "source": "github", "repo": "example/external" } }
  ]
}`)
	f.write(
		"plugins/"+testPluginName+"/.claude-plugin/plugin.json",
		`{ "name": "`+testPluginName+`", "version": "0.1.0", "skills": "./skills/" }`,
	)
	f.write("plugins/"+testPluginName+"/README.md", "# demo\n")

	f.skill(testPluginName+"-usage", testPluginName+"-usage")
	f.skill("do-the-thing", "do-the-thing")
	return f
}

func (f *fixture) write(relPath, content string) {
	f.t.Helper()
	path := filepath.Join(f.root, filepath.FromSlash(relPath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		f.t.Fatalf("mkdir %s: %v", relPath, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		f.t.Fatalf("write %s: %v", relPath, err)
	}
}

// skill writes skills/<dir>/SKILL.md declaring `name: <declaredName>`, which the
// tests deliberately let drift apart from <dir>.
func (f *fixture) skill(dir, declaredName string) {
	f.t.Helper()
	f.write(
		"plugins/"+testPluginName+"/skills/"+dir+"/"+skillFile,
		"---\nname: "+declaredName+"\ndescription: Does a thing.\n---\n\n# body\n",
	)
}

func (f *fixture) removeSkill(dir string) {
	f.t.Helper()
	path := filepath.Join(f.root, "plugins", testPluginName, skillsDir, dir)
	if err := os.RemoveAll(path); err != nil {
		f.t.Fatalf("remove skill %s: %v", dir, err)
	}
}

func (f *fixture) validate() Report {
	f.t.Helper()
	return Validate(f.root)
}

func TestValidateAcceptsWellNamedSkills(t *testing.T) {
	report := newFixture(t).validate()

	if !report.OK() {
		t.Fatalf("expected the fixture to be valid, got errors:\n%s", strings.Join(report.Errors, "\n"))
	}
	if report.PluginCount != 1 {
		t.Errorf("PluginCount = %d, want 1", report.PluginCount)
	}
}

func TestValidateRejectsSkillNamedAfterItsPlugin(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		wantSub []string
	}{
		{
			name:    "exact match",
			dir:     testPluginName,
			wantSub: []string{"skill name must differ from the plugin name", "/demo-plugin:demo-plugin"},
		},
		{
			name:    "differs only by case",
			dir:     "Demo-Plugin",
			wantSub: []string{"skill name must differ from the plugin name"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFixture(t)
			f.removeSkill("do-the-thing")
			f.skill(tt.dir, tt.dir)

			report := f.validate()
			if report.OK() {
				t.Fatal("expected a validation error, got none")
			}
			joined := strings.Join(report.Errors, "\n")
			for _, want := range tt.wantSub {
				if !strings.Contains(joined, want) {
					t.Errorf("errors missing %q:\n%s", want, joined)
				}
			}
		})
	}
}

// A stuttering name is only caught for good if the declared name cannot drift
// away from the directory the check reads.
func TestValidateRejectsSkillFrontmatterNameMismatch(t *testing.T) {
	tests := []struct {
		name         string
		declaredName string
		wantSub      string
	}{
		{
			name:         "drifts from directory",
			declaredName: testPluginName,
			wantSub:      `frontmatter ` + "`name`" + ` must match the skill directory ("do-the-thing"), got "demo-plugin"`,
		},
		{
			name:         "empty",
			declaredName: `""`,
			wantSub:      "frontmatter must set a non-empty `name`",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFixture(t)
			f.skill("do-the-thing", tt.declaredName)

			report := f.validate()
			if report.OK() {
				t.Fatal("expected a validation error, got none")
			}
			if joined := strings.Join(report.Errors, "\n"); !strings.Contains(joined, tt.wantSub) {
				t.Errorf("errors missing %q:\n%s", tt.wantSub, joined)
			}
		})
	}
}

func TestValidateStillRequiresAUsageSkill(t *testing.T) {
	f := newFixture(t)
	f.removeSkill(testPluginName + "-usage")

	report := f.validate()
	if joined := strings.Join(report.Errors, "\n"); !strings.Contains(joined, "a usage skill is required") {
		t.Errorf("errors missing the usage-skill requirement:\n%s", joined)
	}
}

func TestValidateRejectsSkillWithoutFrontmatter(t *testing.T) {
	f := newFixture(t)
	f.write("plugins/"+testPluginName+"/skills/do-the-thing/"+skillFile, "# no frontmatter\n")

	report := f.validate()
	if joined := strings.Join(report.Errors, "\n"); !strings.Contains(joined, "missing or malformed YAML frontmatter") {
		t.Errorf("errors missing the frontmatter requirement:\n%s", joined)
	}
}
