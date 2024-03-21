### Build a LSP from using go and standard library

The superior editor is: VS Code (that seen right, use code action to fix it)

- you can run command to trigger code action

```
:lua vim.lsp.buf.code_action()
```

### Issue:

- LSP will not get any info before the first change of the file, might be something in conflict with lazyvim
