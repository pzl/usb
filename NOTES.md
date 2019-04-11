


USBDEVFS_ syscall |               Argument            |  Data dir
------------------|-----------------------------------|------------
SETINTERFACE      | struct usbdevfs_setinterface      | out
SETCONFIGURATION  | int (config num)                  | out
GETDRIVER         | struct usbdevfs_getdriver         | in
CONNECTINFO       | struct usbdevfs_connectinfo       | in
CLAIMINTERFACE    | unsigned int (interface num)      | out
RELEASEINTERFACE  | unsigned int (interface num)      | out
DISCONNECT_CLAIM  | struct usbdevfs_disconnect_claim  | out
RESETEP           | unsigned int (endpt)              | out
CLEAR_HALT        | unsigned int (endpt)              | out
CLAIM_PORT        | unsigned int (usb hub port num?)  | out
RELEASE_PORT      | unsigned int (usb hub port num)   | out
HUB_PORTINFO*     | struct usbdevfs_hub_portinfo      | in
GET_CAPABILITIES  | __u32                             | in
ALLOC_STREAMS     | struct usbdevfs_streams           | out
FREE_STREAMS      | struct usbdevfs_streams           | out
DROP_PRIVILEGES   | u32                               | out
DISCARDURB        | struct usbdevfs_urb               | out
RESET             | NONE (legit, nothing)             | n/a
DISCONNECT        | // sent through USBDEVFS_IOCTL    | ??
CONNECT           | // sent through USBDEVFS_IOCTL    | ??
GET_SPEED         | NONE                              | n/a
CONTROL           | struct usbdevfs_ctrltransfer      |  ??
BULK              | struct usbdevfs_bulktransfer      | out
SUBMITURB         | struct usbdevfs_urb               | out
REAPURB           | *struct usbdevfs_urb              | in
REAPURBNDELAY     | *struct usbdevfs_urb              | in
DISCSIGNAL        | struct usbdevfs_disconnectsignal  | out
IOCTL^            | struct usbdevfs_ioctl             | ??




Data dir:
    - **out**: passing an argument or value to kernel
    - **in**: passing a struct or pointer to the kernel to have fields filled in

* only sent to USB hub devices

^ talks to kernel driver directly