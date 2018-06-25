## Getting Started
A lot of these endpoints are pretty straightforward and mechanical to wrap. It's also super easy if you sign up for a Auckland Transport Developers account to test what kind of things the endpoints return. You can do that [here](https://dev-portal.at.govt.nz/).

## Submitting a PR
### Wrapping an endpoint
Check the list of endpoints that still need to be wrapped in the README, pick one, and wrap it. 

The general pattern for wrapping an endpoint is:
- Find the corresponding endpoint in the Auckland Transport developer's portal, a link to which can also be found in the README
- Hit it, to see what's returned
- Model that response with structs
- Write a method for the `Client` type to hit that endpoint, marshal the response, and return the relevant field
  - See any of the methods for examples, the pattern is pretty similar
  
### Testing
There needs to be a test for every `Client` method, just to double check there were no issues marshalling, it also ensures that no breaking changes slip through. Tests will be run by Travis CI, and a passing build is required before your PR can be merged. 

- Copy the JSON response into a file in `test/payloads` for use in mocking
- Mock the response to that endpoint for use in tests in `mock_client.go`
- Write a test which calls the method you wrote for the `Client` type
  - I'm using [testify](https://github.com/stretchr/testify) for simple assertions
  - In your test, assert that the returned object from your method is _exactly_ equal to what you see in the mocked JSON payload
- Check off the endpoint you've wrapped in the README

### Linting
I'm using [gometalinter](https://github.com/alecthomas/gometalinter) to lint Go files. It's essentially a linter that just runs a collection of linters.

The linter will also run in CI, so your contribution will need to conform to it too.
