# Callibrating

- Turn on the compass, with the arrow facing north
- Execute the command `go run . CAL_START` (case-sensetive). You should see a green light on the compass
- Slowly rotate the compass two full rotations, ~1 min/rotation
- When finished, execute the command `go run . CAL_END`. The green light on the compass should turn off

# Getting Compass Angle

## Single Angle

Execute the command `go run . GET_ANGLE`

## Continuous Angle Stream

Execute the command `go run . GET_ANGLE_CONTINUED`