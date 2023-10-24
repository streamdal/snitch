Streamdal
=========
[![GitHub](https://img.shields.io/github/license/streamdal/streamdal)](https://github.com/streamdal/streamdal)
[![Discord](https://img.shields.io/discord/123456789?color=blue&label=discord)](https://discord.gg/123456789)

This is the _main_ repo for the [Streamdal](https://streamdal.com) project.

Streamdal is a combination of open source services and SDKs that enable _real-time_:

* Data observability (think `tail -f` for your data!)
* Data governance & enforcement
* Data quality monitoring
* Data transformation + obfuscation + masking + stripping

.. and [a bunch of other things](https://docs.streamdal.com/capabilities).

_This is what it looks like:_
<img src="assets/img.png">

---

<div align="center">

[Overview](#streamdal) •
[Demo](#demo) •
[Getting Started](#getting-started) •
[How Does It Work?](#how-does-it-work) •
[Roadmap](#roadmap) •
[Community](#community)
[Getting Help](#getting-help) •

</div>

---

## Demo

You don't have to install the [server](https://github.com/streamdal/server), 
the [console](https://github.com/streamdal/console)
or [instrument](https://docs.streamdal.com/instrument) any of your apps to see 
Streamdal in action. We've got a live demo :)

**https://demo.streamdal.com**

<sub>The demo has a bunch of sample apps instrumented with the 
[Go SDK](https://github.com/streamdal/go-sdk), doing reads and writes to and 
from various data sources like Kafka, PostgreSQL and HTTP APIs.</sub>

## Getting Started

---

The _easiest_ way to get Streamdal running is to bring up all of the components
on your local machine using [Docker Compose](https://docs.docker.com/compose/).

### Docker Compose

1. Make sure you have [Docker](https://docker.com) installed and running
2. Make sure you have [Docker Compose](https://docs.docker.com/compose/) installed
3. Clone this repo
4. `cd streamdal/install/docker`
5. `docker-compose up -d`
6. `docker ps` -- you should see something like:
```bash
❯ docker ps
CONTAINER ID   IMAGE                             COMMAND                  CREATED              STATUS              PORTS                                            NAMES
32eceb8e4bba   streamdal/snitch-console:latest   "/tini -- docker-ent…"   About a minute ago   Up About a minute   0.0.0.0:3000->3000/tcp                           snitch-console-container
034ec118dcfa   streamdal/snitch-server:latest    "/snitch-server --de…"   About a minute ago   Up About a minute   0.0.0.0:8080->8080/tcp, 0.0.0.0:9090->9090/tcp   snitch-server
8cb6646eb128   redis:latest                      "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:6379->6379/tcp                           redis
ebc04b0e4eee   local-envoy                       "/usr/local/bin/envo…"   3 weeks ago          Up About a minute   0.0.0.0:9091->9091/tcp, 10000/tcp                envoy
```

Open `http://localhost:3000` in your browser and you should see the console.

Ready for [Next Steps](#next-steps)?

### Deploying to Other Platforms

You can install Streamdal on other platforms like Kubernetes, AWS, GCP, Azure,
on-prem, IoT and so on. You can find deployment instructions in the 
[install](./install) dir.

And if something does not work or you get stuck, check out the official [docs](https://docs.streamdal.com).

## How Does It Work?

> [!NOTE]
> For a more in-depth explanation of the architecture, the flow and the various
> components, visit our [docs](https://docs.streamdal.com).

Streamdal consists of **three** main components:

1. [Server](https://github.com/streamdal/server)
    - A distributed server that can be deployed anywhere (cloud, on-prem, IoT, 
      etc.) and it is responsible for facilitating communication between the 
      SDKs and the Console (UI).
2. [Console](https://github.com/streamdal/console)
    - A _beautiful_, figma-like UI that is the primary way you interface with
      your SDKs, define rules, observe data and on.
3. SDKs
    1. [Go](https://github.com/streamdal/sdk-go)
    2. [Node.js](https://github.com/streamdal/sdk-node) 
    3. [Python](https://github.com/streamdal/python)

The basic flow is that you [install](#install) the server and console and wrap
any reads or writes in your app with our SDK. Once that's done, you will be able
to see the data your app is reading or writing in the 
[console](https://github.com/streamdal/console) (or use our 
[CLI](https://github.com/streamdal/cli)).

You will also be able to enforce rules on your data (such as "this should be 
valid JSON", "message should contain a field called `foo`", "strip all email
addresses" and so on).

> [!NOTE]
> [Learn more](https://docs.streamdal.com/arch) about the architecture, the flow and the various components.

## Community

Need help? Have questions? Want to chat? 

Join our [Discord](https://discord.gg/123456789)!

## Contributing

We :heart: contributions! But... before you craft an amazing PR, please read
through our [contributing guidelines](https://docs.streamdal.com/contributing).

## License

This project is licensed under the `Apache-2.0` license. 

See the [LICENSE](LICENSE) file for more info.
