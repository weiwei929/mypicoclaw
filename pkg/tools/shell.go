package tools

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type ExecTool struct {
	workingDir          string
	timeout             time.Duration
	denyPatterns        []*regexp.Regexp
	warnPatterns        []*regexp.Regexp
	allowPatterns       []*regexp.Regexp
	restrictToWorkspace bool
}

func NewExecTool(workingDir string) *ExecTool {
	// 🔴 Blocked: immediately rejected, never executed
	denyPatterns := []*regexp.Regexp{
		// Destructive file operations
		regexp.MustCompile(`\brm\s+-[rf]{1,2}\s+/`),               // rm -rf / (root path)
		regexp.MustCompile(`\bdel\s+/[fq]\b`),                      // Windows del /f
		regexp.MustCompile(`\brmdir\s+/s\b`),                       // Windows rmdir /s
		// Disk wiping
		regexp.MustCompile(`\b(format|mkfs|diskpart)\b\s`),
		regexp.MustCompile(`\bdd\s+if=`),
		regexp.MustCompile(`>\s*/dev/sd[a-z]\b`),
		// System control
		regexp.MustCompile(`\b(shutdown|reboot|poweroff|halt)\b`),
		regexp.MustCompile(`\bsystemctl\s+(stop|disable)\s+mypicoclaw`), // Don't let it stop itself
		// Fork bomb
		regexp.MustCompile(`:\(\)\s*\{.*\};\s*:`),
		// System file overwrite
		regexp.MustCompile(`>\s*/etc/(passwd|shadow|sudoers|fstab)`),
		// Firewall flush (locked out of VPS)
		regexp.MustCompile(`\biptables\s+-F`),
		// Dangerous permission changes on root
		regexp.MustCompile(`\bchmod\s+(-R\s+)?777\s+/\s*$`),
		// User deletion
		regexp.MustCompile(`\buserdel\b`),
	}

	// 🟡 Warned: logged as warning but still executed
	warnPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\brm\b`),                               // Any file deletion
		regexp.MustCompile(`\b(kill|pkill|killall)\b`),              // Process killing
		regexp.MustCompile(`\b(apt|yum|dnf|pacman)\s+install\b`),    // Package installation
		regexp.MustCompile(`\bcurl\b.*\|\s*(ba)?sh`),                // Pipe-to-shell
		regexp.MustCompile(`\bwget\b.*\|\s*(ba)?sh`),                // Pipe-to-shell
		regexp.MustCompile(`\bchmod\b`),                             // Permission changes
		regexp.MustCompile(`\bchown\b`),                             // Ownership changes
		regexp.MustCompile(`\bsystemctl\s+(restart|start|enable)\b`), // Service management
		regexp.MustCompile(`\bcrontab\b`),                           // Scheduled tasks
		regexp.MustCompile(`\bnohup\b.*&`),                          // Background daemons
	}

	return &ExecTool{
		workingDir:          workingDir,
		timeout:             60 * time.Second,
		denyPatterns:        denyPatterns,
		warnPatterns:        warnPatterns,
		allowPatterns:       nil,
		restrictToWorkspace: false,
	}
}

func (t *ExecTool) Name() string {
	return "exec"
}

func (t *ExecTool) Description() string {
	return "Execute a shell command and return its output. Use with caution."
}

func (t *ExecTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"command": map[string]interface{}{
				"type":        "string",
				"description": "The shell command to execute",
			},
			"working_dir": map[string]interface{}{
				"type":        "string",
				"description": "Optional working directory for the command",
			},
		},
		"required": []string{"command"},
	}
}

func (t *ExecTool) Execute(ctx context.Context, args map[string]interface{}) (string, error) {
	command, ok := args["command"].(string)
	if !ok {
		return "", fmt.Errorf("command is required")
	}

	cwd := t.workingDir
	if wd, ok := args["working_dir"].(string); ok && wd != "" {
		cwd = wd
	}

	if cwd == "" {
		wd, err := os.Getwd()
		if err == nil {
			cwd = wd
		}
	}

	if guardError := t.guardCommand(command, cwd); guardError != "" {
		return fmt.Sprintf("Error: %s", guardError), nil
	}

	cmdCtx, cancel := context.WithTimeout(ctx, t.timeout)
	defer cancel()

	cmd := exec.CommandContext(cmdCtx, "sh", "-c", command)
	if cwd != "" {
		cmd.Dir = cwd
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	output := stdout.String()
	if stderr.Len() > 0 {
		output += "\nSTDERR:\n" + stderr.String()
	}

	if err != nil {
		if cmdCtx.Err() == context.DeadlineExceeded {
			return fmt.Sprintf("Error: Command timed out after %v", t.timeout), nil
		}
		output += fmt.Sprintf("\nExit code: %v", err)
	}

	if output == "" {
		output = "(no output)"
	}

	maxLen := 10000
	if len(output) > maxLen {
		output = output[:maxLen] + fmt.Sprintf("\n... (truncated, %d more chars)", len(output)-maxLen)
	}

	return output, nil
}

func (t *ExecTool) guardCommand(command, cwd string) string {
	cmd := strings.TrimSpace(command)
	lower := strings.ToLower(cmd)

	// 🔴 Check deny patterns (block immediately)
	for _, pattern := range t.denyPatterns {
		if pattern.MatchString(lower) {
			log.Printf("[SECURITY] ⛔ BLOCKED command: %s (matched: %s)", cmd, pattern.String())
			return "⛔ Command blocked by safety guard (dangerous pattern detected)"
		}
	}

	// 🟡 Check warn patterns (log warning but allow)
	for _, pattern := range t.warnPatterns {
		if pattern.MatchString(lower) {
			log.Printf("[SECURITY] ⚠️ RISKY command allowed: %s (matched: %s)", cmd, pattern.String())
			break // Only log once even if multiple patterns match
		}
	}

	// Check allowlist if configured
	if len(t.allowPatterns) > 0 {
		allowed := false
		for _, pattern := range t.allowPatterns {
			if pattern.MatchString(lower) {
				allowed = true
				break
			}
		}
		if !allowed {
			return "Command blocked by safety guard (not in allowlist)"
		}
	}

	if t.restrictToWorkspace {
		if strings.Contains(cmd, "..\\") || strings.Contains(cmd, "../") {
			return "Command blocked by safety guard (path traversal detected)"
		}

		cwdPath, err := filepath.Abs(cwd)
		if err != nil {
			return ""
		}

		pathPattern := regexp.MustCompile(`[A-Za-z]:\\[^\\\"']+|/[^\s\"']+`)
		matches := pathPattern.FindAllString(cmd, -1)

		for _, raw := range matches {
			p, err := filepath.Abs(raw)
			if err != nil {
				continue
			}

			rel, err := filepath.Rel(cwdPath, p)
			if err != nil {
				continue
			}

			if strings.HasPrefix(rel, "..") {
				return "Command blocked by safety guard (path outside working dir)"
			}
		}
	}

	return ""
}

func (t *ExecTool) SetTimeout(timeout time.Duration) {
	t.timeout = timeout
}

func (t *ExecTool) SetRestrictToWorkspace(restrict bool) {
	t.restrictToWorkspace = restrict
}

func (t *ExecTool) SetAllowPatterns(patterns []string) error {
	t.allowPatterns = make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		re, err := regexp.Compile(p)
		if err != nil {
			return fmt.Errorf("invalid allow pattern %q: %w", p, err)
		}
		t.allowPatterns = append(t.allowPatterns, re)
	}
	return nil
}
