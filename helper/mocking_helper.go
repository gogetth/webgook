package helper

// ScriptRunnerMockup use for mocking script runner
type ScriptRunnerMockup struct {
}

// RunScript method for mocking script result
func (s ScriptRunnerMockup) RunScript(scriptPath string) error {
	return nil
}
