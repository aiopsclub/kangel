## ⚡️ Quick start
The kangel is the angel daemonset for kubelet, it can make kubelet more stable.

First of all, it require Helm, The chart in the chart dir. 

shell detect:
```shell
#!/bin/bash

main(){
    for i in `docker ps | grep -v "CONTAINER"| awk '{print $NF}'`;do echo "checking $i" && docker inspect $i > /dev/null && docker stats --no-stream $i;done
}
main $@
```


## 📖 kubernets issue   
1. container runtime status check may not have completed yet, PLEG is not healthy.  

knagel can found which container cause this issue and fix it.

## ⚙️ Commands & Options

### `deploy`

```bash
helm install .
```
## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `Create Go App CLI`:

Together, we can make this project **better** every day! 😘
