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

you can omit at the hcl file the user and pass for security reason. To manage this, export the variables like the following commands:
```
export JENKINS_USERNAME=user
export JENKINS_PASSWORD=pass
```
```hcl
provider "jenkins" {
    url = "http://jenkins_url:port"
}
```

`url` [required] - and is the full address of jenkins.

`username`, `password` [optional] - If jenkins has authentication enabled then `username` and `password` are also required, they wil be used to authenticate into jenkins. If you prefer, you can use the export the environment variables (`JENKINS_USERNAME, JENKINS_PASSWORD`).

`ca_cert` [optional] - should be provided if jenkins uses self signed certificate.

## Supported resources

* [credentials](#credentials)
  * [username_credential](#jenkins_username_credential)
  * [ssh_credential](#jenkins_ssh_credential)
  * [credentials](#jenkins_docker_credential)
  * [secret_text](#jenkins_secret_credential)
* [plugin](#plugin)
* [job](#job)

---

### credentials

Common fields for all credentials:

`identifier` [required] - id under which given credential will be stored in jenkins.
`username` [required] - username for given credentials.
`scope` [optional] - scope for credential in jenkins ( default to `global` ).
`domain` [optional] - domain for credential in jenkins ( defaults to `_` ).
`description` [optional] - describes what credential is used for.
`jobpath` [option] - Empty for no job location, or name of job to store under specific job. ( defaults to "" ).

### jenkins_username_credential

```hcl
resource "jenkins_username_credential" "admin" {
    identifier = "cred_id"
    username = "admin"
    password = "admin"
}
```

This will create username credential.

`password` [required] - password for given user.

### jenkins_ssh_credential

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
`jobpath` [option] - Empty for no job location, or name of job to store under specific job. ( defaults to "" ).

This will create ssh credential, ssh key can be provided as inline value or as path on jenkins master.

### jenkins_docker_credential

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
`jobpath` [option] - Empty for no job location, or name of job to store under specific job. ( defaults to "" ).


### jenkins_secret_credential

```hcl
resource "jenkins_secret_credential" "secret_text" {
  identifier = "secret_text_id"
  secret = "secret"
  jobpath = ""
  domain = "_"
  description = "secret text description"
}
```

This will secret text credential.

`secret` [required] - secret text.
`jobpath` [option] - Empty for no job location, or name of job to store under specific job. ( defaults to "" ).
`domain` [optional] - domain for credential in jenkins ( defaults to `_` ).


## plugin

```hcl
"jenkins_plugin" "packer" {
    name = "packer"
    version = "1.4"
}
```

Will try to install plugin with given name and version
`name` [required] - id of plugin to install.
`version` [required] - version of plugin to install

## job

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
