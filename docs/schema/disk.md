# Response Schema for Disk informations

## Disk

| Key        | Type   | Description                       |
| ---------- | ------ | --------------------------------- |
| device     | string | Disk name                         |
| mountpoint | string | Disk mountpoint                   |
| fstype     | string | File system type                  |
| opts       | string | Disk options                      |
| usage      | object | [Disk usage](#disk-usage)         |
| iocounters | object | [Disk io counter](#disk-counters) |

### Disk Usage

| Key               | Type   | Description                   |
| ----------------- | ------ | ----------------------------- |
| path              | string | Disk path                     |
| fstype            | string | Disk file system type         |
| total             | string | Disk total size               |
| free              | string | Disk free space               |
| used              | string | Disk used space               |
| usedPercent       | string | Percentage of disk space used |
| inodesTotal       | string | Total number of inodes        |
| inodesUsed        | string | Number of inodes used         |
| inodesFree        | string | Number of inodes free         |
| inodesUsedPercent | string | Percentage of inodes used     |

### Disk Counters

| Key              | Type   | Description                         |
| ---------------- | ------ | ----------------------------------- |
| readCount        | string | Number of read count                |
| mergedReadCount  | string | Number of merge read count          |
| writeCount       | string | Number of write count               |
| mergedWriteCount | string | Number of merged write count        |
| readBytes        | string | Number of read bytes                |
| writeBytes       | string | Number of write bytes               |
| readTime         | string | Total of time used to read          |
| writeTime        | string | Total of time used to write         |
| iopsInProgress   | string | Number of IO per second in progress |
| ioTime           | string | Total of time used for IO           |
| weightedIO       | string | Number of weighted IO               |
| name             | string | Disk name                           |
| serialnumber     | string | Disk serialnumber                   |
| label            | string | Disk label                          |
