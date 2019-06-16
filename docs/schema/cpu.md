# Response Schema for CPU informations

## CPU

| Key   | Type   | Description                                      |
| ----- | ------ | ------------------------------------------------ |
| count | object | Number of processor. See [cpu count](#cpu-count) |
| info  | Array  | Array of [cpu info](#cpu-info)                   |

### CPU Count

| Key      | Type   | Description                  |
| -------- | ------ | ---------------------------- |
| physical | string | Number of physical processor |
| logical  | string | Number of logical processor  |

### CPU Info

| Key     | Type   | Description                      |
| ------- | ------ | -------------------------------- |
| id      | string | A CPU ID                         |
| percent | string | Usage of the CPU in %            |
| times   | object | A [cpu times](#cpu-times) object |

### CPU Times

| Key       | Type   | Description                            |
| --------- | ------ | -------------------------------------- |
| user      | string | Amounts of time of the user work       |
| system    | string | Amounts of time of the system work     |
| idle      | string | Amounts of time of the idle work       |
| nide      | string | Amounts of time of the nice work       |
| iowait    | string | Amounts of time of the iowait work     |
| irq       | string | Amounts of time of the irq work        |
| softirq   | string | Amounts of time of the softirq work    |
| steal     | string | Amounts of time of the steal work      |
| guest     | string | Amounts of time of the guest work      |
| guestNice | string | Amounts of time of the guest nice work |
| total     | string | Total of time for all kinds of work    |

## CPU Hardware Info

| Key        | Type   | Description              |
| ---------- | ------ | ------------------------ |
| cpu        | number | CPU Identifier           |
| vendorId   | string | Vendor Identifier        |
| family     | string | Family of the CPU        |
| model      | string | Model name               |
| stepping   | number | ???                      |
| physicalId | string | Hardware Identifier      |
| coreId     | string | Something about core     |
| cores      | number | Number of cores          |
| modelName  | string | Another model name       |
| mhz        | number | CPU frequency            |
| cacheSize  | number | Cache size               |
| flags      | Array  | Flags support by the CPU |
| microcode  | string | Microcode info           |
