# VS Code Setup for WSL Go Development

This project is configured to work with VS Code in a WSL (Windows Subsystem for Linux) environment.

## Prerequisites

1. **WSL installed** - Windows Subsystem for Linux with Ubuntu
2. **Go installed in WSL** - Currently using Go 1.25.4 (installed via snap)
3. **VS Code on Windows** with the following extensions:
   - `ms-vscode-remote.remote-wsl` - Remote - WSL
   - `golang.go` - Go language support

## Quick Setup

### 1. Install Required Extensions

When you open this workspace, VS Code should prompt you to install the recommended extensions. Click "Install All" or install them manually:

- Open VS Code
- Press `Ctrl+Shift+X` to open Extensions
- Search for and install:
  - "Remote - WSL" by Microsoft
  - "Go" by Go Team at Google

### 2. Open in WSL

**Important:** You must open this project through WSL for Go tooling to work properly.

```bash
# From Windows Terminal (WSL)
cd ~/workspace/go.workspace/octo.workspace
code .
```

Or from VS Code:
1. Press `F1` or `Ctrl+Shift+P`
2. Type "WSL: Reopen Folder in WSL"
3. Select this folder

### 3. Verify Setup

You should see "WSL: Ubuntu" in the bottom-left green status bar of VS Code.

The Go extension will automatically use the configured settings from `.vscode/settings.json`:
- **GOROOT**: `/snap/go/10984`
- **GOPATH**: `/home/shahid/go`
- **gopls**: `/home/shahid/go/bin/gopls` (language server)
- **dlv**: `/home/shahid/go/bin/dlv` (debugger)

## Features Enabled

### Code Intelligence
- ✅ **IntelliSense** - Auto-completion and suggestions
- ✅ **Go to Definition** - Navigate to function/type definitions
- ✅ **Find References** - Find all usages of symbols
- ✅ **Hover Information** - Documentation on hover

### Code Quality
- ✅ **Format on Save** - Auto-format Go code on save
- ✅ **Organize Imports** - Auto-organize imports on save
- ✅ **Linting** - Real-time code linting (staticcheck via gopls)
- ✅ **Build on Save** - Compile checks on save

### Debugging
- ✅ **Breakpoints** - Set breakpoints in code
- ✅ **Step Through** - Step through code execution
- ✅ **Variable Inspection** - Inspect variables during debugging

## Troubleshooting

### Error: "Failed to find the go binary"

**Solution:** Make sure you're opening the project through WSL:
1. Close VS Code
2. Open WSL terminal
3. Run: `code ~/workspace/go.workspace/octo.workspace`
4. Verify "WSL: Ubuntu" appears in bottom-left corner

### Error: "gopls not found" or Language Server Issues

**Solution:** Reinstall Go tools:
```bash
# In WSL terminal
go install -v golang.org/x/tools/gopls@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
```

Then reload VS Code: `Ctrl+Shift+P` → "Developer: Reload Window"

### Slow IntelliSense or Linting

**Solution:** The first time you open a Go file, gopls needs to index the workspace. This can take 30-60 seconds. Check the output:
- `Ctrl+Shift+U` to open Output panel
- Select "Go" from dropdown
- Wait for "gopls: initialized" message

### Format on Save Not Working

**Solution:** Check your settings:
1. `Ctrl+,` to open Settings
2. Search for "format on save"
3. Ensure "Editor: Format On Save" is checked
4. Ensure "Go" is selected as default formatter for Go files

## Testing the Setup

1. Open any `.go` file
2. You should see:
   - Syntax highlighting
   - IntelliSense when typing
   - No error squiggles (if code is valid)
3. Try:
   - `Ctrl+Space` for auto-completion
   - `F12` to go to definition
   - `Shift+F12` to find references
   - `Ctrl+S` to save (should auto-format)

## Running and Debugging

### Run the Application

**Terminal Method:**
```bash
go run main.go
```

**VS Code Method:**
1. Press `F5` or click "Run and Debug" icon
2. Select "Go: Launch Package"
3. Application starts with debugger attached

### Debug the Application

1. Set breakpoints by clicking left of line numbers
2. Press `F5` to start debugging
3. Use debug controls:
   - `F5` - Continue
   - `F10` - Step Over
   - `F11` - Step Into
   - `Shift+F11` - Step Out
   - `Shift+F5` - Stop

## Additional Resources

- [VS Code Go Extension](https://code.visualstudio.com/docs/languages/go)
- [Debugging Go in VS Code](https://github.com/golang/vscode-go/wiki/debugging)
- [WSL in VS Code](https://code.visualstudio.com/docs/remote/wsl)

## Configuration Files

- `.vscode/settings.json` - Workspace settings for Go
- `.vscode/extensions.json` - Recommended extensions
- `.github/copilot-instructions.md` - AI coding assistant context

Happy coding! 🚀
