# Terraform Jenkins provider

Terraform allows you to keep your infrastructure as a code (HCL language).
This provider supports controlling jenkins over its http api.

## Configuration of provider

```hcl
provider "jenkins" {
    url = "http://jenkins_url:port"
    username = "user"
    password = "pass"
}
```

`url` [required] - and is the full address of jenkins.

`username`, `password` [optional] - If jenkins has authentication enabled then `username` and `password` are also required, they wil be used to authenticate into jenkins.

`ca_cert` [optional] - should be provided if jenkins uses self signed certificate.

## Supported resources

* (#credentials credentials)
  * (#jenkins_username_credential)
  * (#jenkins_ssh_credential)
  * (#docker credentials)
* (#Plugins plugin)
* (#Job job)

---

### [credentials](#credentials)

Common fields for all credentials:

`identifier` [required] - id under which given credential will be stored in jenkins.
`username` [required] - username for given credentials.
`scope` [optional] - scope for credential in jenkins ( default to `global` ).
`domain` [optional] - domain for credential in jenkins ( defaults to `_` ).
`description` [optional] - describes what credential is used for.

### [jenkins_username_credential](#jenkins_username_credential)

```hcl
resource "jenkins_username_credential" "admin" {
    identifier = "cred_id"
    username = "admin"
    password = "admin"
}
```

This will create username credential.

`password` [required] - password for given user.

### [jenkins_ssh_credential](#jenkins_ssh_credential)

```hcl
resource "jenkins_ssh_credential" "ssh_user" {
    identifier = "ssh_cred_id"
    username = "user"
    passphrase = "top-secret"
    value_type = "directValue"
    value = "contents_of_ssh_key"
}
```

`passphrase` [optional] - used when provided ssh key is password protected.
`value_type` [required] - can equal either `directValue` or `fileOnMaster`, used to determine contents of `value` field.
`value` [required] - if `value_Type` equals  `directValue` then this field should contain directly pasted ssh key. If this field equals `fileOnMaster` then it should be path to ssh key on jenkins master.

This will create ssh credential, ssh key can be provided as inline value or as path on jenkins master.

### [jenkins_docker_credential](#jenkins_docker_credential)

```hcl
resource "jenkins_docker_credential" "docker_user" {
    identifier = "docker_cred_id"
    username = "docker-user"
    server_ca_certificate = "ca_cert"
    client_certificate = "client_cert"
    client_key = "client_key"
}
```

This will create docker credential.
`server_ca_certificate` [optional] - server CA certificate value.
`client_certificate` [optional] - client certificate value.
`client_key` [optional] - client key value.


## [plugins](#plugins)

```hcl
"jenkins_plugin" "packer" {
    name = "packer"
    version = "1.4"
}
```

Will try to install plugin with given name and version
`name` [required] - id of plugin to install.
`version` [required] - version of plugin to install

## [job](#job)

```hcl
"jenkins_xml_job" "job" {
    name = "test-job"
    xml = "${file("./job.xml")}"
}
```

example of contents of job.xml:

```xml
<?xml version='1.0' encoding='UTF-8'?>
  <project>
    <description>test my </description>
    <keepDependencies>false</keepDependencies>
    <properties/>
      ....
        some other properties
      ....
    <scm class="hudson.scm.NullSCM"/>
    <triggers class="vector"/>
  </project>
```

Will create job on jenkins with given configuration. Unfortunately current version shows that there are always changes in xml file.

`name` [required] - name of a job that should be created.
`xml` [required] - xml configuration of a job.