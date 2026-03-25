##  Unified Comparison of FreeBSD Jail Tools

The following table compares existing tools and highlights practical differences in features, usability, and workflows.

| Feature / Aspect | iocage | BastilleBSD | CBSD | runj | Proposed CLI (bsdctl) |
|------------------|--------|------------|------|------|------------------------|
| **Primary Purpose** | Jail manager | Jail manager | Virtualization suite | OCI runtime | Developer-focused CLI |
| **Target Users** | Sysadmins | Sysadmins | Advanced users | Low-level users | Developers |
| **Jail Creation (Commands)** | `iocage create -r 14.0-RELEASE -n myjail ip4_addr="lo1\|127.0.0.2"` | `bastille create myjail 14.0-RELEASE 127.0.0.2` | `cbsd jcreate jname=myjail` | Manual OCI setup | `bsdctl create myjail` |
| **Start/Stop Jails** | `iocage start myjail` | `bastille start myjail` | `cbsd jstart jname=myjail` | `runj start` | `bsdctl start / stop` |
| **Access Jail Shell** | `iocage console myjail` | `bastille console myjail` | `cbsd jlogin jname=myjail` | Not direct | `bsdctl exec myjail` |
| **List Jails** | `iocage list` | `bastille list` | `cbsd jls` | Not provided | `bsdctl ps` |
| **Automation of Setup** | Partial | Partial | No | No | Full (filesystem + config) |
| **Command Simplicity** | Complex (flag-based) | Moderate | Complex | Low-level | Simple and minimal |
| **Consistency of CLI** | Inconsistent | Inconsistent | Inconsistent | Inconsistent | Consistent |
| **Developer-Friendly Workflow** | Limited | Limited | Limited | Not applicable | Strong |
| **Container-like UX** | Not available | Not available | Not available | Not available | Supported |
| **Learning Curve** | High | Medium | High | High | Low |
| **Maintainability (Codebase)** | Shell-heavy | Shell-heavy | Mixed | Structured | Go-based (clean) |

## ⚠️ Key Drawbacks in Existing Tools

- Require multiple steps for simple tasks  
- Focus more on system administration than developer experience  
- Inconsistent command syntax across tools  
- Lack of a minimal and unified CLI  
- No direct container-like workflow  
- Heavy reliance on shell scripting (hard to maintain at scale)  

---

##  How Proposed CLI Addresses These Issues

- Provides **single-command workflows** (e.g., `bsdctl create`)  
- Offers **consistent and predictable CLI design**  
- Focuses on **developer usability and quick setup**  
- Reduces manual configuration steps  
- Moves toward a **container-like experience on top of jails**  
- Built in Go for **better structure and maintainability**  

---

##  Example Workflow Comparison

### Existing Workflow

```bash
mkdir /usr/jails/test
tar -xf base.txz -C /usr/jails/test
nano /etc/jail.conf
service jail start
jexec test sh
```

---

### Proposed Workflow

```bash
bsdctl create test
bsdctl exec test
```

 Simplified, faster, and more intuitive.