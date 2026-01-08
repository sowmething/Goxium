# Goxium
System-wide browser fingerprint hiding utility written in Go

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

---

## Troubleshooting

### Certificate or Key File Not Found
**Issue:** Goxium fails to load the certificate or HTTPS interception does not work.

**Cause:**  
The certificate (`goxium_ca.crt`) and private key (`goxium_ca.key`) are not located in the same directory as the executable.

**Solution:**  
Move both files into the same directory where `goxium.exe` is located, then restart Goxium.

---

### Invalid or Mismatched Certificate and Key
**Issue:** TLS errors occur or HTTPS traffic is not intercepted.

**Cause:**  
The certificate and private key do not match, or they are invalid/corrupted.

**Solution:**  
Re-download the correct certificate and key pair, or regenerate them, and ensure they belong to the same CA.

---

### Application Not Started as Administrator
**Issue:** The Root CA certificate is not added to the Trusted Root Certification Authorities store.

**Cause:**  
Goxium was not started with administrator privileges.

**Solution:**  
Close Goxium and restart it using **Run as Administrator**.

---

### Corrupted or Invalid Binary
**Issue:** Goxium crashes on startup or behaves unexpectedly.

**Cause:**  
The binary file is corrupted or was built incorrectly.

**Solution:**  
Rebuild the project using `build.bat`, or re-download the precompiled binaries from the repository and try again.

## FAQ

### Is Goxium a VPN?
No.  
Goxium is **not a VPN** and does not change or hide your IP address.  
It operates as a **local HTTP/HTTPS proxy** that modifies request headers.

---

### Does Goxium provide anonymity?
Kinda.  
Goxium does not hide traffic or bypass network restrictions.  
Its purpose is to **reduce/fake fingerprinting vectors** by normalizing HTTP headers. This will make malicious sites dont recognise you.

---

### Does Goxium decrypt HTTPS traffic?
Yes, locally.  
Goxium uses a custom Root CA to intercept HTTPS traffic **on your own system only**.  
All traffic remains local and is not sent to any external server.

---

### Is it safe to install the Root CA?
Yes.
Installing a Root CA grants high trust.  
You should only use Goxium on systems you fully control and trust.

Goxium automatically removes the Root CA when the application exits.

---

### Why does Goxium require administrator privileges?
Administrator privileges are required to:
- Add the Root CA to the Windows Trusted Root store
- Remove the certificate on shutdown

Without admin rights, HTTPS interception will not work.

---

### Does Goxium log or store traffic?
No.  
Goxium does not store traffic or send data externally.  
Requests are only logged to the console for debugging purposes and doesnt sent to anything.

---

### Can Goxium break websites?
Possibly.  
Some websites rely on specific headers for functionality or security checks.  
Removing or modifying headers may cause unexpected behavior.

---

### Is Goxium detectable by websites?
Rare.  
Advanced anti-bot or anti-fraud systems may detect modified TLS or header behavior.  
Goxium aims to **reduce**, not eliminate, fingerprinting.

---

### Can I use Goxium in production environments?
Goxium is intended for:
- Research
- Development
- Privacy testing
- Debugging

### Goxium doesnt work expected. What should i do?
Look at Troubleshooting

Using it in production environments is **not recommended** without thorough testing.
