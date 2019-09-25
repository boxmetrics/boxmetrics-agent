<a href="https://boxmetrics.github.io/"><img src="https://raw.githubusercontent.com/boxmetrics/boxmetrics-agent/master/.github/boxmetrics-logo.png" width="250" alt="boxmetrics"></a>

# boxmetrics-agent

> This repo contains the boxmetrics agent built with Go

## 📦 Requirements

> Only needed when contributing or building from source

This project should be working as expected with the following minimal version of:

| Dependency | Version |
| ---------- | ------- |
| Go         | >= 1.12 |

## 🚀 Quick start

> The application must be launch with root or a sudoers, otherwise some features can be unavailable

### _From docker image :_

1. **Run this command**

```bash
docker run --rm -p 4455:4455 -p 5544:5544 286829485023.dkr.ecr.us-west-2.amazonaws.com/boxmetrics-agent:<TAG>
```

### _From prebuilt binaries :_

1. **Donwload a binary from Github [release page](https://github.com/boxmetrics/boxmetrics-agent/releases)**

2. **Run application**

```bash
# Made application executable
sudo chmod +x boxmetrics-agent

# Start application
./boxmetrics-agent
```

### _From source code :_

1. **Clone the git repository**

```bash
# cloning git repository
git clone https://github.com/boxmetrics/boxmetrics-agent
```

2. **Build application**

```bash
# go to boxmetrics-agent directory
cd boxmetrics-agent
# run helper command to build
make build
```

3. **Run application**

```bash
# Made application executable
sudo chmod +x bin/boxmetrics-agent

# Start application
./bin/boxmetrics-agent start
```

## 💡 Usage

### Routes

| Path     | Description    |
| -------- | -------------- |
| `/ws/v1` | Websocket root |
| `/`      | Test page      |

### Communication

Both request and response are JSON message

#### Request

| Key     | Type    | Require | Default                | Description      |
| ------- | ------- | ------- | ---------------------- | ---------------- |
| type    | string  | yes     | NA                     | Request type     |
| value   | string  | yes     | NA                     | Type value       |
| options | object  | no      | Default Options Object | Request options  |
| format  | boolean | no      | true                   | Enable formating |

##### Type Values

| Value   | Description                                                    |
| ------- | -------------------------------------------------------------- |
| info    | Return `value` information type                                |
| script  | Run `value` script                                             |
| command | Execute `value` as command _(Use `options` to add parameters)_ |

##### Info Type Values

| Value        | Response                                       | Description                                                   |
| ------------ | ---------------------------------------------- | ------------------------------------------------------------- |
| memory       | [Schema](docs/schema/memory.md#memory)         | Return memory information                                     |
| cpu          | [Schema](docs/schema/cpu.md#cpu)               | Return cpu usage information                                  |
| cpuinfo      | [Schema](docs/schema/cpu.md#cpu-hardware-info) | Return cpu hardware information                               |
| disks        | [Schema](docs/schema/disk.md#disk)             | Return disks information                                      |
| containers   | [Schema](docs/schema/container.md#container)   | Return containers full information                            |
| containersid | Array of string                                | Return containers ID list                                     |
| host         | [Schema](docs/schema/host.md#host)             | Return host information                                       |
| users        | [Schema](docs/schema/host.md#user)             | Return users list                                             |
| sessions     | [Schema](docs/schema/host.md#session)          | Return user sessions list                                     |
| network      | [Schema](docs/schema/network.md#network)       | Return network information                                    |
| connections  | [Schema](docs/schema/network.md#connection)    | Return opened connections list                                |
| processes    | [Schema](docs/schema/process.md#process-light) | Return processes information list                             |
| process      | [Schema](docs/schema/process.md#process)       | Return process full information _(`options.pid` must be set)_ |
| general      | [Schema](docs/schema/general.md)               | Return system wide informations                               |

##### Script Type Values

| Value       | Response | Description                                                                    |
| ----------- | -------- | ------------------------------------------------------------------------------ |
| adduser     | string   | Add user to the system _(`options.args` must be set with corresponding value)_ |
| killprocess | string   | Kill one process _(`options.pid` must be set)_                                 |

##### Options Object

| Key  | Type   | Require | Default | Description                                                                                                                                                                        |
| ---- | ------ | ------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| args | Array  | no      | null    | Array of arguments to pass to the command or to pass to `adduser` script _eg. `["-p <pass>", "-u <username>", "-g group1"]` (`-g` arg is optional and can be used multiple times)_ |
| env  | Array  | no      | null    | Array of environment variable to add before executing command _eg. MY_VAR=abc_                                                                                                     |
| pwd  | string | no      | ""      | Location where the command run, if empty string run in the cwd of the process                                                                                                      |
| pid  | number | no      | 0       | PID use to retrieve information with `process` info type or to kill on script `killprocess`                                                                                        |

#### Response

| Key       | Type   | Description                                                                       |
| --------- | ------ | --------------------------------------------------------------------------------- |
| event     | object | The event send                                                                    |
| data      | object | The data reponse of the event. Corresponding to a specific [schema](docs/schema/) |
| startDate | string | Start date of the response processing                                             |
| endDate   | string | End date of the response processing                                               |
| duration  | string | Duration of the response processing                                               |
| status    | object | Status of the response                                                            |
| error     | string | Error message _(`null` if no error)_                                              |

##### Status object

| Key     | Type   | Description    |
| ------- | ------ | -------------- |
| code    | number | Status code    |
| message | string | Status message |

## 💬 Contributing

1. **Fork the git repository**

2. **Create your feature branch**

3. **Apply your changes**

4. **Run application**

```bash
# run test
make test
# start application in dev mode
make run #  only on first time
go run main.go
```

5. **Open browser to test your change!**

> Project is running at <http://localhost:4455> or <https://localhost:5544>

6. **Commit your changes**

7. **Push it on your fork**

8. **Create new pull request**

## 🧐 What's inside ?

```text
.
├── assets          # Project assets (images, logos, etc)
├── bin             # Project binaries
├── certificates    # Project Certificates
├── cmd             # Main applications for this project
├── configs         # Configuration file templates or default configs
├── docs            # Design and user documents
├── init            # System init and process manager/supervisor configs
├── internal
│   ├── app         # Private application
│   └── pkg         # Private library code
├── scripts         # Scripts to perform various build, install, analysis, etc operations
├── test            # Additional external test apps and test data
├── web             # Web application specific components
├── Dockerfile      # Docker image
├── go.mod          # Module dependencies
├── go.sum          # Ensure dependencies integrity
├── JenkinsFile     # Jenkins pipeline
├── main.go         # Application entry point
├── Makefile        # Helpers command
├── LICENSE
└── README.md
```

## 👥 Contributors

<table width="100%">
  <tbody width="100%">
    <tr width="100%">
      <td align="center" width="33.3333%" valign="top">
        <img style="border-radius: 50%;" width="100" height="100" src="https://github.com/Laurent-PANEK.png?s=100">
        <br>
        <a href="https://github.com/Laurent-PANEK">Laurent Panek</a>
        <p>Security System Integrator</p>
      </td>
     <td align="center" width="33.3333%" valign="top">
        <img style="border-radius: 50%;"  width="100" height="100" src="https://github.com/maxencecolmant.png?s=100">
        <br>
        <a href="https://github.com/maxencecolmant">Maxence Colmant</a>
        <p>DevOps System Integrator</p>
    </td>
          <td align="center" width="33.3333%" valign="top">
        <img style="border-radius: 50%;"  width="100" height="100" src="https://github.com/abdessalamb98.png?s=100">
        <br>
        <a href="https://github.com/abdessalamb98">Abdessalam Benharira</a>
        <p>JavaScript Developer</p>
      </td>
     </tr>
  </tbody>
</table>
