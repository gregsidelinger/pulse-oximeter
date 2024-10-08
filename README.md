



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


Example config file
~/.pulse-oximeter.yaml
```yaml
---
monitor:
  device: /dev/ttyUSB0
  baud-rate: 19022
```



```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    k8s-app: pulse-oximeter
  name: pulse-oximeter
spec:
  externalName: HOSTNAME
  ports:
  - name: metrics
    port: 9100
    protocol: TCP
    targetPort: 9100
  sessionAffinity: None
  type: ExternalName
```


```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  labels:
    k8s-app: pulse-oximeter
  name: pulse-oximeter
spec:
  endpoints:
  - honorLabels: true
    interval: 2s
    path: /metrics
    port: metrics
  namespaceSelector:
    matchNames:
    - krynn
  selector:
    matchLabels:
      k8s-app: pulse-oximeter
```
