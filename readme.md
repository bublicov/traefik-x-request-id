## Overview

This plugin for Traefik generates a unique `X-Request-ID` header for each incoming HTTP request.

## Configuring Traefik to Use the Plugin

To configure Traefik to use the plugin, follow these steps:

1. **Create a Traefik Configuration File**:
    - Create a Traefik configuration file, for example, `traefik.yml`.

2. **Add the Plugin Configuration**:
    - Add your plugin to the configuration file. Here is an example configuration:

3. **Replace params traefikGeoIP**:
    - Replace /usr/local/share/GeoIP/GeoLite2-Country.mmdb with the actual path to your GeoLite2 Country database file
    - adjust the redirectMap as needed.

### Static configuration
```yaml
experimental:
  plugins:
    requestid:
      modulename: "github.com/bublicov/traefik-request-id-plugin"
      version: "v0.1.0"
```

### Dynamic configuration
```yaml
http:
  middlewares:
    requestid:
      plugin:
        requestid: {}

  routers:
    myrouter:
      rule: "Host(`example.com`)"
      middlewares:
        - "requestid"
      service: "myservice"

  services:
    myservice:
      loadBalancer:
        servers:
          - url: "http://127.0.0.1:8080"
```
