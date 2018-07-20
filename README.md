# at-go [![Build Status](https://travis-ci.org/Mungrel/at-go.svg?branch=master)](https://travis-ci.org/Mungrel/at-go)
A Go Wrapper for the [Auckland Transport API](https://dev-portal.at.govt.nz/docs/services/)

## Contributing
Always keen to take contributions. Check out [Contributing](https://github.com/Mungrel/at-go/blob/master/CONTRIBUTING.md) for guidelines and tips on contributing.

## Endpoints

[General Transit Feed (GTFS)](https://dev-portal.at.govt.nz/docs/services/gtfs/operations/580698de7d6df41584d3d0ce)
- [x] Agencies
- [x] Calendar by Service
- [x] Calendar Dates
- [x] Calendar Dates by Service
- [x] Calendars
- [x] Metadata
- [x] Routes
- [x] Routes by Location
- [x] Routes by Long Name
- [x] Routes by Short Name
- [x] Routes by Stop
- [x] Routes Id
- [x] Routes Search
- [x] Shape geometry by Id
- [x] Shapes by Id
- [x] Shapes by Trip
- [x] Stop by Code
- [x] Stop by Id
- [ ] Stop by Trip
- [x] Stop Info By Code
- [x] Stop Times by Id
- [ ] Stop Times by Trip
- [ ] Stop Times by Trip and Sequence
- [x] Stops
- [x] Stops by Location
- [x] Stops Search
- [x] Trips
- [x] Trips by Id
- [x] Trips by Route
- [x] Versions

[Locations](https://dev-portal.at.govt.nz/docs/services/locations/operations/5806972ae7e2890e84718eef)
- [x] Customer service centres
- [x] Parking Locations
- [ ] Points of Interest
- [x] Scheduled Works

[Notifications](https://dev-portal.at.govt.nz/docs/services/notifications/operations/57e490217d6df40f107a663e)

- [x] Notifications
- [x] Notifications by Category
- [x] Notifications by Stop

[Realtime Transit Feed (GTFS)](https://dev-portal.at.govt.nz/docs/services/realtime/operations/580698237d6df41584d3d0c9)

- [ ] Combined Feed
- [ ] Service Alerts
- [ ] Trip Updates
- [x] Vehicle Positions
