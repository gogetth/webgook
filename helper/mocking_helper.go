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

// VerifySecretKey method for mocking verification to true
func (s ScriptRunnerMockup) VerifySecretKey(secretKeyFromCommandLine, secretKey string) (bool, error) {
	return true, nil
}
