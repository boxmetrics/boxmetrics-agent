# Response Schema for Network informations

## Network

| Key        | Type   | Description                                                         |
| ---------- | ------ | ------------------------------------------------------------------- |
| usage      | object | Global network usage. See [network IO counter](#network-io-counter) |
| interfaces | Array  | Array of [network interface](#network-interfaces)                   |

### Network Interfaces

| Key          | Type   | Description                                                    |
| ------------ | ------ | -------------------------------------------------------------- |
| mtu          | string | Interface MTU                                                  |
| name         | string | Interface name                                                 |
| hardwareaddr | string | Interface MAC address                                          |
| flags        | array  | Interface flags                                                |
| addrs        | array  | Interface IP addresses                                         |
| usage        | object | Interface usage. See [network IO counter](#network-io-counter) |

### Network IO Counters

| Key             | Type   | Description                        |
| --------------- | ------ | ---------------------------------- |
| name            | string | Interface name                     |
| bytesSent       | string | Number of bytes sent               |
| bytesRecv       | string | Number of bytes receive            |
| bytesSentPerSec | string | Number of bytes sent per second    |
| bytesRecvPerSec | string | Number of bytes receive per second |
| packetsSent     | string | Number of packets sent             |
| packetsRecv     | string | Number of packets receive          |
| errin           | string | Number of input error              |
| errout          | string | Number of output error             |
| dropin          | string | Number of intput drop              |
| dropout         | string | Number of output drop              |
| fifoin          | string | Number of input fifo               |
| fifoout         | string | Number of output fifo              |

## Connection

| Key        | Type   | Description        |
| ---------- | ------ | ------------------ |
| fd         | number | File descriptor Id |
| family     | string | Connection family  |
| type       | string | Connection type    |
| localaddr  | string | Local address      |
| remoteaddr | string | Remote address     |
| status     | string | Connection status  |
| uids       | Array  | Connection UIDs    |
| pid        | number | Connection PID     |
