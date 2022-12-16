# Service Management (systemd)
systemd is a System Management Daemon which replaces the sysvinit process to become the first process with PID = 1, which gets executed in user space during the Linux start-up process. It is a system that is designed specifically for the Linux kernel. It is now being used as a replacement of init.d to overcome shortcomings of it. It uses systemctl command to perform related operations.

e.g. $ systemctl start [service-name], $ systemctl poweroff