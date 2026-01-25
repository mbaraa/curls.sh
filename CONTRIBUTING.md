# Want to contribute?

First off, thanks for taking the time to contribute! 😇

---

### Bug report:

Did you find a bug?

- Open an issue with full bug description and example, and I'll see what I can do about it 🤠
- If you can fix the bug you've found, consider making a PR, and I'll see if your fix fits 👀
- Make sure that the PR or the issue clearly describes the problem and solution 👌

---

### New scripts:

1. Fork the repo
2. Add the new script under [scripts](/scripts)
3. Test it locally
4. Open a pull request
5. Wait for it...

### New script type:

Under [scripts/script.go](/scripts/script.go) you can see a `Script` interface which is currently implemented by `bashScript` under [scripts/bash.go](/scripts/bash.go) which represents a script fully compatible by [Bash](https://www.gnu.org/software/bash/manual/bash.html), feel free to create more implementations, like [ZSH](https://www.zsh.org/) or [PowerShell](https://learn.microsoft.com/en-us/powershell/)
