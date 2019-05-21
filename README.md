# boxmetrics-agent

> This repo contains the boxmetrics agent built with Go

## 📦 Requirements

> Only if you want to run the source code

This project should be working as expected with the following minimal version of:

| Dependency | Version  |
| ---------- | :------: |
| Go         | >= v1.11 |

The project might be working with older version of Go, if you add vendor directory with `go mod vendor` command _(available from Go 1.11)_. Since this repo is based on go modules dependencies, the vendor directory won't be add, it's your own choice if you want it.

## 🚀 Quick start

1. **Clone the git repository**

   ```bash
   # cloning git repository
   git clone https://github.com/boxmetrics/boxmetrics-agent

   cd boxmetrics-agent
   ```

2. **Run application**

   **Dev Version**

   ```bash
   # start the app
   make run
   ```

   **Prod Version**

   ```bash
   # build the app
   make build
   # run executable
   ./bin/boxmetrics-agent
   ```

3. **Open browser and start editing files!**

> Project is running at <http://localhost:8080> or <https://localhost:9090>

## 🧐 What's inside ?

```text
.
├── assets          # Project assets (images, logos, etc)
├── bin             # Project binaries
├── build
│   ├── package     # Package configurations and scripts (Docker, deb, rpm, pkg)
│   └── ci          # CI configurations and scripts (travis, circle, jenkins)
├── certificates    # Project Certificates
├── cmd             # Main applications for this project
├── configs         # Configuration file templates or default configs
├── deployments     # Deployment configurations and templates (docker-compose, kubernetes/helm)
├── docs            # Design and user documents
├── init            # System init and process manager/supervisor configs
├── internal
│   ├── app         # Private application
│   └── pkg         # Private library code
├── pkg             # Public library code
├── scripts         # Scripts to perform various build, install, analysis, etc operations
├── test            # Additional external test apps and test data
├── web             # Web application specific components
├── go.mod          # Module dependencies
├── go.sum          # Ensure dependencies integrity
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
        <img style="border-radius: 50%;"  width="100" height="100" src="https://github.com/Abdessalam98.png?s=100">
        <br>
        <a href="https://github.com/Abdessalam98">Abdessalam Benharira</a>
        <p>JavaScript Developer</p>
      </td>
     </tr>
  </tbody>
</table>
