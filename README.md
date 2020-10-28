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

### Flags:
##### `-d`, `--default`
Default config file name (without extension). This is base config file for other groups (default "default")
##### `-n`, `--no-default`
Without default file config
###### `-o`, `--out string`
Output file name (default "values.yaml")
###### `-p`, `--path string`
Folder with yaml configs to merge (default "./values")

---

##### Dependencies:
- github.com/spf13/cobra
- github.com/spf13/viper
