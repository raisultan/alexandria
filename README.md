# Alexandria IDL + UI

Simple, yet powerful IDL and UI tool for its visualization. Mostly oriented for projects using [Django Channels](https://github.com/django/channels). However, most cases might fit for the usage of Alexandria IDL.

## Alexandria IDL
Interface Definition Language for defining and documenting WebSocket interfaces. It enforces usage of an event-based approach and the existence of a maximum of two response models: to a group of users and to a specific user.

### Basic struct
```yaml
info:
    title: string
    description: string
    version: string

server:
    url: string
    description: string

actions:
    action:
        description: string
        actor: # type of entity by which action is triggered [system, client]
            type: object
            properties:
                is_system: bool
                description: string
        data:
            # data schema sent by client-side must be defined here
        response_to_group:
            event: string
            # your response_to_group schema must be defined here
        response_to_user:
            event: string
            # your response_to_user schema must be defined here
```

> For any action fields: `data` is optional. Additionally, one of the response schemas may not be declared. All other fields of the basic struct given above are mandatory.

### Example
```yaml
info:
  title: WS Documentation
  description: WS Documentation description
  version: '0.1'

server:
    url: ws://example.example:1337
    description: Stage server backend

actions:
  new_message:
    description: Send new message
    actor:
      type: object
      properties:
        is_system: false
        description: Can be sent by user
    data:
      type: object
      properties:
        text: string
    response_to_group:
      event: new_message_sent
      message:
        type: object
        properties:
          id:
            type: integer
          text:
            type: string
          user_id:
            type: string
```

A full working example with the use of Alexandria UI can be found [here](https://github.com/raisultan/alexandria/example).

## Alexandria UI
The simplest swagger-like UI for visualization of documentation implemented according to IDL described above. UI is ready to use and its image is hosted on [Docker Hub](https://hub.docker.com/r/raisultan/alexandria).

### Usage

#### Requirements
- completely valid `YAML` file written according to Alexandria IDL
- Alexandria docker image requires `ALEXANDRIA_YAML` environment variable that describes the path to your WebSocket interface documenting file

#### Usage in `docker-compose`
```yaml
version: '3'
services:
  alexandria:
    image: raisultan/alexandria:0.1
    volumes:
      - ./ws.yaml:/ws.yaml
    ports:
      - 5050:8080
    environment:
      ALEXANDRIA_YAML: /ws.yaml
```

### Enhancement, contribution, and feedback
Any feedback or contribution to the project is eagerly welcomed. Just create an issue or contact me on ki.xbozz@gmail.com.

### Todo:
- [ ] add syntax highlighting
- [ ] documentation file structure and content validation
