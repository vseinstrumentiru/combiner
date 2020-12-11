# Config Files Combiner
A very helpful tool for generating values.yaml for Helm configuration.

### Usage
```bash
combiner combine group1:file1,file2 file3 group3:file4,file5
```

It will be created file `values.yaml` with:
```yaml
group1:
#  data from ./values/default.yaml
#  merged and overwritten with ./values/file1.yaml
#  merged and overwritten with ./values/file2.yaml

file3:
#  data from ./values/default.yaml
#  merged and overwritten with ./values/file3.yaml

group3:
#  data from ./values/default.yaml
#  merged and overwritten with ./values/file4.yaml
#  merged and overwritten with ./values/file5.yaml
```

#### Flags:
###### `-d`, `--default`
Default config file name (without extension). This is base config file for other groups (default "default")
###### Deprecated `-n`, `--no-default`
Without default file config
####### `-o`, `--out string`
Output file name (default "values.yaml")
####### `-p`, `--path string`
Folder with yaml configs to merge (default "./values")

#### Run with combine.yaml
1. Define `combine.yaml` file with structure:
```yaml
.default: &default:
  files:
  - file1

combine:
- out: "values.yaml" // combined file name. default "values.yaml"
  path: "./values" // where chunks located. default "./values"
  groups:
    # range from 1 to 2: group1, group2
    - name: "group%v"
      range: [1, 2]
      <<: *default
    # add files to group1
    - name: group1
      files:
      - file2
    # add files to group2
    - name: group2
      files:
      - file3
      - file4
- out: "values2.yaml"
  groups:
    - name: group1
      <<: *default
    - name: group1
      files:
      - file3
```
2. Run `combiner` in folder with `combine.yaml`.
It will create files values.yaml (with sections group1 and group2) and values2.yaml (with the section group1)
---

##### Dependencies:
- github.com/spf13/cobra
- github.com/spf13/viper
