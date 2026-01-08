# Goxium
System-wise browser fingerprint hiding utility written in Go

## Build

You can run Goxium in two different ways:

### Option 1: Build from source
Clone the repository and run the build script:

```bat
git clone https://github.com/sowmething/Goxium.git
cd Goxium
build.bat
```
files will bult to bin folder.

### Option 2: Use prebuilt binaries

Precompiled binaries are available in the bin directory for the following architectures:
- x64
- x86
- arm64

No build process is required when using these binaries.

## Installation & Setup

### 1. Run as Administrator

Start **Goxium** with administrator privileges.  
This is required to install and remove the Root CA certificate from the system trust store.

---

### 2. Configure System Proxy Settings

Open your system proxy settings and enter the following values:

- **IP Address:** `127.0.0.1`
- **Port:** `19381`

---

### 3. Proxy Settings Screenshots

**Enable and edit proxy settings:**

![Proxy Edit](https://github.com/sowmething/Goxium/blob/main/edit.png)

**Set proxy IP and port:**

![Proxy Set](https://github.com/sowmething/Goxium/blob/main/set.png)

---

Once the proxy is configured, all system HTTP and HTTPS traffic will be routed through Goxium. (YOU MUST RUN GOXIUM!!!)
