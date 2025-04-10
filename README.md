
# OpenStack Web API Client

A lightweight and modular web API client for managing OpenStack resources via RESTful endpoints using Go and the Gophercloud SDK.

---

## 🧭 Overview

This project is a web client for interacting with OpenStack services. It provides REST APIs built on top of the [Gophercloud](https://github.com/gophercloud/gophercloud) SDK, enabling developers or frontend applications to manage infrastructure programmatically.

---

## ✨ Features

- ✅ RESTful endpoints to interface with OpenStack services  
- ⚙️ Written in idiomatic Go  
- 🔐 Environment-based configuration with `.env` support  
- 🧩 Extensible and modular project structure  
- 📦 Vendor-less dependency management with Go Modules  

---

## 📦 Prerequisites

- Go 1.16 or higher  
- Access to a running OpenStack environment  
- Environment variables or `.env` file for credentials  

---

## 🔧 Installation & Setup

```bash
git clone https://github.com/itera-io/openstack-web-client.git
cd openstack-web-client
go mod tidy
```

### 🛠️ Configuration

Create a `.env` file in the project root with the following content:

```env
OS_AUTH_URL=https://<your-openstack-auth-url>/v3
OS_USERNAME=<your-username>
OS_PASSWORD=<your-password>
OS_PROJECT_NAME=<your-project>
OS_DOMAIN_NAME=<your-domain>
```

Make sure these credentials match your OpenStack environment.

---

## 🚀 Running the Application

```bash
go run main.go
```

By default, the server runs at:

```
http://localhost:8080
```

You can update the listening port or add additional configuration via environment variables or Go flags if implemented.

---

## 📚 API Endpoints

| Method | Endpoint           | Description                  |
|--------|--------------------|------------------------------|
| GET    | `/instances`       | List all compute instances   |
| POST   | `/instances`       | Create a new instance        |
| GET    | `/instances/{id}`  | Get details for one instance |
| DELETE | `/instances/{id}`  | Delete an instance           |

> Note: Endpoint availability depends on the features implemented in the current version.

---

## 🧪 Running Tests

```bash
go test ./...
```

Ensure your OpenStack environment or mock is accessible if tests rely on live resources.

---

## 🤝 Contributing

We welcome contributions! To get started:

1. Fork the repository  
2. Create a feature branch (`git checkout -b feature/my-feature`)  
3. Commit your changes (`git commit -am 'Add my feature'`)  
4. Push to your fork (`git push origin feature/my-feature`)  
5. Open a Pull Request  

Before submitting, make sure to:

- Follow idiomatic Go conventions  
- Add or update relevant tests  
- Keep PRs scoped and descriptive  

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🙌 Acknowledgements

- [Gophercloud](https://github.com/gophercloud/gophercloud) — official Go SDK for OpenStack  
- [OpenStack](https://www.openstack.org/) community for the ecosystem and APIs  

---

Made with ❤️ by [itera-io](https://github.com/itera-io)
