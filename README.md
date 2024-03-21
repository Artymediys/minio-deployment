# MinIO deployment
**Ansible-playbook for running MinIO deployment in Multi-node Multi-drive mode.**


## Settings

### Discs
Before launching the playbook you can edit the default value
variable `minio_mounts` to indicate the actual disks on the servers:

```yaml
vars:
   minio_mounts:
     - '/dev/sdb'
     - '/dev/sdc'
     - '/dev/sdd'
     - '/dev/sde'
```


### Root user and password
A mandatory requirement is to specify the Access key and Secret Key values for MinIO.

Example:
```yaml
vars:
   minio_root_user: "minio-login"
   minio_root_password: "minio-super-secret-password"
```

`minio_root_user` must be at least 3 characters long;
`minio_root_password` must be at least 8 characters long;


### Prefix and DNS-zone
You must also specify the prefix and DNS-Zone.

Example:
```yaml
vars:
   minio_servername_prefix: "minio"
   minio_servername_dns_zone: "company.ru"
```

As a result, the address will look like this: `minio.company.ru`


### Erasure Code setting
To configure fault tolerance, the Erasure Code Parity variable must be specified.
You can calculate this variable in the official calculator: `https://min.io/product/erasure-code-calculator`

Example:
```yaml
minio_standard_ec_parity_drives: "2"
```
