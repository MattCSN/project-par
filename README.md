# Project Par Golf Management System

This project is a comprehensive golf management system designed to facilitate the management of golf courses, including tracking of golf courses, holes, and tees. Built with Go, it leverages the Gin framework for RESTful API services and GORM for ORM functionalities.

## Getting Started

These instructions will guide you through setting up the project on your local machine for development and testing purposes.

### Prerequisites

Ensure you have Go installed on your machine. This project uses Go Modules for dependency management, so Go version 1.11 or higher is required.

```bash
go version
```

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/MattCSN/project-par.git
cd project-par
```

Install the necessary dependencies:

```bash
go mod tidy
```

### Running the Application

To start the server, run:

```bash
go run main.go
```

## Usage

The application provides endpoints for managing golf courses, holes, and tees. Here are some examples of how to interact with the API:

- **Get all golf courses**: `GET /api/golfs`
- **Create a new golf course**: `POST /api/golfs`
- **Add multiple golf courses**: `POST /api/golfs/batch`