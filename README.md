# Go Monitor Service

Simple VPS/SSH monitoring service written in Go.

The service periodically checks TCP availability of a server (for example SSH on port 22) and sends Telegram alerts when the server goes down or comes back online.

---

# Features

* TCP server monitoring
* SSH availability checks
* Telegram notifications
* Recovery alerts
* Stateful monitoring logic
* Linux systemd deployment
* Lightweight and simple

---

# Stack

* Go
* Linux
* TCP networking
* Telegram Bot API
* systemd

---

# How It Works

The service:

1. Connects to a server using TCP
2. Checks if the port is reachable
3. Detects state changes:

   * UP → DOWN
   * DOWN → UP
4. Sends Telegram notifications
5. Runs continuously in the background

---

# Example Alerts

```text
SERVER DOWN
SERVER RECOVERED
```

---

# Installation

## Clone repository

```bash
git clone https://github.com/YOUR_USERNAME/go-monitor-service.git
cd go-monitor-service
```

---

# Configure environment variables

```bash
export BOT_TOKEN="your_bot_token"
export CHAT_ID="your_chat_id"
```

---

# Run locally

```bash
go run main.go
```

---

# Build binary

```bash
go build -o monitor-service
```

---

# Deploy with systemd

## Create directory

```bash
mkdir -p /opt/monitor-service
mv monitor-service /opt/monitor-service/
```

---

## Create systemd service

```bash
nano /etc/systemd/system/monitor.service
```

Paste:

```ini
[Unit]
Description=Go Monitoring Service
After=network.target

[Service]
Type=simple
ExecStart=/opt/monitor-service/monitor-service
Restart=always
RestartSec=5
Environment="BOT_TOKEN=your_bot_token"
Environment="CHAT_ID=your_chat_id"

[Install]
WantedBy=multi-user.target
```

---

# Enable service

```bash
systemctl daemon-reload
systemctl enable monitor
systemctl start monitor
```

---

# Logs

```bash
journalctl -u monitor -f
```

---

# Future Improvements

* Multiple servers monitoring
* HTTP health checks
* Docker support
* Configuration file
* Database for checks history
* Web dashboard
* Prometheus metrics
* Grafana integration

---

# Repository Description

Lightweight Go-based VPS/SSH monitoring service with Telegram alerts and systemd deployment.
