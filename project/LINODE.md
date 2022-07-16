# Deploy to Linode

Linode Cloud Manager
- https://cloud.linode.com/linodes

Provision 2x Ubuntu 20.04 LTS Server ($10 a month, shared CPU)
- node-1
- node-2

## Setup servers

Use the same steps for each server

SSH using credentials provided on Linode dashboard

Add a user
- adduser <name>
Add user to sudo group
- usermod -aG sudo <name>
Firewall
- ufw allow ssh (make sure you do this )
- ufw allow http
- ufw allow https
- ufw allow 2377/tcp
- ufw allow 7946/tcp
- ufw allow 7946/udp
- ufw allow 4789/udp
- ufw allow 8025/tcp
- ufw enable
Update packages:
- apt update
- apt upgrade

Install Docker
- https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository

Set hostname
- sudo hostnamectl set-hostname node-1

Hosts file
- sudo vi /etc/hosts
- add entries for IP address - node-1
- add entries for IP address - node-2

DNS (Google Domains)
- domains.google.com

Add records
- A record, node-1, node-1 IP address
- A record, node-2, node-2 IP address
- A record, swarm, node-1 IP address
- A record, swarm, node-2 IP address (note, same 'swarm' subdomain, we can hit any node in the swarm, actually can't do this on google domains)
- CNAME record, broker, swarm.peacefixation.net

Initialize Docker Swarm

node-1
- docker swarm init --advertise-addr <ip address>

node-2 (note this command is output by the first command)
- docker swarm join --token <token> <ip address>:<port>

Configure Caddy For Production
- cp Caddyfile Caddyfile.production
- change localhost -> node-1
- change backend -> broker.peacefixation.net
- cp caddy.dockerfile caddy.production.dockerfile
- change Caddyfile reference to Caddyfile.production
- build and tag the Caddyfile docker image but change name to gomicro-caddy-production (see Makefile)
- push the image (see Makefile)
- ssh to node-1
- sudo mkdir /swarm
- sudo chown <name>:<name> /swarm
- mkdir /swarm/caddy_data
- mkdir /swarm/caddy_config
- scp swarm.production.yml to node-1:/swarm/