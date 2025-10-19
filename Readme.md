



# Настройка cAdvisor для WSL2
**Проблема:** /var/lib/docker/ в случае wsl2 - пуст. Информация, которую ищет cadvisor, находится в файлах wsl: `\\wsl.localhost\docker-desktop\mnt\docker-desktop-disk\data\docker/`. Надо настроить доступ из ubuntu в файлам wsl.

[Ссылка на issue](https://github.com/vacp2p/wakurtosis/issues/58)

### Сделать при каждом запуске wsl2
```bash
./mount_docker.sh
```
или
``` bash
$ ls /mnt/
c  e  wsl  wslg

$ sudo mkdir /mnt/windows_docker
# Docker Desktop должен работать
$ sudo mount -t drvfs '\\wsl.localhost\docker-desktop\mnt\docker-desktop-disk\data\docker' /mnt/windows_docker
```
### docker compose
```
  cadvisortest:
    image: gcr.io/cadvisor/cadvisor:latest
    container_name: cadvisortest
    hostname: cadvisortest
    ports:
      - "8081:8080" 
    privileged: true
    devices:
      - /dev/kmsg:/dev/kmsg
    volumes:
      - /:/rootfs:ro
      - /var/run:/var/run:rw
      - /sys:/sys:ro
      - /var/lib/docker:/var/lib/docker:ro
      - /dev/disk/:/dev/disk:ro
      - /etc/machine-id:/etc/machine-id:ro
      - /mnt/windows_docker/:/rootfs/var/lib/docker:ro # особенность wsl
    networks:
      - test-network
```

container_cpu_usage_seconds_total{name="deployment-gin-app-1"}
rate(container_cpu_usage_seconds_total{name="deployment-gin-app-1"}[1m])*100

container_memory_usage_bytes{name="deployment-gin-app-1"}

 
container_spec_memory_limit_bytes{name="deployment-gin-app-1"}
container_memory_rss{name="deployment-gin-app-1"} 
container_memory_cache{name="deployment-gin-app-1"} 

container_fs_reads_bytes_total{name="deployment-gin-app-1"} 
container_fs_writes_bytes_total{name="deployment-gin-app-1"} 
  