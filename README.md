# CRT.sh Certificate Query Tool

This Go program queries the crt.sh online certificate database for SSL/TLS certificates associated with a specific domain. It outputs the certificates' details in JSON format.

## Features

- Query SSL/TLS certificates by domain.
- Outputs detailed information about each certificate in JSON format.
- Easy to use command-line interface.

## Requirements

- Go 1.15 or newer

## Installation

Clone this repository to your local machine using:

```bash
git clone https://github.com/storbeck/crtsh-query.git
cd crtsh-query
```

## Usage

To run the program, use the following command in the terminal:

```bash
go run main.go -domain=insecure.com
```

Replace `insecure.com` with the domain you wish to query. The program will output the certificates' details in JSON format.

## Example Output

```json
[
  {
    "issuer_ca_id": 16418,
    "issuer_name": "C=US, O=Let's Encrypt, CN=Let's Encrypt Authority X3",
    "common_name": "insecure.com",
    ...
  },
  ...
]
```

## Contributing

Contributions are welcome. Please open an issue first to discuss what you would like to change.

## License

MIT