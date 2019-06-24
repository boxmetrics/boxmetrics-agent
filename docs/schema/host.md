# Response Schema for Host informations

## Host

| Key                  | Type   | Description                              |
| -------------------- | ------ | ---------------------------------------- |
| hostname             | string | System hostname                          |
| uptime               | string | System uptime                            |
| bootTime             | string | System boot date                         |
| procs                | string | Number of processes                      |
| os                   | string | Os type (ex: "linux")                    |
| platform             | string | Platform name (ex: "ubuntu", "arch")     |
| platformFamily       | string | Platform family (ex: "debian")           |
| platformVersion      | string | Platform version (ex: "Ubuntu 13.10")    |
| kernerlVersion       | string | Kernel version number                    |
| virtualizationSystem | string | Virtualization system (ex: "LXC")        |
| virtualizationRole   | string | Virtualization role (ex: "guest"/"host") |
| hostId               | string | Host identifier                          |
| sensors              | object | Array of [sensor stat](#sensor-stat)     |

### Sensor Stat

| Key         | Type   | Description        |
| ----------- | ------ | ------------------ |
| sensorKey   | string | Sensor identifier  |
| temperature | string | Sensor temperature |

## User

| Key      | Type   | Description                 |
| -------- | ------ | --------------------------- |
| uid      | string | User id                     |
| gid      | string | Primary group id            |
| username | string | Login name                  |
| name     | string | User's real or display name |
| homeDir  | string | User's home directory       |

## Session

| Key      | Type   | Description               |
| -------- | ------ | ------------------------- |
| user     | string | User name                 |
| terminal | string | User default terminal     |
| host     | string | Session host              |
| started  | int    | Number of session started |
