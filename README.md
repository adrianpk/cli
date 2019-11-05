# CLI
A command line tool for [mikrowezel backend](https://gitlab.com/mikrowezel/backend).

## Install

Work in progress.

```shell
$ go get -u https://gitlab.com/mikrowezel/backend/cli
$ alias mw=cli
```

Another option to avoid alias

```shell
$ git clone https://gitlab.com/mikrowezel/backend/cli.git
$ cd cli
$ make install
```

## Usage
### **generate** command

**Input file**

Create an input file for the resource under assets/gen
i.e.: `assets/gen/sample.yaml`
The name of the file sholud be lowercased name of the resource to be generated using `.yaml` as extension.

In this example we are going to create the resource `Sample`

```yaml
---
  name: Sample
  #plural: Samples
  pkg:
    main: gitlab.com/username/project
    servicePath: svcname
  api:
    version: v0.0.1
  propDefs:
    - name: ID
      type: uuid
      length: 36
      #isVirtual: false
      isKey: true
      #isUnique: true
      AdmitNull: false
    - name: OwnerID
      type: uuid
      length: 36
      #isVirtual: false
      #isKey: false
      #isUnique: true
      AdmitNull: false
      ref:
        model: user
        property: ID
        #fkName: user_id_fk
        #targetTable: users
    - name: Name
      type: string
      length: 255
      #isVirtual: false
      #isKey: false
      #isUnique: true
      AdmitNull: true
```

**Command**

This is the generic format for the generator command:

```shell
mw generate <ModelName> [-p pkgName] [--all] [--yaml] [--migration] [--model] [--repo] [--grpc] [--jsonrest] [--service] [-transport] [--web] [--restcl] [--force]
``````

```shell
$ mw generate Sample -all
```

**Where**

  * **generate** is the name of the command.
  * **ModelName** is the name of the resource to be created.

  * Flags:
    * **-p** let overwrite the package name declared in sample.yaml.
    * **--all** generates all the files for the resource named <ModelName>.
    * **--yaml** generates a base modelname.yaml skeleton file under `assets/gen`.
    * **--migration** generates a model file for <ModelName> resource under `pkgName/internal/migration`.
    * **--model** generates a model file for <ModelName> resource under `pkgName/internal/model`.
    * **--model** generates a repo file for <ModelName> resource under `pkgName/internal/repo`.
    * **--grpc** generates grpc endpoint files for <ModelName> resource under `pkgName/pkg/servicePath/grpc`.
    * **--jsonrest** generates JSON REST endpoint files for <ModelName> resource under `pkgName/pkg/servicePath/jsonrest`.
    * **--service** generates service files for <ModelName> resource under `pkgName/pkg/servicePath`.
    * **--transport** generates transport files for <ModelName> resource under `pkgName/pkg/transport`.
    * **--web** generates web files for <ModelName> resource under `pkgName/pkg/web`.
    * **--rest** generate cURL and JSON files for <ModelName> resource under `pkgName/scripts/rest`.
    * **--force** Overwrite files if already exist.

Please, observe that *pkgName* and *servicePath* refer to those defined in the YAML file.

**Note:**

This is a work in progress, generators are being implemented so right now not all the options are available.

At this stage the output should look similar to this:

```shell
2019/11/05 22:35:01 Starting...
2019/11/05 22:35:01 Reading input file: 'assets/gen/sample.yaml'
2019/11/05 22:35:01 Generating metadata
2019/11/05 22:35:01 Migration file: internal/migration/20191105223501createtablesamples.go
2019/11/05 22:35:01 Done!
2019/11/05 22:35:01 Model file: internal/model/sample.go
2019/11/05 22:35:01 Done!
2019/11/05 22:35:01 Repo file: internal/repo/sample.go
2019/11/05 22:35:01 Done!
2019/11/05 22:35:01 Transport file: pkg/auth/transport/sample.go
2019/11/05 22:35:01 Transport file: pkg/auth/transport/samplecnv.go
2019/11/05 22:35:01 Done!
```
