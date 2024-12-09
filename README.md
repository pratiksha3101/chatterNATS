
---

# **ChatterNATS: A Fun CLI Chat Application Powered by NATS JetStream 🎉**

**ChatterNATS** is a lightweight, real-time chat application built with Go and NATS JetStream. Whether you're looking to
dive into messaging systems, explore distributed architecture, or just enjoy coding something interactive, **ChatterNATS
** is the perfect fun and engaging project!

---

## 🚀 **Features**

- 🌟 **Real-Time Messaging**: Chat instantly with multiple participants.
- 🔗 **UUID-Based Message Tracking**: Ensure every message is uniquely identified.
- 💾 **JetStream for Persistence**: Never lose your messages.
- 🖥️ **Minimalistic CLI Interface**: Chatting has never been so simple!
- 🎉 **A Fun Learning Experience**: Dive into the world of messaging systems while having fun.

---

## 🛠️ **Prerequisites**

Before you start, make sure you have:

- [Go](https://go.dev/dl/) (v1.22 or higher)
- [NATS Server](https://docs.nats.io/running-a-nats-service/introduction/installation)

---

## 📦 **Installation**

### 1️⃣ Clone the Repository

```bash
git clone git@github.com:pratiksha3101/chatterNATS.git
cd chatterNATS
```

### 2️⃣ Download Dependencies

```bash
go mod download
```

### 3️⃣ Start the NATS Server

Install NATS Server using Homebrew:

```bash
brew install nats-server
```  

**Create a Configuration File:**  
Refer to
the [NATS Configuration Documentation](https://docs.nats.io/running-a-nats-service/configuration/resource_management)
for guidance on creating a config file.

After creating the configuration file (e.g., `nats-server.conf`), launch the server:

```bash
nats-server -c nats-server.conf
```

### 4️⃣ Run the Chat Server

```bash
go run cmd/server/main.go
```

### 5️⃣ Start Chatting with a Client

Open a new terminal and run:

```bash
go run cmd/client/main.go
```

---

## 🧑‍💻 **Usage**

- **Set Your Username:** Enter your chat identity at the start.
- **Join a Channel:** Type your message and press Enter to send.
- **Leave the Chat:** Type `/quit` when you're done.

---

## 📂 **Project Structure**

- `cmd/`: Main entry points for the chat server and clients.
- `internal/`: Core logic for chat functionalities.
- `pkg/`: Shared utilities, including NATS helpers.

---

## 🏗️ **Architecture**

ChatterNATS leverages **NATS JetStream** for real-time delivery and message persistence. Its modular and easy-to-follow
codebase makes it simple to extend and adapt.

---

## 📚 **Official Resources**

- [NATS Documentation](https://docs.nats.io/)
- [JetStream Overview](https://docs.nats.io/nats-concepts/jetstream)
- [Go Programming Language](https://go.dev/)

---

## 🎉 **Why ChatterNATS?**

Whether you're a beginner or an experienced developer, **ChatterNATS** is a great way to explore distributed messaging
systems in a hands-on, fun way. Unleash your creativity, build chatrooms, and enjoy coding with NATS JetStream and Go.

Get started today and let the chatting begin! 🚀

---
