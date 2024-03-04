# MinIO deployment

**Ansible-playbook for running MinIO deployment in Multi-node Multi-drive mode.**


### Settings

Before launching the playbook, you can edit the default value
variable `minio_mounts` to indicate the actual disks on the nodes:

```yaml
vars:
   minio_mounts:
     - '/dev/sdb'
     - '/dev/sdc'
     - '/dev/sdd'
     - '/dev/sde'
```

You should also specify the Access key and Secret Key for MinIO:
```yaml
vars:
   minio_root_user: "minio-login"
   minio_root_password: "minio-super-secret-password"
```

`minio_root_user` must be at least 3 characters long;
`minio_root_password` must be at least 8 characters long;


### Launch

There are two ways to launch a playbook:
1) Using the command `ansible-playbook main.yml --ask-become-pass`
2) Run the binary file `minio-deployment-<arm/amd>64`
