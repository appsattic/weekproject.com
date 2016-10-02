#!/bin/bash
## --------------------------------------------------------------------------------------------------------------------
#
# Run this file from the main repo directory, not the `install` directory. Thanks. :)
#
## --------------------------------------------------------------------------------------------------------------------

# which user is this
PWD=$(pwd)
USER=$(whoami)
PORT=8504

# make some directories (if they don't yet exist)
sudo mkdir -p /var/log/$USER
sudo chown $USER.$USER /var/log/$USER

# create the supervisord config
cat <<EOF | sudo tee /etc/supervisor/conf.d/weekproject.com.conf
[program:weekproject]
directory = $PWD
command = sudo -E -u $USER ./bin/weekproject
user = $USER
autostart = true
autorestart = true
stdout_logfile = /var/log/$USER/weekproject.com-stdout.log
stderr_logfile = /var/log/$USER/weekproject.com-stderr.log
environment = PORT="8504"
EOF

# now restart/reload supervisord
sudo systemctl reload supervisor.service 

# add the caddy file
cat <<EOF | sudo tee /etc/caddy/vhosts/com.weekproject.conf
weekproject.com {
  proxy / localhost:8504
  tls andychilton@gmail.com
  log stdout
  errors stderr
}
EOF

# and restart caddy
sudo systemctl restart caddy.service

# sudo journalctl -boot -u caddy.service

## --------------------------------------------------------------------------------------------------------------------
