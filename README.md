# wireguard_svc

Wireguard tunnel, runnable as a Windows service.

# How to build

```
mage build
```

# How to use

1. `wireguard_svc.exe` has to be installed as a Windows Service

Sample configuration:
```
Service Name:  "WireGuardTunnel$MyPrettyTunnel"
Display Name:  "My Pretty Tunnel Service"
Service Type:  SERVICE_WIN32_OWN_PROCESS
Start Type:    StartAutomatic
Error Control: ErrorNormal,
Dependencies:  [ "Nsi", "TcpIp" ]
Sid Type:      SERVICE_SID_TYPE_UNRESTRICTED
Executable:    "C:\wireguard_svc.exe -service -config-file=C:\wg.conf"
```

Flags `-service` and `-config-file` are mandatory.

See official wireguard-windows README for more information.

https://git.zx2c4.com/wireguard-windows/about/embeddable-dll-service/README.md

2. The service can now be controlled via Windows Service Manager (Start/Stop)