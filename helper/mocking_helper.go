package helper

// ScriptRunnerMockup use for mocking script runner
type ScriptRunnerMockup struct {
}

// CheckScriptPath method for mocking script path
func (s ScriptRunnerMockup) CheckScriptPath(scriptPath string) (string, error) {
	return "", nil
}

// RunScript method for mocking script result
func (s ScriptRunnerMockup) RunScript(scriptPath string) error {
	return nil
}
