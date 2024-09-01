## Overview

This plugin for Traefik generates a unique `X-Request-ID` header for each incoming HTTP request. It supports different types of request IDs, including UUID, ULID, and Snowflake, and allows you to configure whether to override existing `X-Request-ID` headers.

## Configuration

### Plugin Configuration

The plugin configuration is defined in the `Config` struct, which includes the following fields:

- **RequestIDType** (optional, default: `uuid`): Specifies the type of request ID to generate.
  Possible values are `uuid`, `ulid`, `snowflake`.
- **RequestIDOverride** (optional, default: `false`): Determines whether to override an existing X-Request-ID header if it already exists in the request.
- **NodeID** (optional, default: `1`): Specifies the node ID to use when generating Snowflake IDs.

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
      traefik-x-request-id:
         moduleName: "github.com/bublicov/traefik-x-request-id"
```

### Dynamic configuration
```yaml
http:
   middlewares:
      HeaderRequestId:
         plugin:
            traefik-x-request-id:
               RequestIDType: "uuid"  # Type of request ID to generate (default: "uuid")
               RequestIDOverride: false  # Whether to override existing X-Request-ID headers (default: false)
               NodeID: 1  # Node ID for Snowflake (default: 1)
```

## Installation

To use the plugin, you need to install it as a **LOCAL PLUGIN** for Traefik.

1. **Clone the Plugin Repository**: Clone the repository of the Traefik IP2Location Plugin to your local path
   {root_traefik_dir}/plugins-local/src/github.com/bublicov/traefik-ip2location

    ```sh
    git clone https://github.com/bublicov/traefik-x-request-id.git
    ```

2. **Static configuration**: Modify your Traefik configuration to include the local plugin. Here is an example of how to
   do
   this in your `traefik.yml` file:

  ```yaml
  entryPoints:
     web:
        address: :80
        http:
           middlewares:
              - HeaderRequestId@file
  
  experimental:
     localPlugins:
        traefik-x-request-id:
           moduleName: "github.com/bublicov/traefik-x-request-id"
  ```

3. **Dynamic Configuration**: Create a `dynamic.yml` file to define the middleware configuration for the plugin.

  ```yaml
  http:
    middlewares:
      HeaderRequestId:
        plugin:
          traefik-x-request-id: { }
  ```

### License

This plugin is licensed under the MIT License. See the LICENSE file for more details.