# qrop-coordinator

On-farm coordination server for Qrop autonomous farming robots.

Runs on small ARM SBCs (Raspberry Pi 5, Orange Pi). Manages robot fleet, dispatches missions, and operates fully offline.

## Architecture

- **MQTT broker** — Robot communication (Mosquitto)
- **Robot registry** — Discover, authenticate, monitor connected robots
- **Mission planner** — Break tasks into robot-specific waypoint sequences
- **Safety layer** — Geofencing, collision avoidance, emergency stop
- **Telemetry** — Aggregate sensor data, sync to cloud dashboard

## Development

```bash
go build -o qrop-coordinator .
./qrop-coordinator
```

## Cross-compile for ARM

```bash
GOOS=linux GOARCH=arm64 go build -o qrop-coordinator-arm64 .
```
