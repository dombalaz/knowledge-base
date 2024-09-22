# Rebind protection

DNS rebinding is a method of manipulating resolution of domain names that is commonly used as a form of computer attack

By default, OpenWrt protects against DNS rebinding, when using dnsmasq. This is an help message:

> Discard upstream responses containing RFC1918 addresses. Discard also upstream responses containing RFC4193, Link-Local and private IPv4-Mapped RFC4291 IPv6 Addresses.

This is a problem when I set my own domain to resolve to a private ip address. To enable this, I can either disable this protection or enter my domain to a domain whitelist, which fixes the issue.