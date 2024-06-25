## Overview

This plugin for Traefik generates a unique `X-Request-ID` header for each incoming HTTP request.

## Configuring Traefik to Use the Plugin

To configure Traefik to use the plugin, follow these steps:

1. **Create a Traefik Configuration File**:
    - Create a Traefik configuration file, for example, `traefik.yml`.

2. **Add the Plugin Configuration**:
    - Add your plugin to the configuration file. Here is an example configuration:

### Static configuration
```yaml
entryPoints:
   web:
      address: :80
      http:
         middlewares:
            - HeaderRequestId@file

experimental:
   localPlugins:
      HeaderRequestId:
         moduleName: "github.com/bublicov/traefik-x-request-id"
```

### Dynamic configuration
```yaml
http:
  middlewares:
     HeaderRequestId:
        plugin:
           HeaderRequestId: { }
```
