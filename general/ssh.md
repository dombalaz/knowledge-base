# SSH Stuff

## Generate Key Pair

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

## Copy the key to a server

```bash
ssh-copy-id -i ~/.ssh/ed25519 user@host
```
