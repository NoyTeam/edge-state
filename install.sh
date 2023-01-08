#!/bin/bash
wget https://github.com/ArsFy/edge-state/releases/download/1.0/edge-state
chmod 777 edge-state
pwd = $(pwd)
cat > /etc/systemd/system/edge-state.service <<EOF
[Unit]
Description=edge-state
[Service]
Type=simple
WorkingDirectory=/root/
ExecStart=$pwd/edge-state
Restart=always
RestartSec=5
StartLimitInterval=3
RestartPreventExitStatus=137
 
[Install]
WantedBy=multi-user.target
EOF
service edge-state start