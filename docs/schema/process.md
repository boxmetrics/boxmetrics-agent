# Response Schema for Process informations

## Process

| Key         | Type    | Description                                                            |
| ----------- | ------- | ---------------------------------------------------------------------- |
| pid         | string  | Process PID                                                            |
| name        | string  | Process name                                                           |
| username    | string  | Process user name                                                      |
| status      | string  | Process status                                                         |
| uids        | Array   | Process UIDs                                                           |
| gids        | Array   | Process GIDs                                                           |
| terminal    | string  | Process associated terminal                                            |
| cwd         | string  | Process current working directory                                      |
| exe         | string  | Process executable path                                                |
| cmdArgs     | Array   | Process command line arguments                                         |
| ppid        | string  | Process parent PID                                                     |
| parent      | string  | Process parent info. See [process light](#process-light)               |
| children    | Array   | Process children. See [process light](#process-light)                  |
| cpu         | object  | [Process cpu usage](#process-cpu-info)                                 |
| memory      | object  | [Process memory usage](#process-memory-info)                           |
| network     | object  | [Process network usage](./network.md#network-io-counter)               |
| disk        | object  | [Process disk usage](#process-disk-info)                               |
| connections | Array   | Process connections. See [network connection](./network.md#connection) |
| numFds      | string  | Number of file descriptor                                              |
| openFiles   | Array   | Process opened files . See [process file](#process-file)               |
| numThread   | string  | Number of threads                                                      |
| threads     | Array   | Process threads. See [cpu times](./cpu.md#cpu-times)                   |
| limits      | Array   | Process limits. See [process limit](#process-limit)                    |
| isRunning   | boolean | If process is running                                                  |
| background  | boolean | If process is in background                                            |
| foreground  | boolean | If process is in foreground                                            |
| createTime  | string  | Process create time                                                    |

### Process CPU Info

| Key     | Type   | Description                            |
| ------- | ------ | -------------------------------------- |
| percent | srting | Percentage of CPU used                 |
| times   | object | [cpu times](./cpu.md#cpu-times) object |

### Process Memory Info

| Key     | Type   | Description                                          |
| ------- | ------ | ---------------------------------------------------- |
| percent | string | Percentage of memory used                            |
| usage   | object | [Process Memory Usage](#process-memory-usage) object |

### Process Memory Usage

| Key    | Type   | Description        |
| ------ | ------ | ------------------ |
| rss    | string | RSS memory used    |
| vms    | string | VMS memory used    |
| data   | string | Data memory used   |
| stack  | string | Stack memory used  |
| locked | string | Locked memory used |
| swap   | string | Swap memory used   |

### Process Disk Info

| Key              | Type   | Description                      |
| ---------------- | ------ | -------------------------------- |
| readCount        | string | Number of read                   |
| writeCount       | string | Number of write                  |
| readBytes        | string | Number of bytes read             |
| writeBytes       | string | Number of bytes write            |
| readBytesPerSec  | string | Number of bytes read per second  |
| writeBytesPerSec | string | Number of bytes write per second |

### Process File

| Key  | Type   | Description                |
| ---- | ------ | -------------------------- |
| path | string | Path of the file           |
| fd   | number | File descriptor identifier |

### Process Limit

| Key      | Type   | Description   |
| -------- | ------ | ------------- |
| resource | string | Resource name |
| soft     | string | Soft limit    |
| hard     | string | Hard limit    |
| used     | string | Resource used |

## Process Light

| Key        | Type   | Description                                  |
| ---------- | ------ | -------------------------------------------- |
| pid        | string | Process PID                                  |
| name       | string | Process name                                 |
| username   | string | Process user name                            |
| status     | string | Process status                               |
| uids       | Array  | Process UIDs                                 |
| gids       | Array  | Process GIDs                                 |
| terminal   | string | Process associated terminal                  |
| cwd        | string | Process current working directory            |
| exe        | string | Process executable path                      |
| cmdArgs    | Array  | Process command line arguments               |
| ppid       | string | Process parent PID                           |
| cpu        | object | [Process cpu usage](#process-cpu-info)       |
| memory     | object | [Process memory usage](#process-memory-info) |
| createTime | string | Process create time                          |
