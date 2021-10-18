



cobra init --pkg-name github.com/gregsidelinger/pulse-oximeter




/etc/systemd/system/pulse-oximeter.service
```
[Unit]
Description=Puslse Oximeter
After=network.target

[Service]
User=gate
Restart=always
Type=simple

PermissionsStartOnly=true
ExecStartPre=/bin/chmod a+rw /dev/ttyUSB0

ExecStart=pulse-oximeter monitor -d /dev/ttyUSB0

[Install]
WantedBy=multi-user.target
```
