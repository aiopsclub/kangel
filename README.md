## ⚡️ Quick start
The kangel is the angel daemon for kubelet, it can make kubelet more stable.

First of all, it require Helm, The chart in the chart dir. 

## 📖 Project Wiki

The best way to better explore all the features of the **Create Go App CLI** is to read the project [Wiki](https://github.com/create-go-app/cli/wiki) and take part in [Discussions](https://github.com/create-go-app/cli/discussions) and/or [Issues](https://github.com/create-go-app/cli/issues). Yes, the most frequently asked questions (_FAQ_) are also [here](https://github.com/create-go-app/cli/wiki/FAQ).

## ⚙️ Commands & Options

### `create`

CLI command for create a new project with the interactive console UI.

```bash
cgapp create [OPTION]
```

| Option | Description                                              | Type   | Default | Required? |
| ------ | -------------------------------------------------------- | ------ | ------- | --------- |
| `-t`   | Enables to define custom backend and frontend templates. | `bool` | `false` | No        |

![cgapp_create](https://user-images.githubusercontent.com/11155743/116796937-38160080-aae9-11eb-8e21-fb1be2750aa4.gif)

- 📺 Full demo video: https://recordit.co/OQAwkZBrjN
- 📖 Docs: https://github.com/create-go-app/cli/wiki/Command-create

### `deploy`

CLI command for deploy Docker containers with your project via Ansible to the remote server.

> 🔔 Make sure that you have [Python 3.8+](https://www.python.org/downloads/) and [Ansible 2.9+](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-ansible-on-specific-operating-systems) installed on your computer.

```bash
cgapp deploy [OPTION]
```

| Option | Description                                                                                            | Type   | Default | Required? |
| ------ | ------------------------------------------------------------------------------------------------------ | ------ | ------- | --------- |
| `-k`   | Prompt you to provide the remote user sudo password (_a standard Ansible `--ask-become-pass` option_). | `bool` | `false` | No        |

![cgapp_deploy](https://user-images.githubusercontent.com/11155743/116796941-3c421e00-aae9-11eb-9575-d72550814d7a.gif)

- 📺 Full demo video: https://recordit.co/ishTf0Au1x
- 📖 Docs: https://github.com/create-go-app/cli/wiki/Command-deploy

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `Create Go App CLI`:

- Add a [GitHub Star](https://github.com/create-go-app/cli) to the project.
- Tweet about project [on your Twitter](https://twitter.com/intent/tweet?text=%E2%9C%A8%20Create%20a%20new%20production-ready%20project%20with%20%23Golang%20backend%2C%20%23JavaScript%20or%20%23TypeScript%20frontend%2C%20%23Docker%20and%20%23Ansible%20deploy%20automation%20by%20running%20one%20command.%20%0A%0AFocus%20on%20writing%20code%20and%20thinking%20of%20business-logic%21%0AThe%20CLI%20will%20take%20care%20of%20the%20rest.%0A%0Ahttps%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
- Write interesting articles about project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or personal blog.
- Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **$100** and we get $25).

Together, we can make this project **better** every day! 😘
