# <img alt="short-url" src="https://github.com/gyaan/short-urls/blob/develop/assets/short_url.png" width="220" />

[![Go Version](https://img.shields.io/badge/Go-1.13+-blue.svg)](https://golang.org/)
[![MongoDB](https://img.shields.io/badge/MongoDB-4.2.1+-green.svg)](https://www.mongodb.com/)
[![React](https://img.shields.io/badge/React-16.12.0+-blue.svg)](https://reactjs.org/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

**Short URLs** - A modern URL shortening service built with Go, MongoDB, and React. Create tiny URLs from long ones and track click analytics.

**Explore the docs »** [API Documentation](docs/README.md) · [Report Bug](https://github.com/gyaan/short-urls/issues) · [Request Feature](https://github.com/gyaan/short-urls/issues)

## Table of Contents

1. [About The Project](#about-the-project)
   * [Built With](#built-with)
2. [Getting Started](#getting-started)
   * [Prerequisites](#prerequisites)
   * [Installation](#installation)
3. [Usage](#usage)
4. [API Documentation](#api-documentation)
5. [Architecture](#architecture)
6. [Roadmap](#roadmap)
7. [Contributing](#contributing)
8. [License](#license)
9. [Contact](#contact)
10. [Acknowledgments](#acknowledgments)

## About The Project

Short URLs is a comprehensive URL shortening service that allows users to create shortened versions of long URLs while providing click tracking and analytics. This project demonstrates modern web development practices with a microservices architecture.

### Key Features

* **URL Shortening**: Convert long URLs into short, memorable links
* **Click Tracking**: Monitor and analyze click statistics
* **User Authentication**: Secure user registration and login
* **RESTful API**: Clean, documented API endpoints
* **Modern Frontend**: React-based user interface
* **Docker Support**: Easy deployment with Docker Compose
* **Analytics Dashboard**: Visual representation of URL performance

### Why Short URLs?

* **Space Efficiency**: Perfect for social media, SMS, and email sharing
* **Branding**: Create memorable, branded short links
* **Analytics**: Track engagement and click-through rates
* **Convenience**: Easy to remember and share

## Built With

This section lists the major frameworks and libraries used to build this project.

### Backend
* [Go](https://golang.org/) - Primary programming language
* [Chi](https://github.com/go-chi/chi) - HTTP router and middleware
* [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver
* [Viper](https://github.com/spf13/viper) - Configuration management
* [JWT-Go](https://github.com/dgrijalva/jwt-go) - JWT authentication

### Frontend
* [React](https://reactjs.org/) - JavaScript library for building user interfaces
* [Material-UI](https://material-ui.com/) - React UI framework
* [Recharts](https://recharts.org/) - Composable charting library

### DevOps
* [Docker](https://www.docker.com/) - Containerization
* [Docker Compose](https://docs.docker.com/compose/) - Multi-container orchestration

## Getting Started

This is an example of how you can set up the project locally. Follow these simple steps to get a local copy up and running.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

#### Option 1: Local Development Setup

* **Go** (>= v1.13)
  ```bash
  # macOS
  brew install go
  
  # Ubuntu/Debian
  sudo apt-get install golang-go
  
  # Windows
  # Download from https://golang.org/dl/
  ```

* **MongoDB** (>= v4.2.1)
  ```bash
  # macOS
  brew install mongodb/brew/mongodb-community
  
  # Ubuntu/Debian
  sudo apt-get install mongodb
  
  # Windows
  # Download from https://www.mongodb.com/try/download/community
  ```

* **Node.js** (>= 14.0.0)
  ```bash
  # macOS
  brew install node
  
  # Ubuntu/Debian
  curl -fsSL https://deb.nodesource.com/setup_14.x | sudo -E bash -
  sudo apt-get install -y nodejs
  
  # Windows
  # Download from https://nodejs.org/
  ```

#### Option 2: Docker Setup

* **Docker** (>= 19.03)
  ```bash
  # macOS
  brew install docker
  
  # Ubuntu/Debian
  sudo apt-get install docker.io docker-compose
  
  # Windows
  # Download from https://www.docker.com/products/docker-desktop
  ```

### Installation

#### Docker Installation (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/gyaan/short-urls.git
   cd short-urls
   ```

2. **Run with Docker Compose**
   ```bash
   docker-compose -f build/package/docker-compose.yml up
   ```

3. **Access the application**
   - Backend API: http://localhost:8080
   - Frontend: http://localhost:3000

#### Local Development Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/gyaan/short-urls.git
   cd short-urls
   ```

2. **Set up MongoDB**
   ```bash
   # Start MongoDB service
   sudo systemctl start mongod
   
   # Or run MongoDB manually
   mongod --dbpath /path/to/data/directory
   ```

3. **Configure the application**
   ```bash
   # Copy and edit the configuration file
   cp config/config.yaml.example config/config.yaml
   # Edit config/config.yaml with your settings
   ```

4. **Run the backend**
   ```bash
   # Install dependencies
   go mod download
   
   # Run the application
   go run cmd/short-urls/main.go
   ```

5. **Run the frontend**
   ```bash
   cd web
   npm install
   npm start
   ```

## Usage

### API Endpoints

#### Authentication
```bash
# Register a new user
POST /api/auth/register
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}

# Login
POST /api/auth/login
{
  "email": "john@example.com",
  "password": "password123"
}
```

#### URL Management
```bash
# Create a short URL
POST /api/short-urls
Authorization: Bearer <token>
{
  "url": "https://example.com/very-long-url"
}

# Get all short URLs
GET /api/short-urls?page=1&limit=10
Authorization: Bearer <token>

# Get specific short URL
GET /api/short-urls/{id}
Authorization: Bearer <token>

# Update short URL status
PUT /api/short-urls/{id}
Authorization: Bearer <token>
{
  "status": 1
}

# Delete short URL
DELETE /api/short-urls/{id}
Authorization: Bearer <token>
```

#### URL Redirection
```bash
# Redirect to original URL
GET /{short-url-identifier}
```

### Frontend Usage

1. **Sign up/Login**: Create an account or sign in
2. **Create Short URLs**: Enter a long URL to get a shortened version
3. **Manage URLs**: View, edit, and delete your short URLs
4. **Analytics**: Monitor click statistics and performance

## API Documentation

For detailed API documentation, see the [API Documentation](docs/README.md) or explore the [Postman Collection](api/postman-collections/short-urls.postman_collection.json).

## Architecture

### Project Structure
```
short-urls/
├── cmd/                    # Application entry points
│   └── short-urls/        # Main application
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── handler/          # HTTP request handlers
│   ├── middleware/       # HTTP middleware
│   ├── models/           # Data models
│   ├── repositories/     # Data access layer
│   └── router/           # Route definitions
├── pkg/                  # Public libraries
│   ├── pagination/       # Pagination utilities
│   ├── url/             # URL validation
│   └── url_shortner/    # URL shortening algorithm
├── web/                  # React frontend
├── config/               # Configuration files
├── docs/                 # Documentation
└── api/                  # API documentation
```

### Technology Stack
- **Backend**: Go with Chi router and MongoDB
- **Frontend**: React with Material-UI
- **Authentication**: JWT tokens
- **Database**: MongoDB with official Go driver
- **Deployment**: Docker and Docker Compose

## Roadmap

* [ ] Add URL expiration functionality
* [ ] Implement click analytics dashboard
* [ ] Add bulk URL import/export
* [ ] Create mobile application
* [ ] Add URL customization (custom short codes)
* [ ] Implement rate limiting
* [ ] Add API rate limiting
* [ ] Create admin dashboard
* [ ] Add URL preview functionality
* [ ] Implement URL password protection

See the [open issues](https://github.com/gyaan/short-urls/issues) for a full list of proposed features and known issues.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

1. **Fork the Project**
2. **Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3. **Commit your Changes** (`git commit -m 'Add some AmazingFeature'`)
4. **Push to the Branch** (`git push origin feature/AmazingFeature`)
5. **Open a Pull Request**

### Development Guidelines

* Follow Go best practices and conventions
* Write tests for new functionality
* Update documentation for API changes
* Use conventional commit messages
* Ensure all tests pass before submitting PR

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Gyaan - [@gyaan](https://github.com/gyaan) - gyaan@example.com

Project Link: [https://github.com/gyaan/short-urls](https://github.com/gyaan/short-urls)

## Acknowledgments

* [Best README Template](https://github.com/othneildrew/Best-README-Template) - For this amazing README template
* [Go Chi Router](https://github.com/go-chi/chi) - Lightweight HTTP router
* [Material-UI](https://material-ui.com/) - React UI framework
* [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver
* [Viper](https://github.com/spf13/viper) - Configuration solution for Go applications
* [JWT-Go](https://github.com/dgrijalva/jwt-go) - JWT implementation in Go

---

<div align="center">

**Made with ❤️ by [Gyaan](https://github.com/gyaan)**

[Back to top](#)

</div>