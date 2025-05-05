# Simple NPS Backend

## Description

It's a simple backend project designed to run on a cloud function

![Architecture Diagram](docs/arch.png)

## Features

- Stores a submited survey

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Jardessomonster/nps-back
```

2. Navigate to the project directory:

```bash
cd nps-back
```

3. Install dependencies:

```bash
go mod tidy
```

## Usage

Setup environment values

```bash
MONGODB_URL=your@mongodb/db
```

Run the application:

```bash
# Build and execute the binary
go build -o nps-back
./nps-back

# Or use the make command
make dev
```
