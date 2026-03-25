## ⚠️ Drawbacks in Existing Tools and How Proposed CLI Addresses Them

| Drawback / Limitation | Existing Tools (iocage, BastilleBSD, CBSD) | runj | How Proposed CLI Solves It |
|----------------------|--------------------------------------------|------|----------------------------|
| Multiple manual steps for basic setup | Requires separate commands for filesystem, config, start, and access | Low-level, no automation | ✅ Single command (`bsdctl create`) automates full setup |
| Complex and inconsistent CLI syntax | Different tools use different commands and flags | Not user-friendly | ✅ Unified and simple CLI (`create`, `ps`, `exec`) |
| High learning curve | Requires understanding jail internals and tool-specific workflows | Requires OCI knowledge | ✅ Minimal commands, easy for beginners |
| Not developer-focused | Designed mainly for system administration | Runtime only | ✅ Designed for developer workflows and quick usage |
| Lack of container-like experience | No simple “run and use” workflow | No high-level interface | ✅ Moves toward Docker-like workflow |
| Hard to quickly prototype environments | Setup takes time and multiple steps | Manual setup required | ✅ Fast environment creation in seconds |
| Heavy reliance on shell scripts | Hard to maintain and extend | Not applicable | ✅ Written in Go → better structure and maintainability |
| Inconsistent output and automation difficulty | Different tools produce different outputs | Low-level only | ✅ Clean and predictable CLI output |
| No abstraction layer over jails | Direct exposure to low-level details | Low-level runtime | ✅ Provides clean abstraction over jail operations |
| Limited onboarding experience | Beginners struggle to start quickly | Not beginner-friendly | ✅ Focus on simplicity and quick start |

---

## 🎯 Summary

Existing tools are powerful but:

- Focus more on system administration than developer experience  
- Require multiple steps and deeper system knowledge  
- Lack a consistent and minimal interface  

The proposed CLI focuses on:

- Simplifying workflows  
- Reducing setup complexity  
- Providing a clean, consistent developer experience  

👉 The goal is not to replace existing tools, but to **make FreeBSD jails more accessible and easier to use**.