# /bin/bash

## Build project
make build

## Create systemd service
SERVICE_NAME="gogenovpn"

IS_ACTIVE=$(systemctl is-active $SERVICE_NAME)
if [ "$IS_ACTIVE" == "active" ]; then
  # restart the service
  echo "Service is running"
  echo "Restarting service"
  systemctl restart $SERVICE_NAME
  echo "Service restarted"
else 
  # create service file
  echo "Creating service file"
  cat > /etc/systemd/system/${SERVICE_NAME}.service << EOF
[Unit]
Description=microservice api for generating openvpn files
After=network.target

[Service]
Environment=APP_ENV=production
Environment=AUTH_SECRET=valid-key
ExecStart=$PWD/bin/gogetovpn
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF
  ## Restart daemon, enable, and start service
  systemctl daemon-reload
  systemctl enable ${SERVICE_NAME}
  systemctl start ${SERVICE_NAME}
  echo "Service Started"
fi

exit 0